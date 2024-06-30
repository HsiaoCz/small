run: build
	@./bin/small

build:
	@go build -o bin/small main.go

test:
	@go test -v ./...

tidy:
	@go mod tidy
