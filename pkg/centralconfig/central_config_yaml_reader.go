// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package centralconfig implements an interface to deal with the central configuration.
package centralconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type centralConfigYamlReader struct {
	// configFile is the path to the central config file.
	configFile string
}

// Make sure centralConfigYamlReader implements CentralConfig
var _ CentralConfig = &centralConfigYamlReader{}

// parseConfigFile reads the central config file and returns the parsed yaml content.
// If the file does not exist, it does not return an error because some central repositories
// may choose not to have a central config file.
func (c *centralConfigYamlReader) parseConfigFile() (map[CentralConfigKey]CentralConfigValue, error) {
	// Check if the central config file exists.
	if _, err := os.Stat(c.configFile); os.IsNotExist(err) {
		// The central config file is optional, don't return an error if it does not exist.
		return nil, nil
	}

	bytes, err := os.ReadFile(c.configFile)
	if err != nil {
		return nil, err
	}

	var content map[CentralConfigKey]CentralConfigValue
	err = yaml.Unmarshal(bytes, &content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *centralConfigYamlReader) GetCentralConfigEntry(key CentralConfigKey) (CentralConfigValue, error) {
	values, err := c.parseConfigFile()
	if err != nil {
		return nil, err
	}

	value, isKeySet := values[key]
	if !isKeySet {
		return nil, nil
	}
	return value, nil
}
