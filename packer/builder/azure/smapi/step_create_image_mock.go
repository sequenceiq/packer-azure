package azure

import (
	"github.com/Azure/packer-azure/packer/builder/azure/common/constants"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepCreateImageMock struct {
	TmpServiceName    string
	TmpVmName         string
	UserImageLabel    string
	UserImageName     string
	RecommendedVMSize string
}

func (s *StepCreateImageMock) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get(constants.Ui).(packer.Ui)

	ui.Say("[MOCK] Creating Azure Image. If Succeed This Will Remove the Temorary VM...")

	// CatpureVMImage removes the VM
	state.Put(constants.ImageCreated, 1)
	state.Put(constants.VmExists, 0)

	return multistep.ActionContinue
}

func (s *StepCreateImageMock) Cleanup(state multistep.StateBag) {
}
