BINARY_NAME=simplelog

build:
	go build -o ${BINARY_NAME} simplelog.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test

vet:
	go vet
	
lint:
	golangci-lint run --enable-all