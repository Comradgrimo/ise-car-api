import asyncio

from grpclib.client import Channel

from ozonmp.ise_car_api.v1.ise_car_api_grpc import IseCarApiServiceStub
from ozonmp.ise_car_api.v1.ise_car_api_pb2 import DescribeCarV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = IseCarApiServiceStub(channel)

        req = DescribeCarV1Request(car_id=1)
        reply = await client.DescribeCarV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
