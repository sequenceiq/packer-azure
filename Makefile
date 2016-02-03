VERSION=v2

build-local:
	go build -a -v -o ~/.packer.d/plugins/packer-builder-azure   ./packer/plugin/packer-builder-azure/

release:
	rm -rf release; mkdir release
	GOOS=linux CGO_ENABLED=0 go get -installsuffix cgo ./...
	tar -czvf release/packer.tgz  -C  $(GOPATH)/bin/linux_amd64 packer-builder-azure packer-provisioner-azure-custom-script-extension packer-provisioner-powershell-azure
	gh-release create sequenceiq/packer-azure $(VERSION)

.PHONY:release
