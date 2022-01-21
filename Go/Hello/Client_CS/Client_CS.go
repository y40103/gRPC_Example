package main

import (
	"Hello/Hello/Hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func ClientCS(client Hello.TestClient, name string, content string) *Hello.ResponseType {

	// client 為接口 , 下面有 SayHello1(ctx context.Context, opts ...grpc.CallOption) (Test_SayHello1Client, error)
	// 非直接調用 HelloMethod的SayHello1

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	stream, err := client.SayHello1(ctx) // 可理解成 開啟一個對server SayHello1 函數的 串接器
	//  參數輸入 為 直接對至 Server SayHello1 函數 ,
	//  該 串接器 回傳的 stream 也對應至 server SayHello1 函數內部 stream
	//  stream 為 channel 概念的東西 client,server共用一個queue

	if err != nil {
		log.Fatalf("fail to get stream %v\n", err)
	}

	var full_echo string

	// streaming 5次 (同python迭代器 展開) 給 server
	for i := 0; i < 5; i++ {

		character := fmt.Sprintf("%v%v號", name, i+1)
		full_echo += content

		err := stream.Send(&Hello.RequestType{Name: character, Content: full_echo})
		// 會將 &Hello.RequestType 傳至 Server端的 SayHello1 並返回結果

		if err != nil {
			fmt.Printf("send error %v\n", err)
			continue
		}
	}

	resp, err := stream.CloseAndRecv() // 將 streaming 結束後 , 函數處理完的結果回傳

	if err != nil {
		log.Fatalf("fail to closeAndRecv: %v\n", err)
	}

	return resp

}

func main() {

	serverhost := "localhost:50051"
	conn, err := grpc.Dial(serverhost, grpc.WithInsecure()) // conn 為一個 連線狀態的結構體 包含一些方法 , connect , close ...
	// 連線至 server
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := Hello.NewTestClient(conn) // 連線狀態掛上 Test service 的 stub

	res := ClientCS(client, "小火龍", "嘎嘎 ")

	fmt.Printf("```\n")
	fmt.Printf("%v", res.GetSentence())
	fmt.Printf("```\n")

}
