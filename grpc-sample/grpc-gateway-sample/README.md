### grpc-gateway-sample

同一个端口提供 grpc 和 https 服务

### 效果

```sh
### 编译并启动服务
# ronething @ ashings-macbook-pro in ~/Documents/grpc-sample/grpc-gateway-sample on git:master x [15:27:50]
$ go build

# ronething @ ashings-macbook-pro in ~/Documents/grpc-sample/grpc-gateway-sample on git:master x [15:33:28]
$ ./grpc-gateway-sample server
2020/07/12 15:33:30 gRPC and https listen on: 50052

### 测试
## grpc
# ronething @ ashings-macbook-pro in ~/Documents/grpc-sample/grpc-gateway-sample/client on git:master x [15:33:33]
$ go run client.go
2020/07/12 15:33:39 hello, Grpc

## http
# ronething @ ashings-macbook-pro in ~/Documents/grpc-sample/grpc-gateway-sample/client on git:master x [15:33:39]
$ curl -X POST -k https://localhost:50052/hello_world -d '{"referer": "restful_api"}'
{"message":"hello, restful_api"}%
```

### 注意事项

如果终端启动了本地代理，记得要关掉，不然会出现奇奇怪怪的错误

```sh
{
    "error": "connection error: desc = \"transport: Error while dialing failed to do connect handshake, response: \\\"HTTP/1.1 400 Invalid header received from client\\\\r\\\\nConnection: close\\\\r\\\\nContent-Type: text/plain\\\\r\\\\n\\\\r\\\\nInvalid header received from client.\\\\r\\\\n\\\"\"",
    "code": 14,
    "message": "connection error: desc = \"transport: Error while dialing failed to do connect handshake, response: \\\"HTTP/1.1 400 Invalid header received from client\\\\r\\\\nConnection: close\\\\r\\\\nContent-Type: text/plain\\\\r\\\\n\\\\r\\\\nInvalid header received from client.\\\\r\\\\n\\\"\""
}
```

### 参考资料

https://eddycjy.gitbook.io/golang/di-5-ke-grpcgateway/hello-world
