# Schedule

- Ngày học lí thuyết : 11/03/2024 -> 18/03/2024
- Ngày viết code example : 16/03/2024 -> 21/03/2024 
- [Báo cáo](https://docs.google.com/document/d/1ICQgr_o0y7lcN8kD0tMwNONszG5nbtXR-txgo3wp7oo)

# Structure 
```
└───example
    ├───example-crud
    └───udemy
        └───chatapplication
            ├───chat_grpc
            └───client

```
- example-curd: Chứa code triển khai mẫu cho CRUD với redis
- udemy/chatapplication : triển khai 1 ứng dụng chat thế giới sử dụng gRPC stream và Publish/Subcrible của Redis 

# Install 

### Requirement installed go and protoc

Clone repo and run :
```
go mod tidy
```

# Running

ChatApp: 
- Run server : 
```
go run ./example/udemy/chatapplication
```
- Run client : 
```
go run ./example/udemy/chatapplication/client
```