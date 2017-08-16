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

package wrappers

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	clusterNameEnvVarName = "ECS_CLUSTER"
	keyPairNameEnvVarName = "EC2_KEY_PAIR"
	regionEnvVarName      = "AWS_REGION"
	ecsEndpointEnvVarName = "ECS_ENDPOINT"
	defaultECSClusterName = "E2ETestCluster"

	validateInstanceSleepTime = 1
)

var (
	latestAMIIDByRegion = map[string]string{
		"us-east-1":      "ami-b2df2ca4",
		"us-east-2":      "ami-832b0ee6",
		"us-west-1":      "ami-dd104dbd",
		"us-west-2":      "ami-022b9262",
		"eu-west-1":      "ami-a7f2acc1",
		"eu-west-2":      "ami-3fb6bc5b",
		"eu-central-1":   "ami-ec2be583",
		"ap-northeast-1": "ami-c393d6a4",
		"ap-southeast-1": "ami-a88530cb",
		"ap-southeast-2": "ami-8af8ffe9",
		"ca-central-1":   "ami-ead5688e",
	}
)

func GetClusterName() string {
	cluster := os.Getenv(clusterNameEnvVarName)
	if cluster == "" {
		return defaultECSClusterName
	}
	return cluster
}

func GetKeyPairName() string {
	return os.Getenv(keyPairNameEnvVarName)
}

func getECSEndpoint() (string, error) {
	endpoint := os.Getenv(ecsEndpointEnvVarName)
	if endpoint == "" {
		return "", fmt.Errorf("Empty endpoint. Please specify the ECS endpoint name using the '%s' environment variable", ecsEndpointEnvVarName)
	}

	return endpoint, nil
}

func GetLatestECSOptimizedAMIID() (string, error) {
	region := os.Getenv(regionEnvVarName)
	if region == "" {
		return "", fmt.Errorf("Empty region. Please specify the AWS region.")
	}
	amiID, ok := latestAMIIDByRegion[region]
	if !ok {
		return "", fmt.Errorf("Invalid region name.")
	}
	return amiID, nil

}

// ValidateECSContainerInstance validates that an ECS instance was launched
// within a cluster 'clusterName'
func ValidateECSContainerInstance(ecsWrapper ECSWrapper, clusterName string) error {
	// Takes some time for EC2 instance to come up and for the ECS agent to
	// register the container instance with ECS.
	// Retry ECS list call once every minute for 10 minutes
	found := false
	for i := 0; i < 10; i++ {
		instances, err := ecsWrapper.ListContainerInstances(clusterName)
		if err != nil {
			return err
		}
		if len(instances) >= 1 {
			found = true
			break
		}
		time.Sleep(validateInstanceSleepTime * time.Minute)
	}

	if !found {
		return errors.Errorf("No ECS instance found in cluster with name '%s'. ",
			clusterName)
	}

	return nil
}
