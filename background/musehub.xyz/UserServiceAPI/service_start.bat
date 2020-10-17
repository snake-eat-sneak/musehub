@echo off

go run UserServiceAPI --registry=etcd --registry_address=127.0.0.1:2379

pause