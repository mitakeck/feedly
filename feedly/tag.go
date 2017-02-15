package feedly

import "net/url"

// TagResponse : GET /v3/tags
type TagResponse []struct {
	ID          string `json:"id"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

// Tag : https://developer.feedly.com/v3/tags/
func (f *Feedly) Tag() (TagResponse, error) {
	result := &TagResponse{}
	_, err := f.request("GET", tagURL, result, url.Values{})
	return *result, err
}
