.PHONY: clean all

BINARY_NAME := main

.PHONY: test
test:
	@go test -race -v ./...

.PHONY: build-pb
build-pb:
	@protoc --proto_path=grpc --go_out=plugins=grpc:grpc grpc/*.proto

.PHONY: build
build:
	@make build-pb
	@go build -v -o xtest-bin cmd/main.go

.PHONY: run
run:
	@make build
	./xtest-bin --http-address 0.0.0.0:9000 --grpc-address 0.0.0.0:50051  --log-level debug \
		   --redis-address 127.0.0.1:6379 \
		   --postgres-dsn postgres://postgres:@127.0.0.1/postgres?sslmode=disable \
		   --nsqd-address 127.0.0.1:4150
