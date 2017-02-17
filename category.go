package feedly

import "net/url"

// CategoriesResponse : GET /v3/categories
type CategoriesResponse []struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id"`
}

// Categories : https://developer.feedly.com/v3/categories/
func (f *Feedly) Categories(options ...url.Values) (CategoriesResponse, error) {
	result := &CategoriesResponse{}
	option := url.Values{}
	for _, input := range options {
		f.setOption(&option, input)
	}
	_, err := f.request("GET", categoriesURI, result, option)
	return *result, err
}
