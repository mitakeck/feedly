package feedly

import "net/url"

// MarkersConuntsResponse : GET /v3/markers/counts
type MarkersConuntsResponse struct {
	Unreadcounts []struct {
		Count   int    `json:"count"`
		Updated int64  `json:"updated"`
		ID      string `json:"id"`
	} `json:"unreadcounts"`
	Updated int64 `json:"updated"`
}

// MarkersCount : https://developer.feedly.com/v3/markers/
func (f *Feedly) MarkersCount() (MarkersConuntsResponse, error) {
	result := &MarkersConuntsResponse{}
	_, err := f.request("GET", markerCountURI, result, url.Values{})
	return *result, err
}
