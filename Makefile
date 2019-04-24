.PHONY: proto server client

proto:
	protoc --go_out=plugins=grpc,paths=source_relative:./model --proto_path=./model ./model/model.proto

server:
	go run ./server/server.go

client:
	go run ./client/client.go

