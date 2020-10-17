@echo off

start "prod1" go run UserServiceSRV --server_address :8101 &
start "prod2" go run UserServiceSRV --server_address :8102 &
start "prod3" go run UserServiceSRV --server_address :8103

pause