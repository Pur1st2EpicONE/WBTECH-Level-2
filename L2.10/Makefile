.PHONY: all L2.10 test lint

all : L2.10

L2.10:
	@go run main.go

test:
	@bash script.sh

lint:
	golangci-lint run ./...
