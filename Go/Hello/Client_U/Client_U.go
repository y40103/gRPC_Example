package main

import (
	"Hello/Hello/Hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func ClientU(client Hello.TestClient, name string, content string) (response *Hello.ResponseType) {

	request := &Hello.RequestType{Name: name, Content: content} // 封裝成 protobuf 格式

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 十秒後 啟動withdeadline 結束該function

	defer cancel() // 函數結束 斷掉 context

	response, err := client.SayHello0(ctx, request) // 啟動server 的方法

	if err != nil {
		log.Fatalf("%v.get SayHello0(_) = _ , %v: ", client, err)
	}

	return response

}

func main() {

	serverhost := "localhost:50051"
	conn, err := grpc.Dial(serverhost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := Hello.NewTestClient(conn)

	// 連線設定




	res := ClientU(client, "傑尼龜", "J ni J ni")

	fmt.Printf("%v",res.Sentence)

}
