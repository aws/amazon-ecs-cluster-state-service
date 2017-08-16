// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package regex

import (
	"regexp"
	"strconv"

	log "github.com/cihub/seelog"
)

// IsClusterName validates a cluster name against the cluster name regex
func IsClusterName(clusterName string) bool {
	validClusterName := regexp.MustCompile(ClusterNameRegex)
	if validClusterName.MatchString(clusterName) {
		return true
	}
	return false
}

// IsClusterARN validates a cluster ARN against the cluster ARN regex
func IsClusterARN(clusterARN string) bool {
	validClusterARN := regexp.MustCompile(ClusterARNRegex)
	if validClusterARN.MatchString(clusterARN) {
		return true
	}
	return false
}

// IsTaskARN validates a task ARN against the task ARN regex
func IsTaskARN(taskARN string) bool {
	validTaskARN := regexp.MustCompile(TaskARNRegex)
	if validTaskARN.MatchString(taskARN) {
		return true
	}
	return false
}

// IsInstanceARN validates an instance ARN against the instance ARN regex
func IsInstanceARN(instanceARN string) bool {
	validInstanceARN := regexp.MustCompile(InstanceARNRegex)
	if validInstanceARN.MatchString(instanceARN) {
		return true
	}
	return false
}

// IsEntityVersion validates an entity version as a positive integer
func IsEntityVersion(entityVersion string) bool {
	value, err := strconv.ParseInt(entityVersion, 10, 64)
	if err != nil {
		log.Warnf("Error parsing entity version: %v", err)
		return false
	}

	if value >= 0 {
		return true
	}
	return false
}