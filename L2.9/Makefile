.PHONY: all L2.9 test lint

all : L2.9

L2.9:
	@go run main.go

test:
	go test ./...

lint:
	golangci-lint run ./...
