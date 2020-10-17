用户API服务

gen.bat生成protobuf接口，在执行之前需要下载安装protoc：
1.github上下载一个cpp包：https://github.com/google/protobuf/releases  windows下载win版本，并解压，将bin目录配置到环境变量
  系统应用资源中也有，解压缩即可
2.protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
3.安装protoc-gen-micro
go get github.com/micro/protoc-gen-micro

service_start.bat启动服务

如果网关和SRV服务都启动了，http发送post请求到http://localhost:8080/user/UserLogin，body:{"user_id" : "","password" : ""},
会有返回值