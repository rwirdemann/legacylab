build:
	go build -o ${GOPATH}/bin/lb frequency/main.go

.PHONY: build deploy-dev clean