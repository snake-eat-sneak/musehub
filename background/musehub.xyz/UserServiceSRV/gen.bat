cd proto/user
protoc --go_out=../ models.proto
protoc --micro_out=../ --go_out=../ user.proto
cd .. && cd ..