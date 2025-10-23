.PHONY: all L2.13 test diff_test lint

all : L2.13

L2.13:
	@go build -o cut ./cmd/cut/main.go

test:
	go test ./internal/input -v

diff_test:
	@bash diff_tests.sh

lint:
	golangci-lint run ./...
