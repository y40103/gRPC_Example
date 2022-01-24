import time
import Hello_pb2
import Hello_pb2_grpc
import grpc
from concurrent import futures
import logging
from typing import Generator


class TestExample(Hello_pb2_grpc.TestServicer):
    """ Service中所需要用到的方法 """


    ## 該class包含該 service 所有方法,
    ## 初始輸入與最終輸出函數 皆必須為 protobuf 的特定類型 , 且符合 protobuf schema 所定義

    ## 通常輸入已為 protobuf 的 instance, 可由該instance內部屬性取得數值 , 函數內部可自行定義 , 最終輸出再封裝成 protobuf instance 回傳



    ## 以下測試 定義類型 , 函數方法內容 輸入=輸出

    def TypeSimple(self, request, context) -> Hello_pb2.SimpleType:
        res = Hello_pb2.SimpleType()            # 實例化 protobuf的 instance

        res.int_num = request.int_num
        res.float_num = request.float_num
        res.description = request.description
        res.bool_flag = request.bool_flag

        return res


    def TypeEnum(self, request, context) -> Hello_pb2.Enumeration:

        return request


    def TypeList(self, request, context) -> Hello_pb2.ListType:
        res = Hello_pb2.ListType()

        for each_num in request.int_list:
            res.int_list.append(each_num*10)

        return res

    def TypeNested(self, request, context) -> Hello_pb2.NestedType:
        res = Hello_pb2.NestedType()
        res.RequestType.name = request.RequestType.name
        res.RequestType.content = request.RequestType.content

        return res

    def TypeMap(self, request, context) -> Hello_pb2.MapType:
        res = Hello_pb2.MapType()
        res.map1["巧克力"] = request.map1["巧克力"]
        res.map1["泡芙"] = request.map1["泡芙"]
        res.map1["海苔"] = request.map1["海苔"]
        res.map2["可樂"] = request.map2["可樂"]
        res.map2["牛奶"] = request.map2["牛奶"]
        res.map2["咖啡"] = request.map2["咖啡"]

        return res


    def TypeOneof(self, request, context) -> Hello_pb2.OneofType:
        res = Hello_pb2.OneofType()
        res.text = request.text
        res.num = request.num

        return res


    ## 以下測試連線類型

    def SayHello0(self, request: Hello_pb2.RequestType, context) -> Hello_pb2.ResponseType:


        res = Hello_pb2.ResponseType()

        res.sentence = f"{request.name} say {request.content}"

        return res

    def SayHello1(self, request_iterator: Generator[Hello_pb2.RequestType, None, None],
                  context) -> Hello_pb2.ResponseType:
        sentence = "'''\n"
        for each in request_iterator:
            content = f"{each.name} say {each.content}\n"
            sentence += content
        sentence = sentence + "'''"

        res = Hello_pb2.ResponseType()
        res.sentence = sentence

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
    Hello_pb2_grpc.add_TestServicer_to_server(TestExample(), server)  ## add stub to server , deserialize > handle > serialize
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
