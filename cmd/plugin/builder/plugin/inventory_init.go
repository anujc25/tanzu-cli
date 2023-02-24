// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aunum/log"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/tanzu-cli/pkg/plugininventory"
)

type InventoryInitOptions struct {
	Repository        string
	InventoryImageTag string
	Override          bool
}

func (iio *InventoryInitOptions) InitializeInventory() error {
	// create plugin inventory database image path
	pluginInventoryDBImage := fmt.Sprintf("%s/%s:%s", iio.Repository, PluginInventoryDBImageName, iio.InventoryImageTag)

	if !iio.Override {
		// check if the image already exists or not
		err := resolveImage(pluginInventoryDBImage)
		if err == nil {
			return errors.Errorf("image %q already exists on the repository. Use `--override` flag to override the content", pluginInventoryDBImage)
		}
	}

	// Create plugin inventory database
	dbFile := filepath.Join(os.TempDir(), plugininventory.SQliteDBFileName)
	err := plugininventory.NewSQLiteInventory(dbFile, "").CreateSchema()
	if err != nil {
		return errors.Wrap(err, "error while creating database")
	}
	log.Infof("Create database locally at: %q", dbFile)

	// Publish the database to the remote repository
	log.Infof("Publishing database at: %q", pluginInventoryDBImage)
	err = pushImage(pluginInventoryDBImage, dbFile)
	if err != nil {
		return errors.Wrapf(err, "error while publishing database to the repository as image: %q", pluginInventoryDBImage)
	}
	log.Infof("Successfully published plugin inventory database")

	return nil
}
