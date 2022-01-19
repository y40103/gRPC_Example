import Hello_pb2
from google.protobuf.json_format import MessageToJson,MessageToDict


## 隨機找兩個 protobuf 格式進行轉換


request1 = Hello_pb2.RequestType()  ## 把想發出去的 request參數 封裝
request1.name = "傑尼龜"
request1.content = "J ni J ni"

request2 = Hello_pb2.MapType()
request2.map1["巧克力"] = 20
request2.map1["泡芙"] = 30
request2.map1["海苔"] = 50
request2.map2["可樂"] = 30
request2.map2["牛奶"] = 80
request2.map2["咖啡"] = 50



## to json

json_type1 = MessageToJson(request1)
json_type2 = MessageToJson(request2)

print(json_type1)
print(json_type2)



## to map


map_type1 = MessageToDict(request1)
map_type2 = MessageToDict(request2)

print(map_type1)
print(map_type2)
