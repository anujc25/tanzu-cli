// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package recommendedversion

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/tj/assert"

	"github.com/vmware-tanzu/tanzu-cli/pkg/constants"
	"github.com/vmware-tanzu/tanzu-cli/pkg/datastore"
	"github.com/vmware-tanzu/tanzu-cli/pkg/utils"
)

func TestFindRecommendedMajorVersion(t *testing.T) {
	tests := []struct {
		name        string
		recommended []string
		current     string
		expected    string
	}{
		{
			name:        "Newer major",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.0",
			expected:    "v2.0.2",
		},
		{
			name:        "Newer major skip one",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v0.90.0",
			expected:    "v2.0.2",
		},
		{
			name:        "Same major",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v2.0.0",
			expected:    "",
		},
		{
			name:        "Older major",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v3.3.3",
			expected:    "v2.0.2",
		},
		{
			name:        "Pre-release",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.0.2-rc.0",
			expected:    "v2.1.0-alpha.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			includePreReleases := utils.IsPreRelease(tt.current)

			if got := FindRecommendedMajorVersion(tt.recommended, tt.current, includePreReleases); got != tt.expected {
				t.Errorf("FindRecommendedMajorVersion() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFindRecommendedMinorVersion(t *testing.T) {
	tests := []struct {
		name        string
		recommended []string
		current     string
		expected    string
	}{
		{
			name:        "Newer minor",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.0",
			expected:    "v1.4.4",
		},
		{
			name:        "Newer minor skip one",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.0.0",
			expected:    "v1.4.4",
		},
		{
			name:        "Same minor",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.4.4",
			expected:    "",
		},
		{
			name:        "Older minor",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v2.1.0",
			expected:    "v2.0.2",
		},
		{
			name:        "Pre-release",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.0.2-rc.0",
			expected:    "v1.5.0-beta.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			includePreReleases := utils.IsPreRelease(tt.current)

			if got := FindRecommendedMinorVersion(tt.recommended, tt.current, includePreReleases); got != tt.expected {
				t.Errorf("FindRecommendedMinorVersion() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFindRecommendedPatchVersion(t *testing.T) {
	tests := []struct {
		name        string
		recommended []string
		current     string
		expected    string
	}{
		{
			name:        "Newer patch",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.2",
			expected:    "v1.3.3",
		},
		{
			name:        "Newer patch skip one",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.0",
			expected:    "v1.3.3",
		},
		{
			name:        "Same patch",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.3",
			expected:    "",
		},
		{
			name:        "Older patch",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.3.4",
			expected:    "v1.3.3",
		},
		{
			name:        "Pre-release",
			recommended: strings.Split("v2.1.0-alpha.2,v2.0.2,v1.5.0-beta.0,v1.4.4,,v1.3.3,v1.2.2,v1.1.1,v0.90.0", ","),
			current:     "v1.5.0-alpha.1",
			expected:    "v1.5.0-beta.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			includePreReleases := utils.IsPreRelease(tt.current)

			if got := FindRecommendedPatchVersion(tt.recommended, tt.current, includePreReleases); got != tt.expected {
				t.Errorf("FindRecommendedPatchVersion() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSortRecommendedVersionsDescending(t *testing.T) {
	tests := []struct {
		name        string
		recommended string
		expected    []string
		expectedErr string
	}{
		{
			name:        "Mixed versions",
			recommended: "v1.1.1,v2.0.2,v2.1.0-alpha.2,v1.3.3,v0.90.0,v1.4.4,v1.5.0-beta.0,v1.2.2",
			expected:    []string{"v2.1.0-alpha.2", "v2.0.2", "v1.5.0-beta.0", "v1.4.4", "v1.3.3", "v1.2.2", "v1.1.1", "v0.90.0"},
		},
		{
			name:        "Mixed versions with empty versions",
			recommended: "v1.1.1,v2.0.2,v2.1.0-alpha.2,,v1.3.3,v0.90.0,,v1.4.4,v1.5.0-beta.0,v1.2.2",
			expected:    []string{"v2.1.0-alpha.2", "v2.0.2", "v1.5.0-beta.0", "v1.4.4", "v1.3.3", "v1.2.2", "v1.1.1", "v0.90.0"},
		},
		{
			name:        "Mixed versions with spaces and empty versions",
			recommended: "v1.1.1 ,v2.0.2,   v2.1.0-alpha.2,  ,v1.3.3,  v0.90.0,,v1.4.4,  v1.5.0-beta.0, v1.2.2",
			expected:    []string{"v2.1.0-alpha.2", "v2.0.2", "v1.5.0-beta.0", "v1.4.4", "v1.3.3", "v1.2.2", "v1.1.1", "v0.90.0"},
		},
		{
			name:        "With invalid versions",
			recommended: "v1.1.1,v2.0.2,invalid-version",
			expectedErr: "Invalid Semantic Version",
		},
		{
			name:        "Repeating versions",
			recommended: "v1.1.1,v2.0.2,v2.1.0-alpha.2,v1.1.1,v0.90.0,v1.1.1,v1.1.1,v1.2.2",
			expected:    []string{"v2.1.0-alpha.2", "v2.0.2", "v1.2.2", "v1.1.1", "v0.90.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortRecommendedVersionsDescending(tt.recommended)
			if tt.expectedErr != "" {
				if err == nil {
					t.Errorf("SortRecommendedVersionsDescending() should have returned an error")
				} else {
					if err.Error() != tt.expectedErr {
						t.Errorf("SortRecommendedVersionsDescending() error = %v, want %v", err.Error(), tt.expectedErr)
					}
				}
			} else {
				if !arraysAreEqual(got, tt.expected) {
					t.Errorf("SortRecommendedVersionsDescending() = %v, want %v", got, tt.expected)
				}
			}
		})
	}
}

func TestGetRecommendationDelayValue(t *testing.T) {
	tests := []struct {
		name          string
		delayOverride string
		want          int
	}{
		{
			name:          "No override",
			delayOverride: "",
			want:          24 * 60 * 60, // 24 hours
		},
		{
			name:          "With smaller override",
			delayOverride: "36",
			want:          36,
		},
		{
			name:          "With larger override",
			delayOverride: strconv.Itoa(24*60*60 + 100),
			want:          24*60*60 + 100,
		},
		{
			name:          "With 0 override",
			delayOverride: "0",
			want:          0,
		},
		{
			name:          "With negative override",
			delayOverride: "-100",
			want:          24 * 60 * 60, // 24 hours
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.delayOverride != "" {
				_ = os.Setenv(constants.ConfigVariableRecommendVersionDelay, tt.delayOverride)
				defer func() {
					_ = os.Setenv(constants.ConfigVariableRecommendVersionDelay, "")
				}()
			}
			delay := getRecommendationDelayValue()
			if delay != tt.want {
				t.Errorf("getRecommendationDelayValue() = %v, want %v", delay, tt.want)
			}
		})
	}
}

func TestShouldCheckVersion(t *testing.T) {
	tests := []struct {
		name          string
		delayOverride string
		lastCheckTime interface{}
		want          bool
	}{
		{
			name:          "Last check an hour ago",
			lastCheckTime: time.Now().Add(-time.Hour),
			want:          false,
		},
		{
			name:          "Last check 25 hour ago",
			lastCheckTime: time.Now().Add(-25 * time.Hour),
			want:          true,
		},
		{
			name:          "Empty last check time",
			lastCheckTime: nil,
			want:          true,
		},
		{
			name:          "Invalid last check time",
			lastCheckTime: "not a timestamp",
			want:          true,
		},
		{
			name:          "Disable version check",
			delayOverride: "0",
			lastCheckTime: time.Now().Add(-25 * time.Hour),
			want:          false,
		},
		{
			name:          "Shorten delay, don't check",
			delayOverride: "5",
			lastCheckTime: time.Now().Add(-2 * time.Second),
			want:          false,
		},
		{
			name:          "Shorten delay, do check",
			delayOverride: "5",
			lastCheckTime: time.Now().Add(-6 * time.Second),
			want:          true,
		},
	}

	tmpDataStoreFile, _ := os.CreateTemp("", "data-store.yaml")
	defer os.RemoveAll(tmpDataStoreFile.Name())
	os.Setenv("TEST_CUSTOM_DATA_STORE_FILE", tmpDataStoreFile.Name())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.lastCheckTime != nil {
				_ = datastore.SetDataStoreValue(dataStoreLastVersionCheckKey, tt.lastCheckTime)
			} else {
				_, _ = datastore.DeleteDataStoreValue(dataStoreLastVersionCheckKey)
			}

			if tt.delayOverride != "" {
				_ = os.Setenv(constants.ConfigVariableRecommendVersionDelay, tt.delayOverride)
				defer func() {
					_ = os.Setenv(constants.ConfigVariableRecommendVersionDelay, "")
				}()
			}

			got := shouldCheckVersion()
			if got != tt.want {
				t.Errorf("shouldCheckVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintVersionRecommendations(t *testing.T) {
	tests := []struct {
		name             string
		currentVersion   string
		recommendedPatch string
		recommendedMinor string
		recommendedMajor string
		contains         []string
		timestampNotSet  bool
	}{
		{
			name:            "No recommendation",
			currentVersion:  "v1.3.0",
			contains:        []string{},
			timestampNotSet: true,
		},
		{
			name:             "Recommend patch",
			currentVersion:   "v1.3.0",
			recommendedPatch: "v1.3.3",
			contains:         []string{"Note:", "v1.3.3"},
		},
		{
			name:             "Recommend minor",
			currentVersion:   "v1.3.0",
			recommendedMinor: "v1.4.3",
			contains:         []string{"Note:", "v1.4.3"},
		},
		{
			name:             "Recommend major",
			currentVersion:   "v1.3.0",
			recommendedMajor: "v2.4.3",
			contains:         []string{"Note:", "v2.4.3"},
		},
		{
			name:             "Recommend patch and minor",
			currentVersion:   "v1.3.0",
			recommendedPatch: "v1.3.3",
			recommendedMinor: "v1.4.3",
			contains:         []string{"Note:", "v1.3.3", "v1.4.3"},
		},
		{
			name:             "Recommend patch and major",
			currentVersion:   "v1.3.0",
			recommendedPatch: "v1.3.3",
			recommendedMajor: "v2.4.3",
			contains:         []string{"Note:", "v1.3.3", "v2.4.3"},
		},
		{
			name:             "Recommend major and minor",
			currentVersion:   "v1.3.0",
			recommendedMinor: "v1.4.3",
			recommendedMajor: "v2.4.3",
			contains:         []string{"Note:", "v1.4.3", "v2.4.3"},
		},
		{
			name:             "Recommend patch and minor and major",
			currentVersion:   "v1.3.0",
			recommendedMajor: "v2.4.3",
			recommendedPatch: "v1.3.3",
			recommendedMinor: "v1.4.3",
			contains:         []string{"Note:", "v1.3.3", "v1.4.3", "v2.4.3"},
		},
		{
			name:             "Recommend older patch",
			currentVersion:   "v1.3.9",
			recommendedPatch: "v1.3.3",
			contains:         []string{"WARNING:", "v1.3.3"},
		},
		{
			name:             "Recommend older minor",
			currentVersion:   "v1.9.0",
			recommendedMinor: "v1.4.3",
			contains:         []string{"WARNING:", "v1.4.3"},
		},
		{
			name:             "Recommend older major",
			currentVersion:   "v9.3.0",
			recommendedMajor: "v2.4.3",
			contains:         []string{"WARNING:", "v2.4.3"},
		},
		{
			name:             "With newer pre-release",
			currentVersion:   "v1.3.0-alpha.1",
			recommendedMajor: "v2.4.3-rc.0",
			recommendedMinor: "v1.4.3-alpha.0",
			recommendedPatch: "v1.3.0-beta.0",
			contains:         []string{"Note:", "v2.4.3-rc.0", "v1.4.3-alpha.0", "v1.3.0-beta.0"},
		},
		{
			name:             "Downgrade pre-release",
			currentVersion:   "v1.3.0-alpha.1",
			recommendedMajor: "v2.4.3-rc.0",
			recommendedMinor: "v1.4.3-alpha.0",
			recommendedPatch: "v1.3.0-alpha.0",
			contains:         []string{"WARNING:", "v2.4.3-rc.0", "v1.4.3-alpha.0", "v1.3.0-alpha.0"},
		},
	}

	tmpDataStoreFile, _ := os.CreateTemp("", "data-store.yaml")
	defer os.RemoveAll(tmpDataStoreFile.Name())
	os.Setenv("TEST_CUSTOM_DATA_STORE_FILE", tmpDataStoreFile.Name())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			var buf bytes.Buffer
			printVersionRecommendations(&buf, tt.currentVersion, tt.recommendedMajor, tt.recommendedMinor, tt.recommendedPatch)
			if len(tt.contains) == 0 {
				assert.Empty(buf.String())
			} else {
				for i := range tt.contains {
					assert.Contains(buf.String(), tt.contains[i])
				}
				// Check that the variable to override is mentioned
				assert.Contains(buf.String(), constants.ConfigVariableRecommendVersionDelay)
			}

			// Check that the timestamp is updated
			timestamp := datastore.GetDataStoreValue(dataStoreLastVersionCheckKey)
			if tt.timestampNotSet {
				assert.Nil(timestamp)
			} else {
				assert.WithinDuration(time.Now(), timestamp.(time.Time), 1*time.Second)
			}
		})
	}
}

func arraysAreEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
