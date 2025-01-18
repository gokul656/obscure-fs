run:
	go run cmd/main.go 5001

build:
	GOOS=darwin go build -o bin/node cmd/main.go
