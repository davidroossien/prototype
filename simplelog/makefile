BINARY_NAME=simplelog

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

ifndef $(GOBIN)
    GOBIN=$(shell go env GOBIN)
    export GOBIN
endif

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

deploy:
	cp ${BINARY_NAME} ${GOBIN}/${BINARY_NAME}
