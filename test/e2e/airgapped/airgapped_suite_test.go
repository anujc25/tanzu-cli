// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// airgapped provides airgapped specific E2E test cases
package airgapped

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/vmware-tanzu/tanzu-cli/test/e2e/framework"
	pluginlifecyclee2e "github.com/vmware-tanzu/tanzu-cli/test/e2e/plugin_lifecycle"
	"github.com/vmware-tanzu/tanzu-plugin-runtime/log"
)

func TestAirgapped(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Airgapped E2E Test Suite")
}

var (
	tf                           *framework.Framework
	e2eTestLocalCentralRepoImage string
	e2eAirgappedCentralRepo      string
	e2eAirgappedCentralRepoImage string
	pluginsSearchList            []*framework.PluginInfo
	pluginGroups                 []*framework.PluginGroup
	pluginGroupToPluginListMap   map[string][]*framework.PluginInfo
	pluginSourceName             string
	tempDir                      string
)

// BeforeSuite initializes and set up the environment to execute the airgapped tests
var _ = BeforeSuite(func() {
	tf = framework.NewFramework()
	// check E2E test central repo URL (TANZU_CLI_E2E_TEST_LOCAL_CENTRAL_REPO_URL)
	e2eTestLocalCentralRepoImage = os.Getenv(framework.TanzuCliE2ETestLocalCentralRepositoryURL)
	Expect(e2eTestLocalCentralRepoImage).NotTo(BeEmpty(), fmt.Sprintf("environment variable %s should set with local central repository URL", framework.TanzuCliE2ETestLocalCentralRepositoryURL))

	// check E2E airgapped central repo (TANZU_CLI_E2E_AIRGAPPED_REPO)
	e2eAirgappedCentralRepo = os.Getenv(framework.TanzuCliE2ETestAirgappedRepo)
	Expect(e2eAirgappedCentralRepo).NotTo(BeEmpty(), fmt.Sprintf("environment variable %s should set with airgapped central repository URL", framework.TanzuCliE2ETestAirgappedRepo))

	e2eAirgappedCentralRepoImage = fmt.Sprintf("%s%s", e2eAirgappedCentralRepo, filepath.Base(e2eTestLocalCentralRepoImage))

	// setup the test central repo
	_, err := tf.PluginCmd.UpdatePluginDiscoverySource(&framework.DiscoveryOptions{Name: "default", SourceType: framework.SourceType, URI: e2eTestLocalCentralRepoImage})
	Expect(err).To(BeNil(), "should not get any error for plugin source update")

	// search plugin groups and make sure there plugin groups available
	pluginGroups = pluginlifecyclee2e.SearchAllPluginGroups(tf)

	for _, pg := range pluginGroups {
		log.Infof("%v", pg)
	}

	// check all required plugin groups are available in the central repository with plugin group search output before running airgapped tests
	Expect(framework.IsAllPluginGroupsExists(pluginGroups, framework.PluginGroupsForLifeCycleTests)).Should(BeTrue(), "all required plugin groups for life cycle tests should exists in plugin group search output")

	// search plugins and make sure there are plugins available
	pluginsSearchList = pluginlifecyclee2e.SearchAllPlugins(tf)
	Expect(len(pluginsSearchList)).Should(BeNumerically(">", 0))

	// check all required plugins are available in the central repository with plugin search output before running airgapped tests
	Expect(framework.CheckAllPluginsExists(pluginsSearchList, framework.PluginsForLifeCycleTests)).To(BeTrue())

	// Configure temporary directory to save the plugin bundles for tests
	tempDir, err = os.MkdirTemp("", "")
	Expect(err).To(BeNil())
})

var _ = AfterSuite(func() {
	_ = os.RemoveAll(tempDir)
})
