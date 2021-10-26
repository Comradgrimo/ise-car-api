# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from ozonmp.omp_template_api.v1 import omp_template_api_pb2 as ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2


class OmpTemplateApiServiceStub(object):
    """OmpTemplateApiService - Service for working with templates
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DescribeTemplateV1 = channel.unary_unary(
                '/ozonmp.omp_template_api.v1.OmpTemplateApiService/DescribeTemplateV1',
                request_serializer=ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Request.SerializeToString,
                response_deserializer=ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Response.FromString,
                )


class OmpTemplateApiServiceServicer(object):
    """OmpTemplateApiService - Service for working with templates
    """

    def DescribeTemplateV1(self, request, context):
        """DescribeTemplateV1 - Describe a template
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OmpTemplateApiServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DescribeTemplateV1': grpc.unary_unary_rpc_method_handler(
                    servicer.DescribeTemplateV1,
                    request_deserializer=ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Request.FromString,
                    response_serializer=ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Response.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ozonmp.omp_template_api.v1.OmpTemplateApiService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OmpTemplateApiService(object):
    """OmpTemplateApiService - Service for working with templates
    """

    @staticmethod
    def DescribeTemplateV1(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ozonmp.omp_template_api.v1.OmpTemplateApiService/DescribeTemplateV1',
            ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Request.SerializeToString,
            ozonmp_dot_omp__template__api_dot_v1_dot_omp__template__api__pb2.DescribeTemplateV1Response.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
