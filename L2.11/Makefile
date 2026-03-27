.PHONY: all L2.11 test lint

all : L2.11

L2.11:
	@go run main.go

test:
	go test -cover ./...

lint:
	golangci-lint run ./...
