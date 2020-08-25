# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gwan evm clean

GOBIN = build/bin

linuxDir=$(shell echo gwan-linux-amd64-`cat ./VERSION`-`git rev-parse --short=8 HEAD`)
windowsDir=$(shell echo gwan-windows-amd64-`cat ./VERSION`-`git rev-parse --short=8 HEAD`)
darwinDir=$(shell echo gwan-mac-amd64-`cat ./VERSION`-`git rev-parse --short=8 HEAD`)
# The gwan target build gwan binary
gwan:
	go build -v -o $(GOBIN)/gwan ./cmd/gwan
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gwan\" to launch gwan."

# The evm target build EVM emulator binary
evm:
	go build -v -o $(GOBIN)/evm  ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

# The clean target clear all the build output
clean:
	$(RM) $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen
