build:
	go build -o ${GOPATH}/bin/lb main.go

test:
	go test ./... 

.PHONY: build test