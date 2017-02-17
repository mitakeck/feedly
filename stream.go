package feedly

import (
	"fmt"
	"net/url"
	"strconv"
)

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
	_, e := f.request("GET", fmt.Sprintf(streamIDURL, esid), result, nil)
	return *result, e
}

// StreamContent : https://developer.feedly.com/v3/streams/
func (f *Feedly) StreamContent(streamID string, options ...func(*Feedly, *url.Values) error) (StreamContentsResponse, error) {
	result := &StreamContentsResponse{}
	o := url.Values{}
	for _, opt := range options {
		if err := opt(f, &o); err != nil {
			return *result, err
		}
	}
	esid := url.QueryEscape(streamID)
	_, e := f.request("GET", fmt.Sprintf(streamContentURL, esid), result, o)
	return *result, e
}

// Count : set api option count
func (f *Feedly) Count(count int) func(*Feedly, *url.Values) error {
	return func(f *Feedly, option *url.Values) error {
		return f.setCount(count, option)
	}
}

func (f *Feedly) setCount(count int, param *url.Values) error {
	if count < 0 || 10000 < count {
		return fmt.Errorf("Invalid param `count`. it is must be [1 10000] : %d", count)
	}
	param.Add("count", strconv.Itoa(count))
	return nil
}

// Ranked : set api option ranked
func (f *Feedly) Ranked(ranked string) func(*Feedly, *url.Values) error {
	return func(f *Feedly, option *url.Values) error {
		return f.setRanked(ranked, option)
	}
}

func (f *Feedly) setRanked(ranked string, param *url.Values) error {
	switch ranked {
	case "newest":
		fallthrough
	case "oldest":
		param.Add("ranked", ranked)
		return nil
	default:
		return fmt.Errorf("Invalid param `ranked`. it is must be  'newest' or 'oldest' : %s", ranked)
	}
}
