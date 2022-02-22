BINARY_NAME=semanticsort

build:
	go build -o ${BINARY_NAME} semanticsort.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test

vet:
	go vet
	
lint:
	golangci-lint run --enable-all