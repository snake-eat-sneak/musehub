module UserServiceAPI

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	musehub.xyz/UserServiceSRV v0.0.0-incompatible
)
replace (
    musehub.xyz/UserServiceSRV => ../UserServiceSRV
)
