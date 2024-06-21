.PHONY: all transmitter_proto transmitter_response transmitter

all:

transmitter:
	go run cmd/transmitter/main.go

transmitter_proto:
	protoc --go_out=. --go-grpc_out=. transmitter.proto

transmitter_response:
	grpcurl -plaintext -d '{}' localhost:50051 transmitter.TransmitterService/StreamData

client:
	go run cmd/client/main.go


