package feedly

import "net/url"

// StreamIDsResponse : GET /v3/streams/ids?streamId=:streamId
type StreamIDsResponse struct {
	Ids          []string `json:"ids"`
	Continuation string   `json:"continuation"`
}

// StreamContentsResponse : GET /v3/streams/:streamId/contents
type StreamContentsResponse struct {
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

// StreamID : https://developer.feedly.com/v3/streams/
func (f *Feedly) StreamID(streamID string) (StreamIDsResponse, error) {
	result := &StreamIDsResponse{}
	esid := url.QueryEscape(streamID)
	_, e := f.request("GET", "streams/"+esid+"/ids", result, nil)
	return *result, e
}

// StreamContent : https://developer.feedly.com/v3/streams/
func (f *Feedly) StreamContent(streamID string) (StreamContentsResponse, error) {
	result := &StreamContentsResponse{}
	esid := url.QueryEscape(streamID)
	_, e := f.request("GET", "streams/"+esid+"/contents", result, nil)
	return *result, e
}
