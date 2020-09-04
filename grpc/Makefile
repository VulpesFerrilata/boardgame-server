
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. protoc/auth/auth.proto
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. protoc/user/user.proto