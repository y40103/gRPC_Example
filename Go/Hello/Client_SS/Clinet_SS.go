package main


import (
	"Hello/Hello/Hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func ClientSS(client Hello.TestClient,name string,content string) (stream Hello.Test_SayHello2Client, cancel context.CancelFunc) {

	request := Hello.RequestType{Name: name,Content: content}

	ctx,cancel:= context.WithTimeout(context.Background(),5*time.Second)

	stream, err := client.SayHello2(ctx,&request)
	//  client.SayHello2 可理解成 與 server SayHello2的串接器
	//  參數輸入 為 直接對至 Server SayHello1 函數 ,
	//  該 串接器 回傳的 stream 也對應至 server SayHello1 函數內部 stream
	//  stream 為 channel 概念的東西 client,server共用一個queue

	if err!= nil {
		log.Fatalf("fail to get stream")
	}


	return stream,cancel

}








func main(){
	server_hoat := "localhost:50051"
	conn,err := grpc.Dial(server_hoat,grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("fail to connect to server\n")
	}

	client := Hello.NewTestClient(conn)

	resStream,cancel := ClientSS(client,"皮卡丘","屁卡 ")




	// 展開 stream
	for {

		res,err := resStream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("fail to read response stream %v",err)
		}
		fmt.Printf("%v\n",res.Sentence)

	}

	cancel()

}
