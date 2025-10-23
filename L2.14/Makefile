.PHONY: all L2.14 test lint

all : L2.14

L2.14:
	@go run main.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...
