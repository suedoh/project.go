build:
	@go build -o bin/project

run: build
	@./bin/project

test:
	@go test -v ./...
