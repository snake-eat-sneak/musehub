@echo off

go run UserServiceSRV --registry=etcd --registry_address=localhost:2379

pause