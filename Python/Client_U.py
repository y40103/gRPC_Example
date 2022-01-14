import grpc
import Hello_pb2
import Hello_pb2_grpc

host = grpc.insecure_channel("localhost:50051")

client = Hello_pb2_grpc.TestStub(host)

request = Hello_pb2.RequestType()  ## 把想發出去的 request參數 封裝
request.name = "傑尼龜"
request.content = "J ni J ni"

res = client.SayHello0(request).sentence

print(res)
