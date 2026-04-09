
build:
	@go clean && @go build -o build/gunter cmd/main.go
run:
	@go run cmd/main.go
test:
	@go test ./... -v -cover
