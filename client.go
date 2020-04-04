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

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type Client interface {
	GetCategories(ctx context.Context) (GetCategoriesResp, error)
	GetCollections(ctx context.Context, page, limit string) (GetCollectionsResp, error)
	GetVideos(ctx context.Context, page string) (GetVideosResp, error)
	SearchVideos(ctx context.Context, query, page string) (SearchVideosResp, error)
	SearchJAVs(ctx context.Context, query, page string) (SearchJAVsResp, error)
	GetVideoByVID(ctx context.Context, vid string) (GetVideoByVIDResp, error)
}

var _ Client = &clientImpl{}

type clientImpl struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// NewClient returns a new Client.
func NewClient(opts ...ClientOption) (Client, error) {
	var o options
	for _, fn := range opts {
		fn(&o)
	}

	rawBaseURL := defaultOptions.baseURL
	if o.baseURL != "" {
		rawBaseURL = o.baseURL
	}

	httpClient := defaultOptions.httpClient
	if o.httpClient != nil {
		httpClient = o.httpClient
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse URL")
	}

	return &clientImpl{
		BaseURL:    baseURL,
		HTTPClient: httpClient,
	}, nil
}

// GetCategories retrieves all the video categories.
func (c *clientImpl) GetCategories(ctx context.Context) (GetCategoriesResp, error) {
	rawPath := "/categories"
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return GetCategoriesResp{}, err
	}

	var ret GetCategoriesResp
	if err := decodeBody(resp, &ret); err != nil {
		return GetCategoriesResp{}, err
	}

	return ret, nil
}

// GetCollections retrieves all the video collections.
func (c *clientImpl) GetCollections(ctx context.Context, page, limit string) (GetCollectionsResp, error) {
	const (
		defaultPage  = "0"
		defaultLimit = "50"
	)

	if page == "" {
		page = defaultPage
	}
	if limit == "" {
		limit = defaultLimit
	}

	rawPath := fmt.Sprintf("/collections/%s?limit=%s", page, limit)
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return GetCollectionsResp{}, err
	}

	var ret GetCollectionsResp
	if err := decodeBody(resp, &ret); err != nil {
		return GetCollectionsResp{}, err
	}

	return ret, nil
}

// GetVideos retrieves all the videos in criteria.
func (c *clientImpl) GetVideos(ctx context.Context, page string) (GetVideosResp, error) {
	const (
		defaultPage = "0"
	)

	if page == "" {
		page = defaultPage
	}

	rawPath := fmt.Sprintf("/videos/%s", page)
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return GetVideosResp{}, err
	}

	var ret GetVideosResp
	if err := decodeBody(resp, &ret); err != nil {
		return GetVideosResp{}, err
	}

	return ret, nil
}

// SearchVideos retrieves the videos matches the search query which are in criteria.
func (c *clientImpl) SearchVideos(ctx context.Context, query, page string) (SearchVideosResp, error) {
	const (
		defaultPage = "0"
	)

	if query == "" {
		return SearchVideosResp{}, errors.New("invalid argument: query is empty")
	}
	if page == "" {
		page = defaultPage
	}

	rawPath := fmt.Sprintf("/search/%s/%s", query, page)
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return SearchVideosResp{}, err
	}

	var ret SearchVideosResp
	if err := decodeBody(resp, &ret); err != nil {
		return SearchVideosResp{}, err
	}

	return ret, nil
}

// SearchJAVs retrieves the videos with category (channel, CHID) <= 12 that matches the search query which are in
// criteria, similar to video search.
func (c *clientImpl) SearchJAVs(ctx context.Context, query, page string) (SearchJAVsResp, error) {
	const (
		defaultPage = "0"
	)

	if query == "" {
		return SearchJAVsResp{}, errors.New("invalid argument: query is empty")
	}
	if page == "" {
		page = defaultPage
	}

	rawPath := fmt.Sprintf("/jav/%s/%s", query, page)
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return SearchJAVsResp{}, err
	}

	var ret SearchJAVsResp
	if err := decodeBody(resp, &ret); err != nil {
		return SearchJAVsResp{}, err
	}

	return ret, nil
}

// GetVideoByVID retrieves the video of specified VID.
func (c *clientImpl) GetVideoByVID(ctx context.Context, vid string) (GetVideoByVIDResp, error) {
	rawPath := fmt.Sprintf("/video/%s", vid)
	resp, err := c.get(ctx, rawPath)
	if err != nil {
		return GetVideoByVIDResp{}, err
	}

	var ret GetVideoByVIDResp
	if err := decodeBody(resp, &ret); err != nil {
		return GetVideoByVIDResp{}, err
	}

	if !ret.Success {
		return GetVideoByVIDResp{}, fmt.Errorf("video of VID %s not found", vid)
	}

	return ret, nil
}

func (c *clientImpl) newRequest(ctx context.Context, method, rawPath string, body io.Reader) (*http.Request, error) {
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, rawPath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req.WithContext(ctx), nil
}

func (c *clientImpl) get(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(req)
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
