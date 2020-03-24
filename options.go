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
