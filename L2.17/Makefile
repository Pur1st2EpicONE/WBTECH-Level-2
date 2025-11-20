.PHONY: all L2.17 clean test lint server stop

all: L2.17

L2.17:
	@go build -o telnet ./cmd/telnet/main.go

clean: stop
	@rm -f ./telnet

test: stop
	@go test ./internal/input -v
	@go test ./internal/reader -v
	@go test ./internal/writer -v   

server:
	@bash server.sh & echo $$! > server.pid

stop:
	@if [ -f server.pid ]; then \
		kill $$(cat server.pid) 2>/dev/null || true; \
		rm -f server.pid; \
	fi
	@pids=$$(ps aux | grep '[n]c -l -p 8080' | awk '{print $$2}'); \
	if [ -n "$$pids" ]; then \
		echo "$$pids" | xargs kill; \
		echo "Echo server stopped"; \
	fi

lint:
	golangci-lint run ./...
