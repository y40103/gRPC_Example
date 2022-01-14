import time

import Hello_pb2
import Hello_pb2_grpc
import grpc
from concurrent import futures
import logging
from typing import Generator


class HelloMethod(Hello_pb2_grpc.TestServicer):
    """ Service中所需要用到的所有方法 """

    def SayHello0(self, request: Hello_pb2.RequestType, context) -> Hello_pb2.ResponseType:
        ## from_request 實際上是 protobuf通訊用的 RequestType的instance,
        ## 它具有當初 protobuf 文檔所定義的RequestType屬性 , 可理解成 request 進來的參數

        res = Hello_pb2.ResponseType()
        ## 產生一個 protobuf response 的 instance  同上述 request地位 它具有當初 protobuf 文檔所定義的ResponseType屬性

        res.sentence = f"{request.name} say {request.content}"  ## 該函數定義的內容 可自由發揮

        return res

    def SayHello1(self, request_iterator: Generator[Hello_pb2.RequestType, None, None],
                  context) -> Hello_pb2.ResponseType:
        sentence = "'''\n"
        for each in request_iterator:
            content = f"{each.name} say {each.content}\n"
            sentence += content
        sentence = sentence + "'''"

        res = Hello_pb2.ResponseType()  # 產生一個response type 的 instance
        res.sentence = sentence  # 把pythob string 賦值 給 protobuf 的 response

        return res

    def SayHello2(self, request: Hello_pb2.RequestType, context) -> Generator[Hello_pb2.ResponseType, None, None]:
        res = Hello_pb2.ResponseType()
        for times in range(1, 6):
            res.sentence = f"{request.name} say {request.content * times}"
            yield res

    def SayHello3(self, request: Generator[Hello_pb2.RequestType, None, None], context) -> Generator[
        Hello_pb2.ResponseType, None, None]:
        res = Hello_pb2.ResponseType()
        for num, each_request in enumerate(request):
            res.sentence = f"{each_request.name} say {each_request.content * (num + 1)}"
            yield res


def gRPC_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))  ## create server instance
    Hello_pb2_grpc.add_TestServicer_to_server(HelloMethod(), server)  ## add method to server
    server.add_insecure_port("0.0.0.0:50051")  ## add host
    server.start()
    try:
        while True:
            time.sleep(3600000000)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    gRPC_server()
