// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package imgpkg implements helper function for imgpkg cli
package imgpkg

//go:generate counterfeiter -o ../fakes/imgpkgwrapper.go --fake-name ImgpkgWrapper . ImgpkgWrapper

// CraneWrapper defines the crane command wrapper functions
type CraneWrapper interface {
	// SaveImage image as an archive file
	SaveImage(image, pluginTarGZFilePath string) error
	// PushImage publish the archive file to remote container registry
	PushImage(pluginTarGZFilePath, image string) error
}

// NewCraneWrapper creates new CraneWrapper instance
func NewCraneWrapper() CraneWrapper {
	return &CraneOptions{}
}
