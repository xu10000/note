支持rpc请求 
使用grpc-gateway支持http请求（将http转为grpc):  http://localhost:8080/example/hello?name=1

编译micro和pb文件：
protoc  -I/usr/local/include \-I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --micro_out=.  --go_out=plugins=grpc:. ./proto/*.proto 

编译网关gw文件：
protoc  -I/usr/local/include \-I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis     --grpc-gateway_out=logtostderr=true:. ./proto/*.proto 

生成gw和pb文件registerxx函数名相同会冲突，修改成不同：
%s/RegisterXXX/RegisterGWxxx/g

执行： go run rpc.go



