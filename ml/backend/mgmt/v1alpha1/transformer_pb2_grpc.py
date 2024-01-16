# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from mgmt.v1alpha1 import transformer_pb2 as mgmt_dot_v1alpha1_dot_transformer__pb2


class TransformersServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetSystemTransformers = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/GetSystemTransformers',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersResponse.FromString,
                )
        self.GetUserDefinedTransformers = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformers',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersResponse.FromString,
                )
        self.GetUserDefinedTransformerById = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformerById',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdResponse.FromString,
                )
        self.CreateUserDefinedTransformer = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/CreateUserDefinedTransformer',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerResponse.FromString,
                )
        self.DeleteUserDefinedTransformer = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/DeleteUserDefinedTransformer',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerResponse.FromString,
                )
        self.UpdateUserDefinedTransformer = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/UpdateUserDefinedTransformer',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerResponse.FromString,
                )
        self.IsTransformerNameAvailable = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/IsTransformerNameAvailable',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableResponse.FromString,
                )
        self.ValidateUserJavascriptCode = channel.unary_unary(
                '/mgmt.v1alpha1.TransformersService/ValidateUserJavascriptCode',
                request_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeRequest.SerializeToString,
                response_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeResponse.FromString,
                )


class TransformersServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetSystemTransformers(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserDefinedTransformers(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserDefinedTransformerById(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateUserDefinedTransformer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteUserDefinedTransformer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUserDefinedTransformer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def IsTransformerNameAvailable(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ValidateUserJavascriptCode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TransformersServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetSystemTransformers': grpc.unary_unary_rpc_method_handler(
                    servicer.GetSystemTransformers,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersResponse.SerializeToString,
            ),
            'GetUserDefinedTransformers': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserDefinedTransformers,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersResponse.SerializeToString,
            ),
            'GetUserDefinedTransformerById': grpc.unary_unary_rpc_method_handler(
                    servicer.GetUserDefinedTransformerById,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdResponse.SerializeToString,
            ),
            'CreateUserDefinedTransformer': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateUserDefinedTransformer,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerResponse.SerializeToString,
            ),
            'DeleteUserDefinedTransformer': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteUserDefinedTransformer,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerResponse.SerializeToString,
            ),
            'UpdateUserDefinedTransformer': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateUserDefinedTransformer,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerResponse.SerializeToString,
            ),
            'IsTransformerNameAvailable': grpc.unary_unary_rpc_method_handler(
                    servicer.IsTransformerNameAvailable,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableResponse.SerializeToString,
            ),
            'ValidateUserJavascriptCode': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateUserJavascriptCode,
                    request_deserializer=mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeRequest.FromString,
                    response_serializer=mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'mgmt.v1alpha1.TransformersService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TransformersService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetSystemTransformers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/GetSystemTransformers',
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetSystemTransformersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserDefinedTransformers(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformers',
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserDefinedTransformerById(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformerById',
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.GetUserDefinedTransformerByIdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateUserDefinedTransformer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/CreateUserDefinedTransformer',
            mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.CreateUserDefinedTransformerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteUserDefinedTransformer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/DeleteUserDefinedTransformer',
            mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.DeleteUserDefinedTransformerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateUserDefinedTransformer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/UpdateUserDefinedTransformer',
            mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.UpdateUserDefinedTransformerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def IsTransformerNameAvailable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/IsTransformerNameAvailable',
            mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.IsTransformerNameAvailableResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ValidateUserJavascriptCode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/mgmt.v1alpha1.TransformersService/ValidateUserJavascriptCode',
            mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeRequest.SerializeToString,
            mgmt_dot_v1alpha1_dot_transformer__pb2.ValidateUserJavascriptCodeResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
