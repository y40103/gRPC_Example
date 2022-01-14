import grpc
import Hello_pb2
import Hello_pb2_grpc
from typing import Generator

## python gRPC , stream 的類型 為 用 generator打包


host = grpc.insecure_channel("localhost:50051")

client = Hello_pb2_grpc.TestStub(host)


def declare_stream_request() -> Generator[Hello_pb2.RequestType, None, None]:
    """
    generate a generator
    """

    for num in range(1, 6):
        request = Hello_pb2.RequestType()  ## 產生一個 RequestType 的 instance
        request.name = f"小火龍{num}號"
        request.content = ("嘎嘎 " * num)
        yield request


request_generator = declare_stream_request()  ## 產生 宣告多個 request 的 generator

res = client.SayHello1(request_generator)

print(res.sentence)
