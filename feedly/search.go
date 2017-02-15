package feedly

import (
	"fmt"
	"net/url"
)

// SearchResponse : GET /v3/search/feeds
type SearchResponse struct {
	Results []struct {
		Curated     bool    `json:"curated"`
		Featured    bool    `json:"featured"`
		Subscribers int     `json:"subscribers"`
		Title       string  `json:"title"`
		Velocity    float64 `json:"velocity"`
		FeedID      string  `json:"feedId"`
		Website     string  `json:"website"`
	} `json:"results"`
	Related []string `json:"related"`
	Hint    string   `json:"hint"`
}

// Search : https://developer.feedly.com/v3/search/
func (f *Feedly) Search(query string) (SearchResponse, error) {
	result := &SearchResponse{}
	if query == "" {
		return *result, fmt.Errorf("Search query is required")
	}
	_, err := f.request("GET", searchURL, result, url.Values{
		"query": {query},
	})
	return *result, err
}
