# gRPC_Example



gRPC 隨筆

[Notion](https://handy-lady-8da.notion.site/gRPC-55cde33b0e16430db6b587d3419fbcb5)

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


1.  install python gRpc package

```bash
pip install grpcio grpcio-tools
```


2. activate gRPC Server

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

