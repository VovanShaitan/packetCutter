.PHONY: run test build clean

run:
	go run cmd/cutter/main.go

test:
	go test ./internal/... -v -race -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	go build -o bin/packetcutter cmd/cutter/main.go

clean:
	rm -rf bin/ coverage.out coverage.html
	go clean
	