package feedly

import (
	"fmt"
	"net/url"
)

// FeedResponse : GET /v3/feeds/:feedId
type FeedResponse struct {
	Language    string   `json:"language"`
	Curated     bool     `json:"curated"`
	Topics      []string `json:"topics"`
	Subscribers int      `json:"subscribers"`
	Featured    bool     `json:"featured"`
	Title       string   `json:"title"`
	Velocity    float64  `json:"velocity"`
	State       string   `json:"state"`
	Website     string   `json:"website"`
	ID          string   `json:"id"`
	Sponsored   bool     `json:"sponsored"`
}

// Feed : https://developer.feedly.com/v3/feeds/
func (f *Feedly) Feed(feedID string) (FeedResponse, error) {
	result := &FeedResponse{}
	if feedID == "" {
		return *result, fmt.Errorf("feedID is required")
	}
	u := feedURL + "/" + url.QueryEscape(feedID)
	_, err := f.request("GET", u, result, url.Values{})
	return *result, err
}
