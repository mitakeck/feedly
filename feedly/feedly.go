package feedly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL = "https://cloud.feedly.com"

	clientID    = "feedly"
	redirectURI = "localhost:8080"
	scope       = baseURL + "/subscriptions"
	provider    = "google"
	apiVersion  = "v3"

	authenticateURI  = "auth/auth"
	accessTokenURI   = "auth/token"
	profileURI       = "profile"
	markerCountURI   = "markers/counts"
	subscriptionsURI = "subscriptions"
	categoriesURI    = "categories"
	searchURL        = "search/feeds"
	mixesURL         = "mixes/contents"
	tagURL           = "tags"

	tokensFile = ".feedly-token"
	codeFile   = ".feedly-code"
)

// Feedly : feedly struct
type Feedly struct {
	authToken *AuthTokenResponse
}

// API 呼び出しの URI を組み立て
func (f *Feedly) createURI(suburl string) string {
	return baseURL + "/" + apiVersion + "/" + suburl
}

// JSON 取得処理
func (f *Feedly) request(method string, suburl string, v interface{}, param url.Values) (interface{}, error) {
	client := &http.Client{}
	u := f.createURI(suburl)
	log.Printf("http request %s (%s)\n", method, u)
	res := &http.Response{}

	switch method {
	case "GET":
		url, err := url.Parse(u)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse url (%s) : %v", u, err)
		}
		url.RawQuery = param.Encode()
		req, err := http.NewRequest(method, url.String(), strings.NewReader(param.Encode()))
		if err != nil {
			return nil, fmt.Errorf("Unable to create request (%s) : %v", url.String(), err)
		}
		req.Header.Set("Authorization", "Bearer "+*f.authToken.AccessToken)
		var resErr error
		res, resErr = client.Do(req)
		if resErr != nil {
			return nil, fmt.Errorf("Unable to fetch url (%s) : %v", u, resErr)
		}
	case "POST":
		var resErr error
		res, resErr = http.PostForm(u, param)
		if resErr != nil {
			return nil, fmt.Errorf("Unable to postform (%s) : %v", u, resErr)
		}
	}

	if res.StatusCode == 400 {
		byteArray, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("Bad request (%s) : %s", u, string(byteArray))
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("Unable to decode json : %v", err)
	}
	return v, nil
}
