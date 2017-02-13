package feedly

import "net/url"

// ProfileResponse : GET /v3/profile
type ProfileResponse struct {
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	FamilyName string `json:"familyName"`
	Twitter    string `json:"twitter"`
	Picture    string `json:"picture"`
	Google     string `json:"google"`
	Wave       string `json:"wave"`
	GivenName  string `json:"givenName"`
	Facebook   string `json:"facebook"`
	Reader     string `json:"reader"`
	ID         string `json:"id"`
	Locale     string `json:"locale"`
}

// Profile : https://developer.feedly.com/v3/profile/
func (f *Feedly) Profile() (ProfileResponse, error) {
	result := &ProfileResponse{}
	_, err := f.request("GET", profileURI, result, url.Values{})
	if err != nil {
		return *result, err
	}
	return *result, nil
}
