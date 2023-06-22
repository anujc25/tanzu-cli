// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package crane

import (
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/cmd/crane/cmd"
	"github.com/google/go-containerregistry/pkg/crane"
)

// CraneOptions implements the ImgpkgWrapper interface by using `imgpkg` binary internally
type CraneOptions struct{}

// SaveImage image as an archive file
func (co *CraneOptions) SaveImage(imageName, archiveFile string) error {
	err := os.MkdirAll(filepath.Dir(archiveFile), 0755)
	if err != nil {
		return err
	}
	cranePullCmd := cmd.NewCmdPull(&[]crane.Option{})
	return cranePullCmd.RunE(cranePullCmd, []string{imageName, archiveFile})
}

// PushImage publish the archive file to remote container registry
func (co *CraneOptions) PushImage(archiveFile, image string) error {
	cranePushCmd := cmd.NewCmdPush(&[]crane.Option{})
	return cranePushCmd.RunE(cranePushCmd, []string{archiveFile, image})
}
