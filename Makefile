build:
	@go build -o ./bin/golang-key-value-db ./cmd/main.go

test:
	@go test -v ./...

coverage:
	@go test -cover fmt

run: build
	@./bin/golang-key-value-db