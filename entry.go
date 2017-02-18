package feedly

import (
	"fmt"
	"net/url"
)

// EntriesResponse : GET /v3/entries/:entryId
type EntriesResponse []struct {
	ID          string   `json:"id"`
	Keywords    []string `json:"keywords"`
	OriginID    string   `json:"originId"`
	Fingerprint string   `json:"fingerprint"`
	Origin      struct {
		StreamID string `json:"streamId"`
		Title    string `json:"title"`
		HTMLURL  string `json:"htmlUrl"`
	} `json:"origin"`
	Content struct {
		Content   string `json:"content"`
		Direction string `json:"direction"`
	} `json:"content"`
	Title     string `json:"title"`
	Published int64  `json:"published"`
	Crawled   int64  `json:"crawled"`
	Alternate []struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"alternate"`
	Author  string `json:"author"`
	Summary struct {
		Content   string `json:"content"`
		Direction string `json:"direction"`
	} `json:"summary"`
	Visual struct {
		Processor   string `json:"processor"`
		URL         string `json:"url"`
		Width       int    `json:"width"`
		Height      int    `json:"height"`
		ContentType string `json:"contentType"`
	} `json:"visual"`
	Unread bool `json:"unread"`
}

// Entry : https://developer.feedly.com/v3/entries/
func (f *Feedly) Entry(entryID string, options ...url.Values) (EntriesResponse, error) {
	result := &EntriesResponse{}
	option := url.Values{}
	for _, input := range options {
		f.setOption(&option, input)
	}
	_, e := f.request("GET", fmt.Sprintf(entryURL, entryID), result, option)
	return *result, e
}

// Entries : https://developer.feedly.com/v3/entries/
func (f *Feedly) Entries(entryIDs []string, options ...url.Values) ([]EntriesResponse, error) {
	result := &([]EntriesResponse{})
	option := url.Values{}
	for _, entryID := range entryIDs {
		option.Add("data", entryID)
	}
	for _, input := range options {
		if err := f.setOption(&option, input); err != nil {
			return *result, err
		}
	}
	_, e := f.request("POST", entriesURL, result, option)
	return *result, e
}
