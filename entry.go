package feedly

import (
	"fmt"
	"net/url"
)

// EntriesResponse : GET /v3/entries/:entryId
type EntriesResponse struct {
	Author    string `json:"author"`
	Canonical []struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"canonical"`
	Crawled int64 `json:"crawled"`
	Visual  struct {
		URL         string `json:"url"`
		Height      int    `json:"height"`
		ContentType string `json:"contentType"`
		Width       int    `json:"width"`
	} `json:"visual"`
	Updated        int64   `json:"updated"`
	Title          string  `json:"title"`
	Fingerprint    string  `json:"fingerprint"`
	EngagementRate float64 `json:"engagementRate"`
	Tags           []struct {
		ID    string `json:"id"`
		Label string `json:"label,omitempty"`
	} `json:"tags"`
	Categories []struct {
		Label string `json:"label"`
		ID    string `json:"id"`
	} `json:"categories"`
	Published int64  `json:"published"`
	Recrawled int64  `json:"recrawled"`
	OriginID  string `json:"originId"`
	Enclosure []struct {
		Type string `json:"type,omitempty"`
		Href string `json:"href"`
	} `json:"enclosure"`
	Alternate []struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"alternate"`
	Origin struct {
		StreamID string `json:"streamId"`
		Title    string `json:"title"`
		HTMLURL  string `json:"htmlUrl"`
	} `json:"origin"`
	Engagement int  `json:"engagement"`
	Unread     bool `json:"unread"`
	Thumbnail  []struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	ID       string   `json:"id"`
	Keywords []string `json:"keywords"`
	Summary  struct {
		Content   string `json:"content"`
		Direction string `json:"direction"`
	} `json:"summary"`
}

// Entry : https://developer.feedly.com/v3/entries/
func (f *Feedly) Entry(entryID string, options ...url.Values) (EntriesResponse, error) {
	result := &EntriesResponse{}
	option := url.Values{}
	for _, input := range options {
		f.setOption(&option, input)
	}
	_, e := f.request("GET", fmt.Sprintf(entruesURL, entryID), result, option)
	return *result, e
}
