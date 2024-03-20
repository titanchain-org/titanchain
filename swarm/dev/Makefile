.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/titanchain/titanchain/cmd/swarm
	go build -o bin/titan     github.com/titanchain/titanchain/cmd/titan
	go build -o bin/bootnode github.com/titanchain/titanchain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/titanchain/titanchain/swarm/...
