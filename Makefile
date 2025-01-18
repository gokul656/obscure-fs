run:
	go run main.go --port 5001

build:
	CGO_ENABLED=0 go build -o bin/node main.go
