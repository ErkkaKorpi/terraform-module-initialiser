BINARY_NAME=tfminit

test:
	go test

build: test
	go build -o ${BINARY_NAME}

build-all:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}.exe

PHONY: test