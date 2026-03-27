.PHONY: all L2.10 test diff_test lint clean

all : L2.10

L2.10:
	@go build -o sort ./cmd/sort/main.go

test:
	go test ./internal/sorter -v
	go test ./internal/comparator -v

diff_test:
	@bash diff_tests.sh

lint:
	golangci-lint run ./...

clean:
	@rm -f sort