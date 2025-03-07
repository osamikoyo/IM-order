CC = go
OUTPUT = bin/app

proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/order.proto 
deps:
	go mod download
build:
	go build -o $(OUTPUT) cmd/main.go
run:
	$(MAKE) build
	$(OUTPUT)