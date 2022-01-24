package main

import (
	"Hello/Hello/Hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func TestSimpleClinet(client Hello.TestClient,int_num int32,float_num float32,description string,bool_flag bool) error {
	request := &Hello.SimpleType{IntNum: int_num,FloatNum: float_num,Descprition: description,BoolFlag: bool_flag}
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*time.Duration(5))
	defer cancel()

	res,err := client.TypeSimple(ctx,request)

	if err != nil {
		log.Fatalf("fail to get response from server\n")
	}
	fmt.Printf("####################\n")
	fmt.Printf("%v %T\n",res.IntNum,res.IntNum)
	fmt.Printf("%v %T\n",res.FloatNum,res.FloatNum)
	fmt.Printf("%v %T\n",res.Descprition,res.Descprition)
	fmt.Printf("%t %T\n",res.BoolFlag,res.BoolFlag)

	return nil
}


func TestEnumeration(client Hello.TestClient){
	// golang 本身無枚舉類型  go grpc 枚舉 grpc底層實現主要是依靠 兩個 map , key:value and value:key


	request := &Hello.Enumeration{}
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()
	_,err := client.TypeEnum(ctx,request)
	if err != nil {
		log.Fatalf("can't get response from server\n")
	}
	fmt.Printf("%v %T\n",Hello.EnumerationPickOne_name,Hello.EnumerationPickOne_name)
	fmt.Printf("%v %T\n",Hello.EnumerationPickOne_value,Hello.EnumerationPickOne_value)
	fmt.Printf("\n")
	fmt.Printf("%v\n",Hello.Enumeration_CHOICE0)
	fmt.Printf("%v\n",Hello.EnumerationPickOne_name[0])
	fmt.Printf("%v\n",Hello.EnumerationPickOne(0))

	fmt.Printf("%v\n",Hello.EnumerationPickOne_value["CHOICE0"])
	fmt.Printf("%v\n",Hello.EnumerationPickOne.Number(0))

}


func TestListType(client Hello.TestClient,List ...int32){ // 表示 多個數入 但變 []int32
	request := &Hello.ListType{}

	request.IntList = List

	request.IntList = append(request.IntList,99999)

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*5)
	defer cancel()

	res,err := client.TypeList(ctx,request)

	if err != nil {
		fmt.Printf("can't get response from server %v\n",err)
	}

	fmt.Printf("%v\n",res.IntList)

	for index,val := range res.IntList {
		fmt.Printf("index %v val %v\n",index,val)
	}

}

func TestNestedType(client Hello.TestClient,name string,content string){
	// golang grpc 不會產生 struct nested struct , 則是會創建兩個 struct , nested另一個會叫 A_B

	ctx,cancel := context.WithTimeout(context.Background(),time.Second*5)

	defer cancel()



	request := &Hello.NestedType_RequestType{Name: name,Content: content}
	res,err := client.TypeNested(ctx,request)

	if err != nil {
		log.Fatalf("can't connect to server %v\n",err)
	}

	fmt.Printf("%v不禁留下兩行機油 %v\n",res.Name,res.Content)

}


func TestMap(client Hello.TestClient,food []string,price []int32) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	request := &Hello.MapType{}
	request.Map1 = make(map[string]int32)
	request.Map2 = make(map[string]int32)

	for i:=0;i<len(food)/2;i++{
		request.Map1[food[i]] = price[i]
	}
	for i:=len(food)/2;i<len(food);i++{
		request.Map2[food[i]] = price[i]
	}

	res,err := client.TypeMap(ctx,request)

	if err != nil {
		fmt.Printf("can't get response from server\n",res)
	}

	for key,val := range res.Map1 {
		fmt.Printf("%v %v\n",key,val)
	}

	for key,val := range res.Map2 {
		fmt.Printf("%v %v\n",key,val)
	}


}


func TestOneof(client Hello.TestClient, request *Hello.OneofType) {

	//ctx,cancel := context.WithTimeout(context.Background(),time.Second * 5)
	//defer cancel()

	switch request.GetOne().(type) {
	case *Hello.OneofType_Num:
		fmt.Printf("%v\n",request.GetNum())
	case *Hello.OneofType_Text:
		fmt.Printf("%v\n",request.GetText())
	}



}





func main(){

	server_host := "localhost:50051"
	conn,err:= grpc.Dial(server_host,grpc.WithInsecure())
	defer conn.Close()
	if err!= nil{
		log.Fatalf("fail to coonect server\n")
	}

	client := Hello.NewTestClient(conn)
	fmt.Printf("SimpleType 		####################\n")
	TestSimpleClinet(client,100,3.14,"testing",true)
	fmt.Printf("EnumerationType 	####################\n")
	TestEnumeration(client)
	fmt.Printf("ListType 		####################\n")
	intslice := []int32{100,200,300,400}
	TestListType(client,intslice...) // 表示一個進去變多個   會變 slice
	fmt.Printf("NestedType 		####################\n")
	TestNestedType(client,"機器人","哭R")
	fmt.Printf("MapType 			####################\n")
	food := []string{"巧克力","蛋糕","香蕉","布朗尼","布丁","芒果"}
	price := []int32{100,200,50,130,100,90}
	TestMap(client,food,price)
	fmt.Printf("OneofType 		####################\n")
	//request := &Hello.OneofType{One: &Hello.OneofType_Num{Num: 111}}
	request := &Hello.OneofType{One: &Hello.OneofType_Text{Text: "book"}}
	TestOneof(client,request) // oneof == python中 type hint 中 Union[int,str]

}



