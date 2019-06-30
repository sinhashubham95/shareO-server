package external

import (
	"github.com/sinhashubham95/shareO-server/httpclient"
	"strconv"
	"strings"
)

// Page is the type for paginated page information
type Page struct {
	Count      int  `json:"count"`
	Start      int  `json:"start"`
	NextStart  int  `json:"nextStart"`
	IsLastPage bool `json:"isLastPage"`
}

// Image is the type for paginated image information
type Image struct {
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Type      string    `json:"type"`
	Dimension Dimension `json:"dimension"`
}

// PaginatedImageList is the return type
type PaginatedImageList struct {
	Data []Image `json:"data"`
	Page Page    `json:"page"`
}

// PageInfo is the type for page info
type PageInfo struct {
	Count      int `json:"count"`
	StartIndex int `json:"startIndex"`
}

// Queries is the type for queries
type Queries struct {
	Request  []PageInfo `json:"request"`
	NextPage []PageInfo `json:"nextPage"`
}

// Dimension is the type for dimension
type Dimension struct {
	Height string `json:"height"`
	Width  string `json:"width"`
}

// SearchItem is the type for search item
type SearchItem struct {
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Mime      string    `json:"mime"`
	Dimension Dimension `json:"image"`
}

// CustomImageSearch is the response for image custom google search
type CustomImageSearch struct {
	Queries Queries      `json:"queries"`
	Items   []SearchItem `json:"items"`
}

// GetImageList is used to fetch image list from Google Custom Search API
func GetImageList(searchType string, startIndex int) (*PaginatedImageList, error) {
	// constants
	customSearchURL := "https://www.googleapis.com/customsearch/v1"
	key := "key"
	cx := "cx"
	q := "q"
	safe := "safe"
	searchType = "searchType"
	start := "start"
	// initializing query params
	queryParams := make(map[string]string)
	queryParams[key] = "AIzaSyBaIIUZ9DOGOqQQI8t3339Z7MmJyCbP3ko"
	queryParams[cx] = "000674355350253632341:xp3u_vx45me"
	queryParams[q] = getImageSearchString(searchType)
	queryParams[safe] = "active"
	queryParams[searchType] = "image"
	queryParams[start] = getImageStartIndex(startIndex)
	// make call
	var data CustomImageSearch
	err := httpclient.GET(customSearchURL, nil, queryParams, true, &data)
	if err != nil {
		return nil, err
	}
	return parseImageSearchResponse(data), nil
}

func getImageSearchString(searchType string) string {
	// constants
	good := "good"
	morning := "morning"
	afternoon := "afternoon"
	evening := "evening"
	night := "night"
	// test
	switch strings.ToLower(searchType) {
	case morning:
		return good + morning
	case afternoon:
		return good + afternoon
	case evening:
		return good + evening
	case night:
		return good + night
	}
	return good + morning
}

func getImageStartIndex(startIndex int) string {
	return strconv.Itoa(startIndex)
}

func parseImageSearchResponse(data CustomImageSearch) *PaginatedImageList {
	return &PaginatedImageList{
		Data: getImageListFromSearch(data.Items),
		Page: getPageFromSearch(data.Queries),
	}
}

func getImageListFromSearch(items []SearchItem) []Image {
	list := make([]Image, 0, 0)
	for _, item := range items {
		list = append(list, Image{
			Title:     item.Title,
			Link:      item.Link,
			Type:      item.Mime,
			Dimension: item.Dimension,
		})
	}
	return list
}

func getPageFromSearch(queries Queries) Page {
	page := Page{
		Count: queries.Request[0].Count,
		Start: queries.Request[0].StartIndex,
	}
	if len(queries.NextPage) == 0 {
		page.IsLastPage = true
	} else {
		page.IsLastPage = false
		page.NextStart = queries.NextPage[0].StartIndex
	}
	return page
}
