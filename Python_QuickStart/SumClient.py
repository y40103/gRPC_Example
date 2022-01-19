import grpc
import Sum_pb2_grpc,Sum_pb2



host = grpc.insecure_channel("localhost:50051")

client = Sum_pb2_grpc.MyServiceStub(host) # add client stub

request = Sum_pb2.RequestForSum()
request.num1 = 100.5
request.num2 = 200.5

res = client.Sum(request)

print(f"num1 = {request.num1}, num2 = {request.num2}")
print(f"Sum = {res.sum_num}")
