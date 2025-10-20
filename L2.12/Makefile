.PHONY: all L2.12 test diff_test lint

all : L2.12

L2.12:
	@go build -o grep ./cmd/grep/main.go

test:
	go test ./internal/buffer -v
	go test ./internal/grep -v

diff_test:
	@bash diff_tests.sh

lint:
	golangci-lint run ./...
