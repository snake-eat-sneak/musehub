module UserServiceAPI

go 1.14

require (
	github.com/favadi/protoc-go-inject-tag v1.1.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	musehub.xyz/proto v0.0.0-incompatible
)

replace musehub.xyz/proto => ../proto
