// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package imgpkg

import (
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/vmware-tanzu/tanzu-cli/pkg/utils"
)

// CraneOptions implements the ImgpkgWrapper interface by using `imgpkg` binary internally
type CraneOptions struct{}

// SaveImage image as an archive file
func (co *CraneOptions) SaveImage(imageName, pluginTarGZFilePath string) error {
	img, err := getImageObj(imageName)
	if err != nil {
		return err
	}

	pluginTarFile, err := os.CreateTemp("", "*.tar")
	if err != nil {
		return err
	}
	defer os.Remove(pluginTarFile.Name())

	err = crane.Save(img, imageName, pluginTarFile.Name())
	if err != nil {
		return err
	}

	// convert the tar file into the tar.gz file
	return utils.Gzip(pluginTarFile.Name(), pluginTarGZFilePath)
}

// PushImage publish the archive file to remote container registry
func (co *CraneOptions) PushImage(archiveFile, image string) error {
	return nil
}

func getImageObj(imageName string) (v1.Image, error) {
	o := crane.GetOptions()
	ref, err := name.ParseReference(imageName, o.Name...)
	if err != nil {
		return nil, fmt.Errorf("parsing reference %q: %w", imageName, err)
	}
	rmt, err := remote.Get(ref, o.Remote...)
	if err != nil {
		return nil, err
	}

	img, err := rmt.Image()
	if err != nil {
		return nil, err
	}

	return img, nil
}
