BINARY_NAME=app
ENTRYPOINT_NAME= cmd/webserver/main.go
build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} ${ENTRYPOINT_NAME}

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
