dev:
	go run cmd/api/main.go
build:
	go build -o ./bin/api cmd/api/main.go
test:
	go test -cover ./...
