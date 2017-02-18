package feedly

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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
	feedURL          = "feeds"
	opmlURL          = "opml"
	entryURL         = "entries/%s"
	entriesURL       = "entries/.mget"
	streamIDURL      = "streams/%s/ids"
	streamContentURL = "streams/%s/contents"

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

	accessToken := os.Getenv("FEEDLY_ACCESS_TOKEN")

	if accessToken != "" && f.authToken != nil {
		accessToken = *f.authToken.AccessToken
	}
	if accessToken == "" && method != "AUTH" {
		return v, fmt.Errorf("requred Auth()")
	}

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
		req.Header.Set("Authorization", "Bearer "+accessToken)
		var resErr error
		res, resErr = client.Do(req)
		if resErr != nil {
			return nil, fmt.Errorf("Unable to fetch url (%s) : %v", u, resErr)
		}
	case "POST":
		slc, _ := json.Marshal(param)
		req, err := http.NewRequest(method, u, strings.NewReader(string(slc)))
		if err != nil {
			return nil, fmt.Errorf("Unable to create request (%s) : %v", u, err)
		}
		req.Header.Set("Authorization", "Bearer "+accessToken)
		req.Header.Set("Content-Type", "application/json")
		var resErr error
		res, resErr = client.Do(req)
		if resErr != nil {
			return nil, fmt.Errorf("Unable to fetch url (%s) : %v", u, resErr)
		}
	case "AUTH":
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

// Download : download file
func (f *Feedly) Download(suburl string, fileName string) (int64, error) {
	output, err := os.Create(fileName)
	if err != nil {
		return 0, fmt.Errorf("Error while creating %s : %v", fileName, err)
	}

	defer output.Close()
	u := f.createURI(suburl)
	url, err := url.Parse(u)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse url (%s) : %v", u, err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("Unable to create request (%s) : %v", url.String(), err)
	}

	req.Header.Set("Authorization", "Bearer "+*f.authToken.AccessToken)
	client := &http.Client{}
	res, resErr := client.Do(req)
	if resErr != nil {
		return 0, fmt.Errorf("Error while downloading %s : %v", url, err)
	}
	defer res.Body.Close()

	n, err := io.Copy(output, res.Body)
	if err != nil {
		return 0, fmt.Errorf("Error while downloading %s : %v", url, err)
	}

	return n, nil
}

func (f *Feedly) setOption(option *url.Values, input url.Values) error {
	for key, values := range input {
		for _, value := range values {
			if option.Get(key) == "" {
				option.Add(key, value)
			} else {
				return fmt.Errorf("Duplicate url.Values key (%s)", key)
			}
		}
	}
	return nil
}
