@echo off

start "prod1" go run gateway --server_address :8001 &
start "prod2" go run gateway --server_address :8002 &
start "prod3" go run gateway --server_address :8003

pause