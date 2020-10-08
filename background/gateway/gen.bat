cd models/protos
protoc --micro_out=../ --go_out=../ gateway.proto
cd .. && cd ..