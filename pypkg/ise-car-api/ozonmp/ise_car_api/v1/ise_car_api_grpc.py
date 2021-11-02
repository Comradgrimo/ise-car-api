# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/ise_car_api/v1/ise_car_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.ise_car_api.v1.ise_car_api_pb2


class IseCarApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribeCarV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Request, ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListCarsV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Request, ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def CreateCarV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Request, ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemoveCarV1(self, stream: 'grpclib.server.Stream[ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Request, ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.ise_car_api.v1.IseCarApiService/DescribeCarV1': grpclib.const.Handler(
                self.DescribeCarV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Request,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Response,
            ),
            '/ozonmp.ise_car_api.v1.IseCarApiService/ListCarsV1': grpclib.const.Handler(
                self.ListCarsV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Request,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Response,
            ),
            '/ozonmp.ise_car_api.v1.IseCarApiService/CreateCarV1': grpclib.const.Handler(
                self.CreateCarV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Request,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Response,
            ),
            '/ozonmp.ise_car_api.v1.IseCarApiService/RemoveCarV1': grpclib.const.Handler(
                self.RemoveCarV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Request,
                ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Response,
            ),
        }


class IseCarApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribeCarV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_car_api.v1.IseCarApiService/DescribeCarV1',
            ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Request,
            ozonmp.ise_car_api.v1.ise_car_api_pb2.DescribeCarV1Response,
        )
        self.ListCarsV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_car_api.v1.IseCarApiService/ListCarsV1',
            ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Request,
            ozonmp.ise_car_api.v1.ise_car_api_pb2.ListCarsV1Response,
        )
        self.CreateCarV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_car_api.v1.IseCarApiService/CreateCarV1',
            ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Request,
            ozonmp.ise_car_api.v1.ise_car_api_pb2.CreateCarV1Response,
        )
        self.RemoveCarV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.ise_car_api.v1.IseCarApiService/RemoveCarV1',
            ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Request,
            ozonmp.ise_car_api.v1.ise_car_api_pb2.RemoveCarV1Response,
        )
