package feedly

import "net/url"

// SubscriptionsResponse : GET /v3/subscriptions
type SubscriptionsResponse []struct {
	Updated    int64  `json:"updated"`
	Title      string `json:"title"`
	Categories []struct {
		Label string `json:"label"`
		ID    string `json:"id"`
	} `json:"categories"`
	Website   string `json:"website"`
	ID        string `json:"id"`
	VisualURL string `json:"visualUrl,omitempty"`
	Sortid    string `json:"sortid"`
	Added     int64  `json:"added"`
}

// Subscriptions : https://developer.feedly.com/v3/subscriptions/
func (f Feedly) Subscriptions() (SubscriptionsResponse, error) {
	result := &SubscriptionsResponse{}
	_, err := f.request("GET", subscriptionsURI, result, url.Values{})
	return *result, err
}
