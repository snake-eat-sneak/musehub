@echo off

start "user_srv_prod1" go run UserServiceSRV --registry=etcd --registry_address=localhost:2379 --server_address=localhost:8011 &
start "user_srv_prod2" go run UserServiceSRV --registry=etcd --registry_address=localhost:2379 --server_address=localhost:8012

pause