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

const (
	validClusterName = "clust_er-1"
	validClusterARN  = "arn:aws:ecs:us-east-1:123456789123:cluster/" + validClusterName

	invalidClusterName                 = "cluster1/cluster1"
	invalidClusterARNWithNoName        = "arn:aws:ecs:us-east-1:123456789123:cluster/"
	invalidClusterARNWithInvalidName   = "arn:aws:ecs:us-east-1:123456789123:cluster/" + invalidClusterName
	invalidClusterARNWithInvalidPrefix = "arn/cluster"

	validTaskARN                    = "arn:aws:ecs:us-east-1:123456789012:task/271022c0-f894-4aa2-b063-25bae55088d5"
	invalidTaskARNWithNoID          = "arn:aws:ecs:us-east-1:123456789123:task/"
	invalidTaskARNWithInvalidID     = "arn:aws:ecs:us-east-1:123456789123:task/271022c0-f894-4aa2-b063-25bae55088d5/-"
	invalidTaskARNWithInvalidPrefix = "arn/task"

	validInstanceARN                    = "arn:aws:ecs:us-east-1:123456789123:container-instance/4b6d45ea-a4b4-4269-9d04-3af6ddfdc597"
	invalidInstanceARNWithNoID          = "arn:aws:ecs:us-east-1:123456789123:container-instance/"
	invalidInstanceARNWithInvalidID     = "arn:aws:ecs:us-east-1:123456789123:container-instance/4b6d45ea-a4b4-4269-9d04-3af6ddfdc597/-"
	invalidInstanceARNWithInvalidPrefix = "arn/container-instance"

	validEntityVersion                      = "123"
	invalidEntityVersionFloatingPointNumber = "123.123"
	invalidEntityVersionNegativeNumber      = "-123"
	invalidEntityVersionNonNumber           = "invalidEntityVersion"

)
