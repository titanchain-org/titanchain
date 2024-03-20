.PHONY: titan titan-cross evm all test clean
.PHONY: titan-linux titan-linux-386 titan-linux-amd64 titan-linux-mips64 titan-linux-mips64le
.PHONY: titan-darwin titan-darwin-386 titan-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.13.1
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

titan:
	go run build/ci.go install ./cmd/titan
	@echo "Done building."
	@echo "Run \"$(GOBIN)/titan\" to launch titan."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

titan-cross: titan-windows-amd64 titan-darwin-amd64 titan-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/titan-*

titan-linux: titan-linux-386 titan-linux-amd64 titan-linux-mips64 titan-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-*

titan-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/titan
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep 386

titan-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/titan
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep amd64

titan-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/titan
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep mips

titan-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/titan
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep mipsle

titan-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/titan
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep mips64

titan-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/titan
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/titan-linux-* | grep mips64le

titan-darwin: titan-darwin-386 titan-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/titan-darwin-*

titan-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/titan
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/titan-darwin-* | grep 386

titan-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/titan
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/titan-darwin-* | grep amd64

titan-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) -buildmode=mode -x --targets=windows/amd64 -v ./cmd/titan
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/titan-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
