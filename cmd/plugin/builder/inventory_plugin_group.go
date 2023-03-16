// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/imgpkg"
	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/inventory"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
)

// newInventoryPluginCmd creates a new command for plugin inventory operations.
func newInventoryPluginGroupCmd() *cobra.Command {
	var inventoryPluginCmd = &cobra.Command{
		Use:   "plugin-group",
		Short: "Plugin-Group Inventory Operations",
	}

	inventoryPluginCmd.SetUsageFunc(cli.SubCmdUsageFunc)

	inventoryPluginCmd.AddCommand(
		newInventoryPluginGroupAddCmd(),
	)

	return inventoryPluginCmd
}

type inventoryPluginGroupAddFlags struct {
	GroupName             string
	Repository            string
	InventoryImageTag     string
	ManifestFile          string
	Publisher             string
	Vendor                string
	DeactivatePluginGroup bool
	Override              bool
}

func newInventoryPluginGroupAddCmd() *cobra.Command {
	var ipgaFlags = &inventoryPluginGroupAddFlags{}

	var pluginGroupAddCmd = &cobra.Command{
		Use:          "add",
		Short:        "Add the plugin-group to the inventory database available on the remote repository",
		SilenceUsage: true,
		Example:      ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			pgaOptions := inventory.InventoryPluginGroupUpdateOptions{
				GroupName:               ipgaFlags.GroupName,
				Repository:              ipgaFlags.Repository,
				InventoryImageTag:       ipgaFlags.InventoryImageTag,
				PluginGroupManifestFile: ipgaFlags.ManifestFile,
				Vendor:                  ipgaFlags.Vendor,
				Publisher:               ipgaFlags.Publisher,
				DeactivatePluginGroup:   ipgaFlags.DeactivatePluginGroup,
				Override:                ipgaFlags.Override,
				ImgpkgOptions:           imgpkg.NewImgpkgCLIWrapper(),
			}
			return pgaOptions.PluginGroupAdd()
		},
	}

	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.GroupName, "name", "", "", "name of the plugin group")
	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.Repository, "repository", "", "", "repository to publish plugin inventory image")
	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.InventoryImageTag, "plugin-inventory-image-tag", "", "latest", "tag to which plugin inventory image needs to be published")
	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.ManifestFile, "manifest", "", "", "manifest file specifying plugin details that needs to be processed")
	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.Vendor, "vendor", "", "", "name of the vendor")
	pluginGroupAddCmd.Flags().StringVarP(&ipgaFlags.Publisher, "publisher", "", "", "name of the publisher")
	pluginGroupAddCmd.Flags().BoolVarP(&ipgaFlags.DeactivatePluginGroup, "deactivate", "", false, "mark plugin-group as deactivated")
	pluginGroupAddCmd.Flags().BoolVarP(&ipgaFlags.Override, "override", "", false, "override the plugin-group if already exists")

	_ = pluginGroupAddCmd.MarkFlagRequired("name")
	_ = pluginGroupAddCmd.MarkFlagRequired("repository")
	_ = pluginGroupAddCmd.MarkFlagRequired("vendor")
	_ = pluginGroupAddCmd.MarkFlagRequired("publisher")
	_ = pluginGroupAddCmd.MarkFlagRequired("manifest")

	return pluginGroupAddCmd
}
