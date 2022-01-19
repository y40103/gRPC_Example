import Sum_pb2
import Sum_pb2_grpc
import grpc
from concurrent import futures
import time
import logging


class QuickStart(Sum_pb2_grpc.MyServiceServicer):
    """ 該class 需包含Service所有方法"""

    def Sum(self, request, context):  # request 為 client 輸入的參數 , 在server 端輸入時為 protobuf 類型

        res = Sum_pb2.ResponseForSum()  # 實例 protonuf 類型 的 instance  回傳必須為此格式
        res.sum_num = request.num1 + request.num2  # 求和函數 內容

        return res


def gRPC_server():
    """ 運行server """

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))  # create server
    Sum_pb2_grpc.add_MyServiceServicer_to_server(QuickStart(), server)  # add QuickStaart Service Stub to server
    server.add_insecure_port("0.0.0.0:50051")  # add host to server
    server.start()  # server start
    try:
        while True:
            time.sleep(3600000000)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    gRPC_server()
