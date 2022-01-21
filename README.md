# gRPC_Example



gRPC 隨筆

[Notion](https://handy-lady-8da.notion.site/gRPC-55cde33b0e16430db6b587d3419fbcb5)

整理整理著 把 gRPC 所有連線形式與感覺比較常用的資料類型 用簡單的範例實做一次

內容累積下來感覺也不少 變得有點雜亂而不易閱讀

可以優先看QuickStart

裡面用一個最簡單的範例實現gRPC

可從該範例理解整個gRPC的架構

其他只是依需求 抽換需要的資料類型, 函數內容, 連線方式

Python
======


```bash
gRPC_Example/
└── Python
    ├── Client_BDS.py
    ├── Client_CS.py
    ├── Client_SS.py
    ├── Client_Types.py
    ├── Client_U.py
    ├── Hello_pb2_grpc.py
    ├── Hello_pb2.py
    ├── Hello.proto
    └── Server.py
    └───Python_Protobuf_Type
        ├── Hello_pb2_grpc.py
        ├── Hello_pb2.py
        └── protobuf_type.py

```


1.  install python gRPC package

```bash
pip install grpcio grpcio-tools
```


2. run gRPC Server

```bash
python Server.py

```

3. execute exmaples

  - about protobuf schema variable definition in python

  ```bash
  python Client_Types.py
  ```

  - about protbuf schema types of api in gRPC in python
  ```bash
  python Client_U.py        // Unary
  python Client_CS.py       // Client Streaming
  python Client_SS.py       // Server Streaming
  python Client_BDS.py      // Bi Directional Streaming
  ```
  
- convert protobuf type to json and map
 ```bash
 python protobuf_type.py
 ```

Golang
======
增加Go基本連線模式測試, 與python架構/功能

```bash
gRPC_Example/
├── Go
    ├── go.mod
    ├── go.sum
    └── Hello
        ├── Client_BDS
        │   └── Client_BDS.go
        ├── Client_CS
        │   └── Client_CS.go
        ├── Client_SS
        │   └── Clinet_SS.go
        ├── Client_U
        │   └── Client_U.go
        ├── grpc_server
        │   └── main.go
        └── Hello
            ├── Hello_grpc.pb.go
            ├── Hello.pb.go
            └── Hello.proto

```

1. Init
 - set project directory and package
 
    gRPC_Example/Go
    ```
    go mod init "Hello"
    ```
    ```bash
    go get google.golang.org/grpc
      ```




2. test go server <> python client

  - run Go server

      gRPC_Example/Go/Hello/grpc_server
      ```bash
      go run main.go

      ```

  - run server

      gRPC_Example/Python
      ```bash
      python Client_BDS.py
      ```




