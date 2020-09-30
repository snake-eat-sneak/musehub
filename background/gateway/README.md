网关服务

在本目录下执行命令行

go mod init gateway (若存在go.mod文件，则不需要执行此命令)
go build gateway
./gateway --server_address :8001

启动网关

或者本目录下有一个service_start.bat文件，可直接启动3个服务注册到consul中，端口分别为8001 8002 8003