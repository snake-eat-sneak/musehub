cd proto/user
protoc --go_out=../ models.proto
protoc --micro_out=../ --go_out=../ user.proto
protoc-go-inject-tag -input=../models.pb.go
protoc-go-inject-tag -input=../user.pb.go
cd .. && cd ..