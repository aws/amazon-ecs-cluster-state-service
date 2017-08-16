#!/bin/bash
# Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

# This script is meant for use inside the build dockerfile
# (scripts/dockerfiles/Dockerfile.build) and makes assumptions related to that.
# It should not be run on its own.

echo "Starting clean build"

cd /src/blox
DIRTY_WARNING=$(cat <<EOW
***WARNING***
You currently have uncommitted or unstaged changes in your git repository.
The release build will not include those and the result may behave differently
than expected due to that. Please commit, stash, or remove all uncommitted or
unstaged files before creating a release build.
EOW
)
[ ! -z "$(git status --porcelain)" ] && echo "$DIRTY_WARNING"

# Fresh clone to ensure our build doesn't rely on anything outside of vcs
git clone --quiet /src/blox /go/src/github.com/blox/blox

cd /go/src/github.com/blox/blox/cluster-state-service/canary
exec ./scripts/build_binary.sh /out
