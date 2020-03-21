# avgle-go
The unofficial Go client for Avgle API

API Reference: https://avgle.github.io/doc/

## Installation
```
go get github.com/sapuri/avgle-go
```

## Usage
### Methods
```go
type Client interface {
	GetCategories(ctx context.Context) (GetCategoriesResp, error)
	GetCollections(ctx context.Context, page, limit string) (GetCollectionsResp, error)
	GetVideos(ctx context.Context, page string) (GetVideosResp, error)
	SearchVideos(ctx context.Context, query, page string) (SearchVideosResp, error)
	SearchJAVs(ctx context.Context, query, page string) (SearchJAVsResp, error)
	GetVideoByVID(ctx context.Context, vid string) (GetVideoByVIDResp, error)
}
```

### Example
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapuri/avgle-go"
)

func main() {
	ctx := context.Background()

	client, err := avgle.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.SearchVideos(ctx, "三上悠亜", "0")
	if err != nil {
		log.Fatal(err)
	}

	videos := result.Response.Videos
	for _, video := range videos {
		fmt.Println(video)
	}
}

```
