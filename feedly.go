package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://cloud.feedly.com"

	clientID    = "feedly"
	redirectURI = "localhost:8080"
	scope       = baseURL + "/subscriptions"
	provider    = "google"

	authenticateURI = baseURL + "/v3/auth/auth"
	accessTokenURI  = baseURL + "/v3/auth/token"
	profileURI      = baseURL + "/v3/profile"
	markerCountURI  = baseURL + "/v3/markers/counts"
)

func getCode() (string, error) {
	OpenBrowser(authenticateURI + "?client_id=" + clientID + "&redirect_uri=http://" + redirectURI + "&scope=" + scope + "&response_type=code&provider=" + provider + "&migrate=false")

	l, err := net.Listen("tcp", redirectURI)
	if err != nil {
		return "", fmt.Errorf("Unable to listen (%s): %v", redirectURI, err)
	}
	defer l.Close()

	quit := make(chan string)
	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// feedly からのリダイレクトを受け取る
		w.(http.Flusher).Flush()
		// URL 中の code を取得する
		quit <- req.URL.Query().Get("code")
	}))

	return <-quit, nil
}

func getAccessToken(code string) (AuthTokenResponse, error) {
	at := AuthTokenResponse{}

	resp, err := http.PostForm(accessTokenURI,
		url.Values{
			"client_id":     {clientID},
			"client_secret": {"0XP4XQ07VVMDWBKUHTJM4WUQ"},
			"grant_type":    {"authorization_code"},
			"redirect_uri":  {"http://www.feedly.com/feedly.html"},
			"code":          {code},
		},
	)
	if err != nil {
		return at, fmt.Errorf("Unable to postform : %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&at); err != nil {
		return at, fmt.Errorf("Unable to decode : %v", err)
	}
	return at, nil
}

// Powered by https://mholt.github.io/json-to-go/

// AuthTokenResponse : POST /v3/auth/token
type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Plan         string `json:"plan"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	ID           string `json:"id"`
	Provider     string `json:"provider"`
}

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

// CategoriesResponse : GET /v3/categories
type CategoriesResponse []struct {
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	ID          string `json:"id"`
}

// MarkersConuntsResponse : GET /v3/markers/counts
type MarkersConuntsResponse struct {
	Unreadcounts []struct {
		Count   int    `json:"count"`
		Updated int64  `json:"updated"`
		ID      string `json:"id"`
	} `json:"unreadcounts"`
	Updated int64 `json:"updated"`
}

func getProfile(accessToken string) (ProfileResponse, error) {
	result := ProfileResponse{}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", profileURI, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("Unable to fetch profile : %v", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("Unable to convert profile json : %v", err)
	}
	return result, nil
}

func getMarkersCount(accessToken string) (MarkersConuntsResponse, error) {
	result := MarkersConuntsResponse{}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", markerCountURI, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		return result, fmt.Errorf("Unable to fetch profile : %v", err)
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("Unable to convert markers count json : %v", err)
	}
	return result, nil
}
