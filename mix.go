package feedly

import (
	"fmt"
	"net/url"
)

// MixesResponse : GET /v3/mixes/:streamId/contents
type MixesResponse struct {
	Continuation string `json:"continuation"`
	Updated      int64  `json:"updated"`
	Title        string `json:"title"`
	Alternate    []struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"alternate"`
	Self []struct {
		Href string `json:" href"`
	} `json:"self"`
	Direction string `json:"direction"`
	ID        string `json:"id"`
	Items     []struct {
		Author  string `json:"author"`
		Crawled int64  `json:"crawled"`
		Updated int64  `json:"updated"`
		Title   string `json:"title"`
		Content struct {
			Content   string `json:"content"`
			Direction string `json:"direction"`
		} `json:"content"`
		Published int64 `json:"published"`
		Tags      []struct {
			Label string `json:"label"`
			ID    string `json:"id"`
		} `json:"tags"`
		Categories []struct {
			Label string `json:"label"`
			ID    string `json:"id"`
		} `json:"categories"`
		Alternate []struct {
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"alternate"`
		Origin struct {
			StreamID string `json:"streamId"`
			Title    string `json:"title"`
			HTMLURL  string `json:"htmlUrl"`
		} `json:"origin"`
		Engagement int    `json:"engagement"`
		Unread     bool   `json:"unread"`
		ID         string `json:"id"`
	} `json:"items"`
}

// Mix : https://developer.feedly.com/v3/mixes/
func (f *Feedly) Mix(streamID string, options ...url.Values) (MixesResponse, error) {
	result := &MixesResponse{}
	if streamID == "" {
		return *result, fmt.Errorf("streamID is required")
	}
	option := url.Values{
		"streamId": {streamID},
	}
	for _, input := range options {
		f.setOption(&option, input)
	}
	_, err := f.request("GET", mixesURL, result, option)
	return *result, err
}
