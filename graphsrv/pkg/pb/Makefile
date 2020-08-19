
build:
	protoc -I . -I $(GOPATH)/src --go_out=plugins=grpc:. ./graphsrv.proto

.PHONY: build