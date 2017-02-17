package feedly

import "net/url"

// TagResponse : GET /v3/tags
type TagResponse []struct {
	ID          string `json:"id"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}

// Tag : https://developer.feedly.com/v3/tags/
func (f *Feedly) Tag(options ...url.Values) (TagResponse, error) {
	result := &TagResponse{}
	option := url.Values{}
	for _, input := range options {
		f.setOption(&option, input)
	}
	_, err := f.request("GET", tagURL, result, option)
	return *result, err
}
