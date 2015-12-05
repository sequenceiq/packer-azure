// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See the LICENSE file in the project root for license information.

package azure

import "fmt"

// This is the common builder ID to all of these artifacts.
const BuilderId = "Azure.ServiceManagement.VMImage"

// Artifact is the result of running the azure builder.
type artifact struct {
	imageLabel    string
	imageName     string
	mediaLocation string
}

func (*artifact) BuilderId() string {
	return BuilderId
}

func (a *artifact) Files() []string {
	return nil
}

func (a *artifact) Id() string {
	return a.imageName
}

func (a *artifact) State(name string) interface{} {
	switch name {
	case "atlas.artifact.metadata":
		metadata := make(map[string]string)
		metadata["imageLabel"] = a.imageLabel
		metadata["imageName"] = a.imageName
		metadata["mediaLocation"] = a.mediaLocation
		return metadata
	default:
		return nil
	}
}

func (a *artifact) String() string {
	return fmt.Sprintf("{%s,%s,%s}",
		fmt.Sprintf("imageLabel: '%s'", a.imageLabel),
		fmt.Sprintf("imageName: '%s'", a.imageName),
		fmt.Sprintf("mediaLocation: '%s'", a.mediaLocation),
	)
}

func (a *artifact) Destroy() error {

	// TODO: remove image and vhd
	return nil
}
