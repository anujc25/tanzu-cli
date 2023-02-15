// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"crypto/sha256"
	"fmt"
	"github.com/vmware-tanzu/tanzu-cli/pkg/distribution"
	"github.com/vmware-tanzu/tanzu-cli/pkg/plugininventory"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/aunum/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/vmware-tanzu/tanzu-cli/pkg/carvelhelpers"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
	"github.com/vmware-tanzu/tanzu-cli/pkg/publisher"
	"github.com/vmware-tanzu/tanzu-cli/pkg/utils"
	configtypes "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
)

const PublisherPluginAssociationURL = "https://gist.githubusercontent.com/marckhouzam/5b653daf0afb815152f45aade5bc5d08/raw/7cd7e79c55361b492f611c8c090640daa5be1d9d"

type PublisherOptions struct {
	DryRun             bool
	ArtifactDir        string
	Publisher          string
	Vendor             string
	Repository         string
	PluginManifestFile string
	PublishScriptFile  string
	bashScript         string
}

type PublisherImpl interface {
	PublishPlugins() error
}

func (po *PublisherOptions) PublishPlugins() error {
	po.bashScript = "#!/usr/bin/env bash\n"

	log.Infof("Starting plugin publishing process...")

	if po.PluginManifestFile == "" {
		po.PluginManifestFile = filepath.Join(po.ArtifactDir, cli.PluginManifestFileName)
	}
	if po.PublishScriptFile == "" {
		po.PublishScriptFile = "plugin_publish_script.sh"
	}

	dbImage := fmt.Sprintf("%s/central:latest", po.Repository)
	tempDBDir, err := getTempDirectory("db")
	if err != nil {
		return errors.Wrap(err, "error creating temporary directory")
	}

	pluginManifest, err := po.getPluginManifest()
	if err != nil {
		return err
	}

	log.Infof("Using plugin location: %q, Publisher: %q, Vendor: %q, Repository: %q, PluginManifest: %q",
		po.ArtifactDir, po.Publisher, po.Vendor, po.Repository, po.PluginManifestFile)

	log.Info("Verifying plugin artifacts...")
	if err := po.verifyPluginBinaryArtifacts(pluginManifest); err != nil {
		return errors.Wrap(err, "error while verifying artifacts")
	}
	log.Info("Successfully verified plugin artifacts")

	//log.Info("Verifying plugin and publisher association...")
	//if err := po.verifyPluginAndPublisherAssociation(pluginManifest); err != nil {
	//	return errors.Wrap(err, "error while verifying artifacts")
	//}
	log.Info("Skipping plugin and publisher association verification")

	mapPluginArtifacts, err := po.organizePluginArtifactsForPublishing(pluginManifest)
	if err != nil {
		return errors.Wrapf(err, "unable to create temp artifacts directory for publishing")
	}

	log.Info("Insert and verify plugins on database...")
	err = po.insertPluginsToLocalDatabase(dbImage, tempDBDir, mapPluginArtifacts)
	if err != nil {
		return errors.Wrapf(err, "error while updating central database index")
	}

	log.Info("Publishing plugin binaries to the repository...")
	err = po.publishPluginsFromPluginArtifacts(mapPluginArtifacts)
	if err != nil {
		return errors.Wrapf(err, "error while publishing plugins to the repository")
	}

	log.Info("Publishing database...")
	err = po.publishDatabase(dbImage, tempDBDir)
	if err != nil {
		return errors.Wrapf(err, "error while updating central database index")
	}

	if po.DryRun {
		log.Infof("Saving the publish script to file: %q", po.PublishScriptFile)
		return po.savePublishScriptToFile()
	}
	return nil
}

// verifyPluginBinaryArtifacts verifies that the specified plugin binaries are available for all required OS_Arch
func (po *PublisherOptions) verifyPluginBinaryArtifacts(pluginManifest *cli.Manifest) error {
	var errList []error
	for i := range pluginManifest.Plugins {
		for _, osArch := range cli.MinOSArch {
			for _, version := range pluginManifest.Plugins[i].Versions {
				pluginFilePath := filepath.Join(po.ArtifactDir, osArch.OS(), osArch.Arch(),
					pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name, version,
					cli.MakeArtifactName(pluginManifest.Plugins[i].Name, osArch))

				if !utils.PathExists(pluginFilePath) {
					errList = append(errList, errors.Errorf("unable to verify artifacts for "+
						"plugin: %q, target: %q, osArch: %q, version: %q. File %q doesn't exist",
						pluginManifest.Plugins[i].Name, pluginManifest.Plugins[i].Target, osArch.String(), version, pluginFilePath))
				}
			}
		}
	}
	return kerrors.NewAggregate(errList)
}

func (po *PublisherOptions) verifyPluginAndPublisherAssociation(pluginManifest *cli.Manifest) error {
	f, err := os.CreateTemp("", "*.yaml")
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s-%s.yaml", PublisherPluginAssociationURL, po.Vendor, po.Publisher)
	log.Infof("Using url: %s", url)
	err = utils.DownloadFile(f.Name(), url)
	if err != nil {
		return errors.Wrapf(err, "error while downloading plugin publisher association file %q", url)
	}
	b, err := os.ReadFile(f.Name())
	if err != nil {
		return errors.Wrapf(err, "error while reading downloaded plugin publisher association file %q", f.Name())
	}

	registeredPluginsForPublisher := &publisher.PublisherPluginAssociation{}
	err = yaml.Unmarshal(b, registeredPluginsForPublisher)
	if err != nil {
		return errors.Wrapf(err, "error while unmarshaling downloaded plugin publisher association file %q", f.Name())
	}

	var errList []error
	for i := range pluginManifest.Plugins {
		found := false
		for j := range registeredPluginsForPublisher.Plugins {
			if pluginManifest.Plugins[i].Name == registeredPluginsForPublisher.Plugins[j].Name &&
				configtypes.StringToTarget(strings.ToLower(pluginManifest.Plugins[i].Target)) ==
					configtypes.StringToTarget(strings.ToLower(registeredPluginsForPublisher.Plugins[j].Target)) {
				found = true
			}
		}
		if !found {
			errList = append(errList, errors.Errorf("plugin: %q with target: %q is not registered for vendor: %q, publisher: %q",
				pluginManifest.Plugins[i].Name, pluginManifest.Plugins[i].Target, po.Vendor, po.Publisher))
		}
	}
	return kerrors.NewAggregate(errList)
}

// getPluginManifest reads the PluginManifest file and returns Manifest object
func (po *PublisherOptions) getPluginManifest() (*cli.Manifest, error) {
	data, err := os.ReadFile(po.PluginManifestFile)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read the plugin manifest file")
	}

	pluginManifest := &cli.Manifest{}
	err = yaml.Unmarshal(data, pluginManifest)
	if err != nil {
		return nil, errors.Wrap(err, "fail to read the plugin manifest file")
	}
	return pluginManifest, nil
}

// organizePluginArtifactsForPublishing organizes the plugin artifacts for publishing and returns the map of plugin name_target to PluginInventoryEntry
// As part of organizing the plugin artifacts function implements following steps:
// - Create temporary directory to hold plugin binary artifacts
// - Iterate through all the plugin specified in the plugin manifest for all supported OS_ARCH and create PluginInventoryEntry
func (po *PublisherOptions) organizePluginArtifactsForPublishing(pluginManifest *cli.Manifest) (map[string]plugininventory.PluginInventoryEntry, error) {
	tmpDir, err := getTempDirectory("plugin_artifacts")
	if err != nil {
		return nil, err
	}

	mapPluginArtifacts := make(map[string]plugininventory.PluginInventoryEntry)
	for i := range pluginManifest.Plugins {
		for _, osArch := range cli.AllOSArch {
			for _, version := range pluginManifest.Plugins[i].Versions {
				locationWithinBaseDir := filepath.Join(osArch.OS(), osArch.Arch(),
					pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name, version,
					cli.MakeArtifactName(pluginManifest.Plugins[i].Name, osArch))

				pluginFilePath := filepath.Join(po.ArtifactDir, locationWithinBaseDir)
				tmpPluginFilePath := filepath.Join(tmpDir, locationWithinBaseDir)

				if !utils.PathExists(pluginFilePath) {
					continue
				}

				err := utils.CopyFile(pluginFilePath, tmpPluginFilePath)
				if err != nil {
					return nil, err
				}

				key := fmt.Sprintf("%s-%s", pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name)
				pa, exists := mapPluginArtifacts[key]
				if !exists {
					pa = plugininventory.PluginInventoryEntry{
						Name:        pluginManifest.Plugins[i].Name,
						Target:      configtypes.Target(pluginManifest.Plugins[i].Target),
						Description: pluginManifest.Plugins[i].Description,
						Publisher:   po.Publisher,
						Vendor:      po.Vendor,
						Artifacts:   make(map[string]distribution.ArtifactList),
					}
					mapPluginArtifacts[key] = pa
				}
				_, exists = pa.Artifacts[version]
				if !exists {
					pa.Artifacts[version] = make([]distribution.Artifact, 0)
				}

				digest, err := getDigest(tmpPluginFilePath)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to calculate digest for path %v", tmpPluginFilePath)
				}
				am := distribution.Artifact{
					OS:     osArch.OS(),
					Arch:   osArch.Arch(),
					Digest: digest,
					URI:    tmpPluginFilePath,
					Image: fmt.Sprintf("%s/%s/%s/%s/%s/%s:%s", po.Vendor, po.Publisher, osArch.OS(), osArch.Arch(),
						pluginManifest.Plugins[i].Target, pluginManifest.Plugins[i].Name, version),
				}
				pa.Artifacts[version] = append(pa.Artifacts[version], am)
			}
		}
	}
	return mapPluginArtifacts, nil
}

// publishPluginsFromPluginArtifacts
func (po *PublisherOptions) publishPluginsFromPluginArtifacts(mapPluginArtifacts map[string]plugininventory.PluginInventoryEntry) error {
	var errList []error
	for _, pa := range mapPluginArtifacts {
		for _, artifacts := range pa.Artifacts {
			for _, a := range artifacts {
				pluginImage := fmt.Sprintf("%s/%s", po.Repository, a.Image)

				if po.DryRun {
					cmd := fmt.Sprintf("imgpkg push -i %s -f %s", pluginImage, filepath.Dir(a.URI))
					po.bashScript = po.bashScript + "\n" + cmd
				} else {
					err := carvelhelpers.UploadImage(pluginImage, filepath.Dir(a.URI))
					if err != nil {
						errList = append(errList, err)
					}
				}
			}
		}
	}
	return kerrors.NewAggregate(errList)
}

func (po *PublisherOptions) insertPluginsToLocalDatabase(centralDBImage, tempDir string, mapPluginArtifacts map[string]plugininventory.PluginInventoryEntry) error {
	err := carvelhelpers.DownloadImageAndSaveFilesToDir(centralDBImage, tempDir)
	if err != nil {
		return errors.Wrapf(err, "failed to download image '%s'", centralDBImage)
	}

	sqliteDB := plugininventory.NewSQLiteInventory("", tempDir, "")
	for _, pluginInventoryEntry := range mapPluginArtifacts {
		err := sqliteDB.InsertPlugin(&pluginInventoryEntry)
		if err != nil {
			log.Warningf("error while inserting plugin: %v", pluginInventoryEntry)
			//return errors.Wrapf(err, "error while inserting plugin: %v", pluginInventoryEntry)
		}
	}
	return nil
}

// publishDatabase
func (po *PublisherOptions) publishDatabase(centralDBImage, tempDir string) error {
	if po.DryRun {
		cmd := fmt.Sprintf("imgpkg push -i %s -f %s", centralDBImage, tempDir)
		po.bashScript = po.bashScript + "\n" + cmd
	} else {
		err := carvelhelpers.UploadImage(centralDBImage, tempDir)
		if err != nil {
			return errors.Wrapf(err, "failed to upload image '%s' to update central database image", centralDBImage)
		}
	}
	return nil
}

func (po *PublisherOptions) savePublishScriptToFile() error {
	err := utils.SaveFile(po.PublishScriptFile, []byte(po.bashScript))
	if err != nil {
		return errors.Wrapf(err, "error while saving publishing script to %s", po.PublishScriptFile)
	}
	return nil
}

func getTempDirectory(dirName string) (string, error) {
	tempDir := filepath.Join(os.TempDir(), dirName)
	err := os.RemoveAll(tempDir)
	if err != nil {
		return "", err
	}
	err = os.Mkdir(tempDir, 0755)
	if err != nil {
		return "", err
	}
	return tempDir, nil
}

// getDigest computes the sha256 digest of the specified file
func getDigest(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
