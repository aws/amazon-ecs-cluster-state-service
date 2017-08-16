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

package v1

const (
	// 4xx error messages
	instanceNotFoundClientErrMsg             = "Instance not found"
	taskNotFoundClientErrMsg                 = "Task not found"
	invalidStatusClientErrMsg                = "Invalid status"
	unsupportedFilterClientErrMsg            = "At least one of the filters provided is unsupported"
	redundantFilterClientErrMsg              = "At least one of the filters provided is specified multiple times"
	invalidClusterClientErrMsg               = "Invalid cluster ARN or name"
	unsupportedFilterCombinationClientErrMsg = "The combination of filters provided are not supported"
	invalidEntityVersionClientErrMsg         = "Invalid entity version"
	outOfRangeEntityVersionClientErrMsg      = "Entity version is out of range"

	// 5xx error messages
	internalServerErrMsg = "Unexpected internal server error"
	encodingServerErrMsg = "Unexpected server error while encoding response"
	routingServerErrMsg  = "Unexpected server error related to api handler function routing"
)
