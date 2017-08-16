// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package httpclient provides a thin, but testable, wrapper around http.Client.
// It adds an Blox User agent header to requests and provides an interface

package httpclient

import (
	"fmt"
	"net/http"
	"time"

	"github.com/blox/blox/daemon-scheduler/versioning"
)

const (
	httpClientTimeout = 30 * time.Second
	userAgentHeader   = "User-Agent"
)

var userAgent string

// bloxRoundTripper helps set a custom user agent on HTTP requests.
type bloxRoundTripper struct {
	transport http.RoundTripper
}

func bloxCSSUserAgent() string {
	return fmt.Sprintf("Blox/%s Cluster State Service", versioning.Version)
}

func (rt *bloxRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(userAgentHeader, userAgent)
	return rt.transport.RoundTrip(req)
}

func init() {
	userAgent = bloxCSSUserAgent()
}

// New returns an Blox httpClient that will insert custom HTTP UA header.
func New() *http.Client {
	transport := http.DefaultTransport

	client := &http.Client{
		Transport: &bloxRoundTripper{transport},
		Timeout: httpClientTimeout,
	}

	return client
}
