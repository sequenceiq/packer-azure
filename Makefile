VERSION=v1

build-local:
	go build -a -v -o ~/.packer.d/plugins/packer-builder-azure   ./packer/plugin/packer-builder-azure/

release:
	rm -rf release; mkdir release
	CGO_ENABLED=0 go get -installsuffix cgo ./...
	tar -czvf release/packer.tgz  -C  $(GOPATH)/bin packer-builder-azure packer-provisioner-azure-custom-script-extension packer-provisioner-powershell-azure
	gh-release create sequenceiq/packer-azure $(VERSION)

.PHONY:release
