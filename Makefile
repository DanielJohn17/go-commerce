build:
	@go build -o bin/go-commerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-commerce