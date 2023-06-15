// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package imgpkg

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/helpers"
	"github.com/vmware-tanzu/tanzu-cli/pkg/utils"
)

// ImgpkgOptions implements the ImgpkgWrapper interface by using `imgpkg` binary internally
type ImgpkgOptions struct{}

// ResolveImage invokes `imgpkg tag resolve -i <image>` command
func (io *ImgpkgOptions) ResolveImage(image string) error {
	output, err := exec.Command(imgpkgBinary(), "tag", "resolve", "-i", image).CombinedOutput()
	return errors.Wrapf(err, "output: %s", string(output))
}

// PushImage invokes `imgpkg push -i <image> -f <filepath>` command
func (io *ImgpkgOptions) PushImage(image, filePath string) error {
	output, err := exec.Command(imgpkgBinary(), "push", "-i", image, "-f", filePath).CombinedOutput()
	return errors.Wrapf(err, "output: %s", string(output))
}

// PullImage invokes `imgpkg pull -i <image> -o <dirPath>` command
func (io *ImgpkgOptions) PullImage(image, dirPath string) error {
	output, err := exec.Command(imgpkgBinary(), "pull", "-i", image, "-o", dirPath).CombinedOutput()
	return errors.Wrapf(err, "output: %s", string(output))
}

// GetFileDigestFromImage invokes `PullImage` to fetch the image and returns the digest of the specified file
func (io *ImgpkgOptions) GetFileDigestFromImage(image, fileName string) (string, error) {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", errors.Wrap(err, "unable to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	// Pull image to the temporary directory
	err = io.PullImage(image, tempDir)
	if err != nil {
		return "", errors.Wrapf(err, "unable to find image at %q", image)
	}

	// find the digest of the specified file
	digest, err := helpers.GetDigest(filepath.Join(tempDir, fileName))
	if err != nil {
		return "", errors.Wrapf(err, "unable to calculate digest for path %v", filepath.Join(tempDir, fileName))
	}
	return digest, nil
}

// CopyArchiveToRepo invokes `imgpkg copy --tar <archivePath> --to-repo <imageRepo>` command
func (io *ImgpkgOptions) CopyArchiveToRepo(imageRepo, pluginTarGZFilePath string) error {
	pluginTarFile, err := os.CreateTemp("", "*.tar")
	if err != nil {
		return err
	}
	defer os.Remove(pluginTarFile.Name())
	err = utils.UnGzip(pluginTarGZFilePath, pluginTarFile.Name())
	if err != nil {
		return err
	}

	output, err := exec.Command(imgpkgBinary(), "copy", "--tar", pluginTarFile.Name(), "--to-repo", imageRepo).CombinedOutput()
	return errors.Wrapf(err, "output: %s", string(output))
}

// CopyImageToArchive invokes `imgpkg copy -i <image> --to-tar <archivePath>` command
func (io *ImgpkgOptions) CopyImageToArchive(image, pluginTarGZFilePath string) error {
	err := os.MkdirAll(filepath.Dir(pluginTarGZFilePath), 0755)
	if err != nil {
		return err
	}

	pluginTarFile, err := os.CreateTemp("", "*.tar")
	if err != nil {
		return err
	}
	defer os.Remove(pluginTarFile.Name())

	output, err := exec.Command(imgpkgBinary(), "copy", "-i", image, "--to-tar", pluginTarFile.Name()).CombinedOutput()
	if err != nil {
		return errors.Wrapf(err, "output: %s", string(output))
	}

	// convert the tar file into the tar.gz file
	return utils.Gzip(pluginTarFile.Name(), pluginTarGZFilePath)
}

func imgpkgBinary() string {
	imgpkgBinary := os.Getenv("IMGPKG_BIN")
	if imgpkgBinary != "" {
		return imgpkgBinary
	}
	return "imgpkg"
}
