// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package inventory

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/helpers"
	"github.com/vmware-tanzu/tanzu-cli/pkg/carvelhelpers"
	"github.com/vmware-tanzu/tanzu-cli/pkg/plugininventory"
)

func inventoryDBDownload(imageOperationsImpl carvelhelpers.ImageOperationsImpl, pluginInventoryDBImage, tempDir string) (string, error) {
	err := imageOperationsImpl.DownloadImageAndSaveFilesToDir(pluginInventoryDBImage, tempDir)
	if err != nil {
		return "", errors.Wrapf(err, "error while pulling database from the image: %q", pluginInventoryDBImage)
	}
	return filepath.Join(tempDir, plugininventory.SQliteDBFileName), nil
}

func inventoryDBUpload(imageOperationsImpl carvelhelpers.ImageOperationsImpl, pluginInventoryDBImage, dbFile string) error {
	err := imageOperationsImpl.PushImage(pluginInventoryDBImage, []string{dbFile})
	if err != nil {
		return errors.Wrapf(err, "error while publishing inventory database to the repository as image: %q", pluginInventoryDBImage)
	}
	return nil
}

// getFileDigestFromImage invokes `PullImage` to fetch the image and returns the digest of the specified file
func getFileDigestFromImage(imageOperationsImpl carvelhelpers.ImageOperationsImpl, image, fileName string) (string, error) {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", errors.Wrap(err, "unable to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	// Pull image to the temporary directory
	err = imageOperationsImpl.DownloadImageAndSaveFilesToDir(image, tempDir)
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
