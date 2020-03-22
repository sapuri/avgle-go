package avgle

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	cases := map[string]struct {
		baseURL    string
		httpClient *http.Client
	}{
		"default": {},
		"with BaseURL": {
			baseURL: "https://api.example.com",
		},
		"with HTTPClient": {
			httpClient: http.DefaultClient,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient(
				WithBaseURL(tc.baseURL),
				WithHTTPClient(tc.httpClient),
			)
			assert.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

func TestClientImpl_GetCategories(t *testing.T) {
	cases := map[string]struct {
	}{
		"success": {},
	}

	for name := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.GetCategories(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_GetCollections(t *testing.T) {
	cases := map[string]struct {
		page  string
		limit string
	}{
		"success": {
			page:  "0",
			limit: "50",
		},
		"page and limit are empty": {},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.GetCollections(context.Background(), tc.page, tc.limit)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_GetVideos(t *testing.T) {
	cases := map[string]struct {
		page string
	}{
		"success": {
			page: "0",
		},
		"page is empty": {},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.GetVideos(context.Background(), tc.page)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_SearchVideos(t *testing.T) {
	cases := map[string]struct {
		query   string
		page    string
		success bool
	}{
		"success": {
			query:   "三上悠亜",
			page:    "0",
			success: true,
		},
		"query is empty": {
			page:    "0",
			success: false,
		},
		"page is empty": {
			query:   "三上悠亜",
			success: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.SearchVideos(context.Background(), tc.query, tc.page)
			if !tc.success {
				assert.Error(t, err)
				assert.Equal(t, resp, SearchVideosResp{})
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_SearchJAVs(t *testing.T) {
	cases := map[string]struct {
		query   string
		page    string
		success bool
	}{
		"success": {
			query:   "SSNI-388",
			page:    "0",
			success: true,
		},
		"query is empty": {
			page:    "0",
			success: false,
		},
		"page is empty": {
			query:   "SSNI-388",
			success: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.SearchJAVs(context.Background(), tc.query, tc.page)
			if !tc.success {
				assert.Error(t, err)
				assert.Equal(t, resp, SearchJAVsResp{})
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_GetVideoByVID(t *testing.T) {
	cases := map[string]struct {
		vid     string
		success bool
	}{
		"success": {
			vid:     "374462",
			success: true,
		},
		"not found": {
			vid:     "0",
			success: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := NewClient()
			require.NoError(t, err)

			resp, err := client.GetVideoByVID(context.Background(), tc.vid)
			if !tc.success {
				assert.Error(t, err)
				assert.Equal(t, resp, GetVideoByVIDResp{})
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestClientImpl_GetCategoriesWithTestServer(t *testing.T) {
	const path = "/categories"

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/get_categories.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.GetCategories(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClientImpl_GetCollectionsWithTestServer(t *testing.T) {
	const (
		page  = "0"
		limit = "50"
	)
	path := fmt.Sprintf("/collections/%s?limit=%s", page, limit)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/get_collections.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.GetCollections(context.Background(), page, limit)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClientImpl_GetVideosWithTestServer(t *testing.T) {
	const page = "0"
	path := fmt.Sprintf("/videos/%s", page)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/get_videos.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.GetVideos(context.Background(), page)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClientImpl_SearchVideosWithTestServer(t *testing.T) {
	const (
		query = "三上悠亜"
		page  = "0"
	)
	path := fmt.Sprintf("/search/%s/%s", query, page)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/search_videos.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.SearchVideos(context.Background(), query, page)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClientImpl_SearchJAVsWithTestServer(t *testing.T) {
	const (
		query = "SSNI-388"
		page  = "0"
	)
	path := fmt.Sprintf("/jav/%s/%s", query, page)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/search_javs.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.SearchJAVs(context.Background(), query, page)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestClientImpl_GetVideoByVIDWithTestServer(t *testing.T) {
	const vid = "374462"
	path := fmt.Sprintf("/video/%s", vid)

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/get_video_by_vid.json")
	})

	client, err := NewClient(WithBaseURL(server.URL))
	require.NoError(t, err)

	resp, err := client.GetVideoByVID(context.Background(), vid)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
