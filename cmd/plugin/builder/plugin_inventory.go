// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/plugin"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
)

// newPluginInventoryCmd creates a new command for plugin inventory operations.
func newPluginInventoryCmd() *cobra.Command {
	var pluginInventoryCmd = &cobra.Command{
		Use:   "inventory",
		Short: "Plugin Inventory Operations",
	}

	pluginInventoryCmd.SetUsageFunc(cli.SubCmdUsageFunc)

	pluginInventoryCmd.AddCommand(
		newPluginInventoryInitCmd(),
	)

	return pluginInventoryCmd
}

type pluginInventoryInitFlags struct {
	Repository        string
	InventoryImageTag string
	Override          bool
}

func newPluginInventoryInitCmd() *cobra.Command {
	var piiFlags = &pluginInventoryInitFlags{}

	var pluginInventoryInitCmd = &cobra.Command{
		Use:          "init",
		Short:        "Initialize empty plugin inventory database and publish it to the remote repository",
		Example:      ``,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			iiOptions := plugin.InventoryInitOptions{
				Repository:        piiFlags.Repository,
				InventoryImageTag: piiFlags.InventoryImageTag,
				Override:          piiFlags.Override,
			}
			return iiOptions.InitializeInventory()
		},
	}

	pluginInventoryInitCmd.Flags().StringVarP(&piiFlags.Repository, "repository", "", "", "repository to publish plugin inventory image")
	pluginInventoryInitCmd.Flags().StringVarP(&piiFlags.InventoryImageTag, "plugin-inventory-image-tag", "", "latest", "tag to which plugin inventory image needs to be published")
	pluginInventoryInitCmd.Flags().BoolVarP(&piiFlags.Override, "override", "", false, "override the inventory database image if already exists")
	_ = pluginInventoryInitCmd.MarkFlagRequired("repository")

	return pluginInventoryInitCmd
}
