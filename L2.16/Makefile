.PHONY: all L2.16 test lint

all : L2.16

L2.16:
	@go build -o wget ./cmd/wget/main.go

test:
	@go test ./internal/config -v
	@go test ./internal/downloader -v
	@go test ./internal/input -v   
	@go test ./internal/utils -v  

lint:
	golangci-lint run ./...