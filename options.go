// Copyright 2020 sapuri
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package avgle

import "net/http"

// ClientOption represents options func.
type ClientOption func(*options)

type options struct {
	baseURL    string
	httpClient *http.Client
}

var defaultOptions = &options{
	baseURL:    defaultBaseURL,
	httpClient: http.DefaultClient,
}

// WithBaseURL sets baseURL as ClientOption.
func WithBaseURL(baseURL string) ClientOption {
	return func(o *options) {
		o.baseURL = baseURL
	}
}

// WithHTTPClient sets httpClient as ClientOption.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(o *options) {
		o.httpClient = httpClient
	}
}
