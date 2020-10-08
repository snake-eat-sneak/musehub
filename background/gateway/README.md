网关服务

在本目录下执行命令行

go mod init gateway (若存在go.mod文件，则不需要执行此命令)
go build gateway
./gateway --server_address :8001

启动网关

或者本目录下有一个service_start.bat文件，可直接启动3个服务注册到consul中，端口分别为8001 8002 8003


gen.bat生成protobuf接口，在执行之前需要下载安装protoc：
1.github上下载一个cpp包：https://github.com/google/protobuf/releases  windows下载win版本，并解压，将bin目录配置到环境变量
  系统应用资源中也有，解压缩即可
2.protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
3.安装protoc-gen-micro
go get github.com/micro/protoc-gen-micro