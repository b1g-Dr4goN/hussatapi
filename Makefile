build:
	@go build -o bin/hussatapi cmd/main.go

test:
	@go test -v ./..

run: build
	@./bin/hussatapi