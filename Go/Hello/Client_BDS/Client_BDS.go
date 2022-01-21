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

var intchan = make(chan int,20)
var exitchan = make(chan int,10)

func BSDReceive(stream Hello.Test_SayHello3Client) {

	for {
		<- intchan
		res,err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("BDS client receive finish...\n")
			break
		} else if err != nil {
			log.Fatalf("BDS fail to receive from server\n")
		}

		fmt.Printf("%v\n",res.GetSentence())
		exitchan <- 1

	}

	close(exitchan)


}





func BDSSend(stream Hello.Test_SayHello3Client,IP []string,echo []string){

	var request = Hello.RequestType{}

	for index:=0; index<len(IP);index++ {

		request.Name = IP[index]
		request.Content = echo[index]

		err := stream.Send(&request)
		intchan <- 1
		if err != nil {
			log.Fatalf("fail to send request to server %v",err)
		}


	}

	close(intchan)

	stream.CloseSend()
	// 非常重要 , 全部結束後 需將 stream 關閉 , 這樣 stream send最後才會是 EOF
	//, 否則 stream為空 還送出, 讀取端 err 則會自動產出 cantext caancel


}


func ExitDetector() {

	for {

		_,ok := <- exitchan

		if ok == false {

			break

		}

	}

	fmt.Printf("finish main process\n")



}



func ClientBDS(client Hello.TestClient,IP []string,echo []string) {

	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)

	BSDstream,err := client.SayHello3(ctx)

	if err != nil {
		log.Fatalf("can't get stream: %v %v",client,err)
	}

	go BSDReceive(BSDstream)
	go BDSSend(BSDstream,IP,echo)

}




func main() {

	var IP = []string{"卡比獸", "嘎啦嘎啦", "多啦a夢", "約得爾賤畜", "嚕嚕咪"}
	var echo = []string{"ZZ ", "嘎啦 ", "大雄啊 ", "1234 ", "@&^*$ "}

	server_host := "localhost:50051"
	conn,err := grpc.Dial(server_host,grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("fail to dial: %v\n",err)
	}


	client := Hello.NewTestClient(conn)

	ClientBDS(client,IP,echo)

	ExitDetector()




}
