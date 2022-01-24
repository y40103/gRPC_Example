import grpc
import Hello_pb2
import Hello_pb2_grpc

HOST = grpc.insecure_channel("localhost:50051")
CLIENT = Hello_pb2_grpc.TestStub(HOST)

def Test_Simple():
    request = Hello_pb2.SimpleType()  ## 把想發出去的 request參數 封裝
    request.int_num = 10
    request.float_num = 10.2
    request.description = "testing"
    res = CLIENT.TypeSimple(request)

    print(res.int_num,type(res.int_num))
    print(res.float_num,type(res.float_num))
    print(res.description,type(res.description))
    print(res.bool_flag,type(res.bool_flag))

def Test_Enum():
    enum = Hello_pb2.Enumeration()
    res = CLIENT.TypeEnum(enum)
    print(res.CHOICE0)
    print(res.pick_one.CHOICE0)
    print(res.pick_one.Value("CHOICE0"))
    # 上三者 相同 key:value
    print(res.pick_one.Name(1)) ## value:key





def Test_List():
    request = Hello_pb2.ListType()  ## 把想發出去的 request參數 封裝
    request.int_list.append(1)
    request.int_list.append(2)
    request.int_list.append(3)
    res = CLIENT.TypeList(request).int_list
    print(res,type(res))
    print(len(res))
    print(res[0])


def Test_Nested():
    request = Hello_pb2.NestedType()
    request.RequestType.name = "大毛"
    request.RequestType.content = "睡覺"
    res = CLIENT.TypeNested(request)
    print(res.RequestType.name, type(res.RequestType.name))
    print(res.RequestType.content, type(res.RequestType.content))

def Test_Map():
    request = Hello_pb2.MapType()
    request.map1["巧克力"] = 20
    request.map1["泡芙"] = 30
    request.map1["海苔"] = 50
    request.map2["可樂"] = 30
    request.map2["牛奶"] = 80
    request.map2["咖啡"] = 50
    res = CLIENT.TypeMap(request)

    del res.map2["咖啡"]  ## 刪除某個 key

    for key in res.map1:
        print(key,res.map1[key])
    for key in res.map2:
        print(key, res.map2[key])

    print(type(res.map1))

def Test_Oneof():
    request = Hello_pb2.OneofType() ## 該類型 實際上等同 type hint 中 Union[int,str]
    request.text = "二"
    request.num = 2   # 這個類型 只能給其中之一賦值 最後一個有效 request 實際只有 num有賦值, text為空
    res = CLIENT.TypeOneof(request)

    print(res.text) ## 為空
    print(res.num)

if __name__ == "__main__":
    print("Simple Type   #########################")
    Test_Simple()
    print("Enum Type     #########################")
    Test_Enum()
    print("List Type     #########################")
    Test_List()
    print("Nested Type   #########################")
    Test_Nested()
    print("Map Type      #########################")
    Test_Map()
    print("Oneof Type    #########################")
    Test_Oneof()