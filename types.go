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

type Category struct {
	CHID        string `json:"CHID"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	TotalVideos int    `json:"total_videos"`
	CategoryURL string `json:"category_url"`
	CoverURL    string `json:"cover_url"`
}

type Collection struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Keyword       string `json:"keyword"`
	CoverURL      string `json:"cover_url"`
	TotalViews    int    `json:"total_views"`
	VideoCount    int    `json:"video_count"`
	CollectionURL string `json:"collection_url"`
}

type Video struct {
	Vid         string  `json:"vid"`
	UID         string  `json:"uid"`
	Title       string  `json:"title"`
	Keyword     string  `json:"keyword"`
	Channel     string  `json:"channel"`
	Duration    float64 `json:"duration"`
	Framerate   float64 `json:"framerate"`
	Hd          bool    `json:"hd"`
	AddTime     int     `json:"addtime"`
	ViewNumber  int     `json:"viewnumber"`
	Likes       int     `json:"likes"`
	Dislikes    int     `json:"dislikes"`
	VideoURL    string  `json:"video_url"`
	EmbeddedURL string  `json:"embedded_url"`
	PreviewURL  string  `json:"preview_url"`
}

type GetCategoriesResp struct {
	Success  bool `json:"success"`
	Response struct {
		Categories []Category `json:"categories"`
	} `json:"response"`
}

type GetCollectionsResp struct {
	Success  bool `json:"success"`
	Response struct {
		HasMore          bool         `json:"has_more"`
		TotalCollections int          `json:"total_collections"`
		CurrentOffset    int          `json:"current_offset"`
		Limit            int          `json:"limit"`
		Collections      []Collection `json:"collections"`
	} `json:"response"`
}

type GetVideosResp struct {
	Success  bool `json:"success"`
	Response struct {
		HasMore       bool    `json:"has_more"`
		TotalVideos   int     `json:"total_videos"`
		CurrentOffset int     `json:"current_offset"`
		Limit         int     `json:"limit"`
		Videos        []Video `json:"videos"`
	} `json:"response"`
}

type SearchVideosResp GetVideosResp

type SearchJAVsResp GetVideosResp

type GetVideoByVIDResp struct {
	Success  bool `json:"success"`
	Response struct {
		Video Video `json:"video"`
	} `json:"response"`
}
