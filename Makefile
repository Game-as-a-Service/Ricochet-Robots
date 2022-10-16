init:
	chmod 755 ./init.sh ./add.sh && ./init.sh

mock:
	mockery --all --keeptree

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/api/main.go

run:
	docker-compose up -d

.PHONY: init mock run