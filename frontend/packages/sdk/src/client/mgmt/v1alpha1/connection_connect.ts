// @generated by protoc-gen-connect-es v1.0.0 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/connection.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CheckConnectionConfigRequest, CheckConnectionConfigResponse, CheckSqlQueryRequest, CheckSqlQueryResponse, CreateConnectionRequest, CreateConnectionResponse, DeleteConnectionRequest, DeleteConnectionResponse, GetConnectionDataStreamRequest, GetConnectionDataStreamResponse, GetConnectionForeignConstraintsRequest, GetConnectionForeignConstraintsResponse, GetConnectionInitStatementsRequest, GetConnectionInitStatementsResponse, GetConnectionRequest, GetConnectionResponse, GetConnectionSchemaRequest, GetConnectionSchemaResponse, GetConnectionsRequest, GetConnectionsResponse, IsConnectionNameAvailableRequest, IsConnectionNameAvailableResponse, UpdateConnectionRequest, UpdateConnectionResponse } from "./connection_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Service for managing datasource connections.
 * This is a primary data model in Neosync and is used in reference when hooking up Jobs to synchronize and generate data.
 *
 * @generated from service mgmt.v1alpha1.ConnectionService
 */
export const ConnectionService = {
  typeName: "mgmt.v1alpha1.ConnectionService",
  methods: {
    /**
     * Returns a list of connections associated with the account
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnections
     */
    getConnections: {
      name: "GetConnections",
      I: GetConnectionsRequest,
      O: GetConnectionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Returns a single connection
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnection
     */
    getConnection: {
      name: "GetConnection",
      I: GetConnectionRequest,
      O: GetConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Creates a new connection
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.CreateConnection
     */
    createConnection: {
      name: "CreateConnection",
      I: CreateConnectionRequest,
      O: CreateConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Updates an existing connection
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.UpdateConnection
     */
    updateConnection: {
      name: "UpdateConnection",
      I: UpdateConnectionRequest,
      O: UpdateConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Removes a connection from the system.
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.DeleteConnection
     */
    deleteConnection: {
      name: "DeleteConnection",
      I: DeleteConnectionRequest,
      O: DeleteConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Connections have friendly names, this method checks if the requested name is available in the system based on the account
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.IsConnectionNameAvailable
     */
    isConnectionNameAvailable: {
      name: "IsConnectionNameAvailable",
      I: IsConnectionNameAvailableRequest,
      O: IsConnectionNameAvailableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Checks if the connection config is connectable by the backend.
     * Used mostly to verify that a connection is valid prior to creating a Connection object.
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.CheckConnectionConfig
     */
    checkConnectionConfig: {
      name: "CheckConnectionConfig",
      I: CheckConnectionConfigRequest,
      O: CheckConnectionConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Returns the schema for a specific connection. Used mostly for SQL-based connections
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnectionSchema
     */
    getConnectionSchema: {
      name: "GetConnectionSchema",
      I: GetConnectionSchemaRequest,
      O: GetConnectionSchemaResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Checks a constructed SQL query against a sql-based connection to see if it's valid based on that connection's data schema
     * This is useful when constructing subsets to see if the WHERE clause is correct
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.CheckSqlQuery
     */
    checkSqlQuery: {
      name: "CheckSqlQuery",
      I: CheckSqlQueryRequest,
      O: CheckSqlQueryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Streaming endpoint that will stream the data available from the Connection to the client.
     * Used primarily by the CLI sync command.
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnectionDataStream
     */
    getConnectionDataStream: {
      name: "GetConnectionDataStream",
      I: GetConnectionDataStreamRequest,
      O: GetConnectionDataStreamResponse,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * For a specific connection, returns the foreign key constraints. Mostly useful for SQL-based Connections.
     * Used primarily by the CLI sync command to determine stream order.
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnectionForeignConstraints
     */
    getConnectionForeignConstraints: {
      name: "GetConnectionForeignConstraints",
      I: GetConnectionForeignConstraintsRequest,
      O: GetConnectionForeignConstraintsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * For a specific connection, returns the init table statements. Mostly useful for SQL-based Connections.
     * Used primarily by the CLI sync command to create table schema init statement.
     *
     * @generated from rpc mgmt.v1alpha1.ConnectionService.GetConnectionInitStatements
     */
    getConnectionInitStatements: {
      name: "GetConnectionInitStatements",
      I: GetConnectionInitStatementsRequest,
      O: GetConnectionInitStatementsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

