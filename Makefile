init:
	chmod +x ./init.sh ./add.sh && ./init.sh

mock:
	mockery --all --keeptree

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/api/main.go

run:
	docker-compose up -d

protoc: protobuf/skeleton.proto
	protoc --go_out=protobuf protobuf/skeleton.proto
	protoc --go-grpc_out=protobuf protobuf/skeleton.proto

run_grpc:
	go run cmd/grpc/main.go

grpcui:
	grpcui -plaintext localhost:50051

.PHONY: init mock run run_grpc grpcui