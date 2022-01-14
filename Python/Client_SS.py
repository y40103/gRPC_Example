import grpc
import Hello_pb2
import Hello_pb2_grpc

host = grpc.insecure_channel("localhost:50051")

client = Hello_pb2_grpc.TestStub(host)

request = Hello_pb2.RequestType()

request.name = "皮卡丘"
request.content = "屁卡 "

res = client.SayHello2(request)

for each_res in res:
    print(each_res.sentence)
