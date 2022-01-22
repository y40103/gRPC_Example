package main

import (
	"Hello/Hello/Hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)



type HelloMethod struct {
	Hello.UnimplementedTestServer
}

func (self *HelloMethod) TypeSimple(ctx context.Context,request *Hello.SimpleType) (*Hello.SimpleType,error) {

	res := &Hello.SimpleType{}  // 輸入輸出為 結構體

	res.FloatNum = request.FloatNum
	res.BoolFlag = request.BoolFlag
	res.Descprition = request.Descprition
	res.IntNum = request.IntNum

	return res,nil

}



func (self *HelloMethod) TypeEnum(ctx context.Context,request *Hello.Enumeration) (*Hello.Enumeration,error) {

	fmt.Printf("%v\n",Hello.Enumeration_CHOICE0)
	fmt.Printf("%v\n",Hello.EnumerationPickOne_name)
	fmt.Printf("%v\n",Hello.EnumerationPickOne_value)
	fmt.Printf("%v\n",Hello.EnumerationPickOne(1))
	fmt.Printf("%v\n",Hello.EnumerationPickOne.Number(0))

	return request,nil

}

func (self *HelloMethod) TypeList(ctx context.Context,request *Hello.ListType) (*Hello.ListType,error) {

	res := &Hello.ListType{}
	res.IntList = request.IntList

	return res,nil


}


func (self *HelloMethod) TypeNested(ctx context.Context,request *Hello.NestedType_RequestType) (*Hello.NestedType_RequestType,error) {

	res := &Hello.NestedType_RequestType{}
	res = request

	return res,nil


}


func (self *HelloMethod) TypeMap(ctx context.Context,request *Hello.MapType) (*Hello.MapType,error) {

	res := &Hello.MapType{}
	res = request
	return res,nil
}






func (self *HelloMethod) SayHello0 ( ctx context.Context, request *Hello.RequestType ) ( *Hello.ResponseType,error) {
	fmt.Printf("SayHello0...\n")

	//name := request.GetName()
	name := request.GetName()
	content := request.GetContent()

	res := Hello.ResponseType{
		Sentence: name + " say "+ content,
	}
	return &res,nil
}



func (self *HelloMethod) SayHello1 (reqStream Hello.Test_SayHello1Server) error {
	var res Hello.ResponseType

	for {
		request , err := reqStream.Recv()

		if err == io.EOF {

			fmt.Printf("finish streaming...\n")

			return reqStream.SendAndClose(&res)
		}

		temp := fmt.Sprintf("%v say %v\n",request.Name,request.Content)
		res.Sentence += temp


	}

}


func (self *HelloMethod) SayHello2 (request *Hello.RequestType,resStream Hello.Test_SayHello2Server) error {
	var name,content string

	var res = Hello.ResponseType{}

	for i:=0;i<5;i++ {

		name = request.Name + fmt.Sprintf("%v號",i+1)
		content += request.Content
		res.Sentence = fmt.Sprintf("%v say %v",name,content)
		resStream.Send(&res)

	}

	return nil
}

func multiContent(content string, count int) string {
	// 增加content次數

	var res = ""
	for i:=0;i<count;i++{
		res += content
	}
	return res
}





func (self *HelloMethod) SayHello3(resStream Hello.Test_SayHello3Server) error {

	res := &Hello.ResponseType{}

	count:= 1

	for {
		request,err := resStream.Recv()

		if err == io.EOF {
			fmt.Printf("BDS server finish reciecve request from client \n")
			break
		} else if err != nil {
			fmt.Printf("BDS server fail to receive request %v\n",err)
			break
		}

		request.Content = multiContent(request.Content,count)
		count +=1
		// 增加content次數


		res.Sentence = fmt.Sprintf("%v say %v",request.Name,request.Content)

		resStream.Send(res)
	}

	return nil

}


// 流程

// clinet端
// 1. 建立一個 連線特定host結構體
// 2. 將 該連線 掛上 Test service 的 stub
// 3. 開通該連線與stub某函數的接口
// 4. 利用該接口的方法 , 第一個函數為 context 控制連線開關,
//    U 第二個函數為 protobuf request類型 ,  直接回傳 protobuf response
//    CS 輸入若為串流 則無第二個參數 返回 一個 stream類型 , 利用.send()方法把 protobuf reqest多次傳入 , 最後用 .CloseAndRecv 把 protobuf response回傳

func main(){
	fmt.Printf("start gRPC server ... \n")
	lis,err := net.Listen("tcp","127.0.0.1:50051")
	if err != nil {
		log.Fatalf("fail to listen %v\n",err)
	}

	gRCPServer := grpc.NewServer()

	Hello.RegisterTestServer(gRCPServer,&HelloMethod{})

	if err := gRCPServer.Serve(lis); err != nil {
		log.Fatalf("fail to serve %v\n",err)
	}
}
