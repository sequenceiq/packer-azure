#VERSION=v4
VERSION=$(shell git  describe --tags --abbrev=0)

build-local:
	go build -a -v -o ~/.packer.d/plugins/packer-builder-azure   ./packer/plugin/packer-builder-azure/
	go build -a -v -o ~/.packer.d/plugins/packer-builder-azure-arm ./packer/plugin/packer-builder-azure-arm/

release: build-local
	rm -rf release; mkdir release
	GOOS=linux CGO_ENABLED=0 go get -installsuffix cgo ./...
	tar -czvf release/packer.tgz  -C  $(GOPATH)/bin/linux_amd64 $(shell gfind ~/go/bin/linux_amd64 -name \*azure\* -printf "%P ")
	gh-release create sequenceiq/packer-azure $(VERSION) circle-ci

.PHONY:release
