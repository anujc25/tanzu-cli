// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package plugin implements plugin specific publishing functions
package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"

	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/helpers"
	"github.com/vmware-tanzu/tanzu-cli/cmd/plugin/builder/imgpkg"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
	"github.com/vmware-tanzu/tanzu-cli/pkg/utils"
)

type DownloadPluginPackageOptions struct {
	PackageArtifactDir string
	SourceRepository   string
	PluginManifestFile string

	ImgpkgOptions imgpkg.ImgpkgWrapper
}

func (dpo *DownloadPluginPackageOptions) DownloadPluginPackages() error {
	if dpo.PluginManifestFile == "" {
		return errors.New("plugin manifest file cannot be empty")
	}
	if !utils.PathExists(dpo.PackageArtifactDir) {
		err := os.MkdirAll(dpo.PackageArtifactDir, 0755)
		if err != nil {
			return err
		}
	}

	pluginManifest, err := helpers.ReadPluginManifest(dpo.PluginManifestFile)
	if err != nil {
		return err
	}

	for i := range pluginManifest.Plugins {
		for _, osArch := range cli.AllOSArch {
			for _, version := range pluginManifest.Plugins[i].Versions {
				pluginTarFilePath := filepath.Join(dpo.PackageArtifactDir, helpers.GetPluginArchiveRelativePath(pluginManifest.Plugins[i], osArch, version))
				image := fmt.Sprintf("%s/%s/%s/%s/%s:%s", dpo.SourceRepository, osArch.OS(), osArch.Arch(), pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name, version)

				if err := dpo.ImgpkgOptions.ResolveImage(image); err != nil {
					log.Infof("skipping %s_%s_%s_%s: unable to resolve image: %q", osArch.OS(), osArch.Arch(), pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name, image)
					continue
				}

				log.Infof("downloading plugin package for 'plugin:%s' 'target:%s' 'os:%s' 'arch:%s' 'version:%s'", pluginManifest.Plugins[i].Name, pluginManifest.Plugins[i].Target, osArch.OS(), osArch.Arch(), version)
				err = dpo.ImgpkgOptions.CopyImageToArchive(image, pluginTarFilePath)
				if err != nil {
					return errors.Wrapf(err, "unable to download package for plugin: %s, target: %s, os: %s, arch: %s, version: %s", pluginManifest.Plugins[i].Name, pluginManifest.Plugins[i].Target, osArch.OS(), osArch.Arch(), version)
				}

				log.Infof("downloaded plugin package at %q", pluginTarFilePath)
			}
		}
	}

	// copy plugin_manifest.yaml to PackageArtifactDir
	err = utils.CopyFile(dpo.PluginManifestFile, filepath.Join(dpo.PackageArtifactDir, cli.PluginManifestFileName))
	if err != nil {
		return errors.Wrap(err, "unable to copy plugin manifest file")
	}

	log.Infof("saved plugin manifest at %q", filepath.Join(dpo.PackageArtifactDir, cli.PluginManifestFileName))

	return nil
}
