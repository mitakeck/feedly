package feedly

// CategoriesResponse : GET /v3/categories
type CategoriesResponse []struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id"`
}

// Categories : https://developer.feedly.com/v3/categories/
func (f *Feedly) Categories() (CategoriesResponse, error) {
	result := &CategoriesResponse{}
	_, err := f.fetch("GET", categoriesURI, result)
	if err != nil {
		return *result, err
	}
	return *result, nil
}
