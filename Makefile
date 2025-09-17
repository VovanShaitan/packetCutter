.PHONY: run test build clean

run:
	go run cmd/matcher/main.go

test:
	go test ./internal/... -v


build:
	go build -o bin/matcher cmd/matcher/main.go

clean:
	rm -rf bin/
	go clean
	