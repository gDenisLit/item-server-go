build: 
	@go build -o ./bin/item-server

run: build
	@./bin/item-server

test: 
	@go test -v ./..

install:
	@go mod tidy