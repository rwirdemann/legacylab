build:
	go build -o ${GOPATH}/bin/hs cli/hotspot/main.go
	env GOOS=linux CGO_ENABLED=0 go build ${LDFLAGS} -a -installsuffix cgo -o bin/legacylab http/main.go

deploy: build
	# ssh 95.217.222.60 "pkill legacylab"
	scp bin/legacylab 95.217.222.60:~/legacylab/bin
	ssh 95.217.222.60 "nohup /home/ralf/legacylab/bin/legacylab &"

test:
	go test ./... 

.PHONY: build test