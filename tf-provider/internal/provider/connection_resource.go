package provider

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
)

var _ resource.Resource = &ConnectionResource{}
var _ resource.ResourceWithImportState = &ConnectionResource{}

func NewConnectionResource() resource.Resource {
	return &ConnectionResource{}
}

type ConnectionResource struct {
	client    mgmtv1alpha1connect.ConnectionServiceClient
	accountId *string
}

type ConnectionResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	AccountId types.String `tfsdk:"account_id"`

	Postgres *Postgres `tfsdk:"postgres"`
}

type Postgres struct {
	Url types.String `tfsdk:"url"`

	Host    types.String `tfsdk:"host"`
	Port    types.Int64  `tfsdk:"port"`
	Name    types.String `tfsdk:"name"`
	User    types.String `tfsdk:"user"`
	Pass    types.String `tfsdk:"pass"`
	SslMode types.String `tfsdk:"ssl_mode"`

	Tunnel *SSHTunnel `tfsdk:"tunnel"`
}

type SSHTunnel struct {
	Host               types.String `tfsdk:"host"`
	Port               types.Int64  `tfsdk:"port"`
	User               types.String `tfsdk:"user"`
	KnownHostPublicKey types.String `tfsdk:"known_host_public_key"`

	PrivateKey types.String `tfsdk:"private_key"`
	Passphrase types.String `tfsdk:"passphrase"`
}

func (r *ConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection"
}

func (r *ConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The unique friendly name of the connection",
				Required:            true,
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the account. Can be pulled from the API Key if present, or must be specified if using a user access token",
				Optional:            true,
				Computed:            true,
				PlanModifiers:       []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"postgres": schema.SingleNestedAttribute{
				Description: "The postgres database that will be associated with this connection",
				Optional:    true,
				// PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Description: "Standard postgres url connection string. Must be uri compliant",
						Optional:    true,
					},

					"host": schema.StringAttribute{
						Description: "The host name of the postgres server",
						Optional:    true,
					},
					"port": schema.Int64Attribute{
						Description: "The post of the postgres server",
						Optional:    true,
						// Default:     int64default.StaticInt64(5432),
					},
					"name": schema.StringAttribute{
						Description: "The name of the database that will be connected to",
						Optional:    true,
					},
					"user": schema.StringAttribute{
						Description: "The name of the user that will be authenticated with",
						Optional:    true,
					},
					"pass": schema.StringAttribute{
						Description: "The password that will be authenticated with",
						Optional:    true,
						Sensitive:   true,
					},
					"ssl_mode": schema.StringAttribute{
						Description: "The SSL mode for the postgres server",
						Optional:    true,
					},
					"tunnel": schema.SingleNestedAttribute{
						Description: "SSH tunnel that is used to access databases that are not publicly accessible to the internet",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description: "The host name of the server",
								Required:    true,
							},
							"port": schema.Int64Attribute{
								Description: "The post of the ssh server",
								Required:    true,
								// Default:     int64default.StaticInt64(22),
							},
							"user": schema.StringAttribute{
								Description: "The name of the user that will be authenticated with",
								Required:    true,
							},
							"known_host_public_key": schema.StringAttribute{
								Description: "The known SSH public key of the tunnel server.",
								Optional:    true,
							},
							"private_key": schema.StringAttribute{
								Description: "If using key authentication, this must be a pem encoded private key",
								Optional:    true,
								Sensitive:   true,
							},
							"passphrase": schema.StringAttribute{
								Description: "If not using key authentication, a password must be provided. If a private key is provided, but encrypted, provide the passphrase here as it will be used to decrypt the private key",
								Optional:    true,
								Sensitive:   true,
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier of the connection",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
		},
	}
}

func (r *ConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*ConfigData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected mgmtv1alpha1connect.ConnectionServiceClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = providerData.ConnectionClient
	r.accountId = providerData.AccountId
}

func hydrateTunnelFromTunnelConfig(t *mgmtv1alpha1.SSHTunnel) *SSHTunnel {
	if t == nil {
		return nil
	}

	tunnel := &SSHTunnel{
		Host:               types.StringValue(t.Host),
		Port:               types.Int64Value(int64(t.Port)),
		User:               types.StringValue(t.User),
		KnownHostPublicKey: types.StringPointerValue(t.KnownHostPublicKey),
	}

	if t.Authentication != nil {
		switch auth := t.Authentication.AuthConfig.(type) {
		case *mgmtv1alpha1.SSHAuthentication_PrivateKey:
			tunnel.PrivateKey = types.StringValue(auth.PrivateKey.Value)
			tunnel.Passphrase = types.StringPointerValue(auth.PrivateKey.Passphrase)
		case *mgmtv1alpha1.SSHAuthentication_Passphrase:
			tunnel.Passphrase = types.StringValue(auth.Passphrase.Value)
		}
	}
	return tunnel
}

func hydrateResourceModelFromConnectionConfig(cc *mgmtv1alpha1.ConnectionConfig, data *ConnectionResourceModel) error {
	switch config := cc.Config.(type) {
	case *mgmtv1alpha1.ConnectionConfig_PgConfig:
		switch pgcc := config.PgConfig.ConnectionConfig.(type) {
		case *mgmtv1alpha1.PostgresConnectionConfig_Connection:
			data.Postgres = &Postgres{
				Host:    types.StringValue(pgcc.Connection.Host),
				Port:    types.Int64Value(int64(pgcc.Connection.Port)),
				Name:    types.StringValue(pgcc.Connection.Name),
				User:    types.StringValue(pgcc.Connection.User),
				Pass:    types.StringValue(pgcc.Connection.Pass),
				SslMode: types.StringPointerValue(pgcc.Connection.SslMode),
				Tunnel:  hydrateTunnelFromTunnelConfig(config.PgConfig.Tunnel),
			}

			return nil
		case *mgmtv1alpha1.PostgresConnectionConfig_Url:
			data.Postgres = &Postgres{
				Url:    types.StringValue(pgcc.Url),
				Tunnel: hydrateTunnelFromTunnelConfig(config.PgConfig.Tunnel),
			}
			return nil
		default:
			return errors.New("unable to find a config to hydrate connection resource model")
		}
	default:
		return errors.New("unable to find a config to hydrate connection resource model")
	}
}

func getConnectionConfigFromResourceModel(data *ConnectionResourceModel) (*mgmtv1alpha1.ConnectionConfig, error) {
	if data.Postgres != nil {
		var tunnel *mgmtv1alpha1.SSHTunnel
		if data.Postgres.Tunnel != nil {
			tunnel = &mgmtv1alpha1.SSHTunnel{
				Host:               data.Postgres.Tunnel.Host.ValueString(),
				Port:               int32(data.Postgres.Tunnel.Port.ValueInt64()),
				User:               data.Postgres.Tunnel.User.ValueString(),
				KnownHostPublicKey: data.Postgres.Tunnel.KnownHostPublicKey.ValueStringPointer(),
			}
			if data.Postgres.Tunnel.PrivateKey.ValueString() != "" {
				tunnel.Authentication = &mgmtv1alpha1.SSHAuthentication{
					AuthConfig: &mgmtv1alpha1.SSHAuthentication_PrivateKey{
						PrivateKey: &mgmtv1alpha1.SSHPrivateKey{
							Value:      data.Postgres.Tunnel.PrivateKey.ValueString(),
							Passphrase: data.Postgres.Tunnel.Passphrase.ValueStringPointer(),
						},
					},
				}
			} else if data.Postgres.Tunnel.Passphrase.ValueString() != "" {
				tunnel.Authentication = &mgmtv1alpha1.SSHAuthentication{
					AuthConfig: &mgmtv1alpha1.SSHAuthentication_Passphrase{
						Passphrase: &mgmtv1alpha1.SSHPassphrase{
							Value: data.Postgres.Tunnel.Passphrase.ValueString(),
						},
					},
				}
			}
		}
		if data.Postgres.Url.ValueString() != "" {
			return &mgmtv1alpha1.ConnectionConfig{
				Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
					PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
						ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Url{
							Url: data.Postgres.Url.ValueString(),
						},
						Tunnel: tunnel,
					},
				},
			}, nil
		} else {
			pg := data.Postgres
			if pg.Host.ValueString() == "" || pg.Port.ValueInt64() == 0 || pg.Name.ValueString() == "" || pg.User.ValueString() == "" || pg.Pass.ValueString() == "" {
				return nil, fmt.Errorf("invalid postgres config")
			}
			return &mgmtv1alpha1.ConnectionConfig{
				Config: &mgmtv1alpha1.ConnectionConfig_PgConfig{
					PgConfig: &mgmtv1alpha1.PostgresConnectionConfig{
						ConnectionConfig: &mgmtv1alpha1.PostgresConnectionConfig_Connection{
							Connection: &mgmtv1alpha1.PostgresConnection{
								Host:    pg.Host.ValueString(),
								Port:    int32(pg.Port.ValueInt64()),
								Name:    pg.Name.ValueString(),
								User:    pg.User.ValueString(),
								Pass:    pg.Pass.ValueString(),
								SslMode: pg.SslMode.ValueStringPointer(),
							},
						},
						Tunnel: tunnel,
					},
				},
			}, nil
		}
	}
	return nil, errors.New("invalid connection config")
}

// nolint
func (r *ConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ConnectionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var accountId string
	if data.AccountId.ValueString() == "" {
		if r.accountId != nil {
			accountId = *r.accountId
		}
	} else {
		accountId = data.AccountId.ValueString()
	}
	if accountId == "" {
		resp.Diagnostics.AddError("no account id", "must provide account id either on the resource or provide through environment configuration")
		return
	}

	cc, err := getConnectionConfigFromResourceModel(&data)
	if err != nil {
		resp.Diagnostics.AddError("connection config error", err.Error())
		return
	}
	connResp, err := r.client.CreateConnection(ctx, connect.NewRequest(&mgmtv1alpha1.CreateConnectionRequest{
		Name:             data.Name.ValueString(),
		AccountId:        accountId, // compute account id
		ConnectionConfig: cc,
	}))
	if err != nil {
		resp.Diagnostics.AddError("create connection error", err.Error())
		return
	}

	connection := connResp.Msg.Connection

	data.Id = types.StringValue(connection.Id)
	data.Name = types.StringValue(connection.Name)
	data.AccountId = types.StringValue(connection.AccountId)
	err = hydrateResourceModelFromConnectionConfig(connection.ConnectionConfig, &data)
	if err != nil {
		resp.Diagnostics.AddError("connection config hydration error", err.Error())
		return
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created connection resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// nolint
func (r *ConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ConnectionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	connResp, err := r.client.GetConnection(ctx, connect.NewRequest(&mgmtv1alpha1.GetConnectionRequest{
		Id: data.Id.ValueString(),
	}))
	if err != nil {
		resp.Diagnostics.AddError("Unable to get connection", err.Error())
		return
	}

	connection := connResp.Msg.Connection

	data.Id = types.StringValue(connection.Id)
	data.Name = types.StringValue(connection.Name)
	data.AccountId = types.StringValue(connection.AccountId)
	err = hydrateResourceModelFromConnectionConfig(connection.ConnectionConfig, &data)
	if err != nil {
		resp.Diagnostics.AddError("connection config hydration error", err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// nolint
func (r *ConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ConnectionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	cc, err := getConnectionConfigFromResourceModel(&data)
	if err != nil {
		resp.Diagnostics.AddError("connection config error", err.Error())
		return
	}

	connResp, err := r.client.UpdateConnection(ctx, connect.NewRequest(&mgmtv1alpha1.UpdateConnectionRequest{
		Id:               data.Id.ValueString(),
		Name:             data.Name.ValueString(),
		ConnectionConfig: cc,
	}))
	if err != nil {
		resp.Diagnostics.AddError("Unable to update connection", err.Error())
		return
	}

	connection := connResp.Msg.Connection

	data.Id = types.StringValue(connection.Id)
	data.Name = types.StringValue(connection.Name)
	data.AccountId = types.StringValue(connection.AccountId)
	err = hydrateResourceModelFromConnectionConfig(connection.ConnectionConfig, &data)
	if err != nil {
		resp.Diagnostics.AddError("connection config hydration error", err.Error())
		return
	}

	tflog.Trace(ctx, "updated connection")
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// nolint
func (r *ConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ConnectionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DeleteConnection(ctx, connect.NewRequest(&mgmtv1alpha1.DeleteConnectionRequest{
		Id: data.Id.ValueString(),
	}))
	if err != nil {
		resp.Diagnostics.AddError("Unable to delete connection", err.Error())
		return
	}

	tflog.Trace(ctx, "deleted connection")
}

func (r *ConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	// todo
}