import grpc
import Hello_pb2
import Hello_pb2_grpc
from typing import Generator

host = grpc.insecure_channel("localhost:50051")

client = Hello_pb2_grpc.TestStub(host) ## TestStub 可理解成 打包 與 解包工具 , Test這邊是指 protobuf schema定義的service名稱

request = Hello_pb2.RequestType()


def declare_generator() -> Generator[Hello_pb2.RequestType, None, None]:
    """
    generate a generator
    """
    IP = ["卡比獸", "嘎啦嘎啦", "多啦a夢", "約得爾賤畜", "嚕嚕咪"]
    echo = ["ZZ ", "嘎啦 ", "大雄啊 ", "1234 ", "@&^*$ "]
    req = Hello_pb2.RequestType()
    for index in range(0, 5):
        req.name = IP[index]
        req.content = echo[index]
        yield req


request = declare_generator()
res = client.SayHello3(request)

for each in res:
    print(each.sentence)
