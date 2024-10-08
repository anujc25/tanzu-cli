// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package centralconfig

import "strconv"

// GetDefaultTanzuEndpoint returns default endpoint for the tanzu platform from the default
// central configuration file
func (c *centralConfigYamlReader) GetDefaultTanzuEndpoint() (string, error) {
	endpoint := ""
	err := c.GetCentralConfigEntry(KeyDefaultTanzuEndpoint, &endpoint)
	return endpoint, err
}

// GetPluginDBCacheRefreshThresholdSeconds returns default value for central db cache refresh in seconds
// from the default central configuration file
func (c *centralConfigYamlReader) GetPluginDBCacheRefreshThresholdSeconds() (int, error) {
	secondsStr := ""
	err := c.GetCentralConfigEntry(KeyDefaultPluginDBCacheRefreshThresholdSeconds, &secondsStr)
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.ParseInt(secondsStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(seconds), nil
}

// GetInventoryRefreshTTLSeconds returns default value for central db refresh TTL in seconds
// from the default central configuration file
func (c *centralConfigYamlReader) GetInventoryRefreshTTLSeconds() (int, error) {
	secondsStr := ""
	err := c.GetCentralConfigEntry(KeyDefaultInventoryRefreshTTLSeconds, &secondsStr)
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.ParseInt(secondsStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(seconds), nil
}

// GetTanzuPlatformEndpointToServiceEndpointMap returns Map of tanzu platform endpoint to service endpoints
// from the default central configuration file
func (c *centralConfigYamlReader) GetTanzuPlatformEndpointToServiceEndpointMap() (TanzuPlatformEndpointToServiceEndpointMap, error) {
	endpointMap := TanzuPlatformEndpointToServiceEndpointMap{}
	err := c.GetCentralConfigEntry(KeyTanzuEndpointMap, &endpointMap)
	return endpointMap, err
}

// GetTanzuPlatformSaaSEndpointList returns list of tanzu platform saas endpoints which can be a regular
// expression. When comparing the result please make sure to use regex match instead of string comparison
func (c *centralConfigYamlReader) GetTanzuPlatformSaaSEndpointList() []string {
	saasEndpointList := []string{}
	err := c.GetCentralConfigEntry(KeyTanzuPlatformSaaSEndpointsAsRegularExpression, &saasEndpointList)
	if err != nil {
		return defaultSaaSEndpoints
	}
	return saasEndpointList
}

// GetTanzuConfigEndpointUpdateVersion returns current version for the local configuration file update
// If the version specified here does not match with the local version stored in the datastore that means
// the local configuration file endpoint updates are required
func (c *centralConfigYamlReader) GetTanzuConfigEndpointUpdateVersion() (string, error) {
	endpointUpdateVersion := ""
	err := c.GetCentralConfigEntry(KeyTanzuConfigEndpointUpdateVersion, &endpointUpdateVersion)
	return endpointUpdateVersion, err
}

// GetTanzuConfigEndpointUpdateMapping returns mapping of old endpoints to new endpoints that needs to be updated
// in the user's local configuration file
func (c *centralConfigYamlReader) GetTanzuConfigEndpointUpdateMapping() (map[string]string, error) {
	endpointUpdateMapping := map[string]string{}
	err := c.GetCentralConfigEntry(KeyTanzuConfigEndpointUpdateMapping, &endpointUpdateMapping)
	return endpointUpdateMapping, err
}
