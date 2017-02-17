package feedly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
)

// AuthTokenResponse : POST /v3/auth/token
type AuthTokenResponse struct {
	AccessToken  *string `json:"access_token"`
	Plan         *string `json:"plan"`
	RefreshToken *string `json:"refresh_token"`
	ExpiresIn    *int64  `json:"expires_in"`
	TokenType    *string `json:"token_type"`
	ID           *string `json:"id"`
	Provider     *string `json:"provider"`
}

// Auth : 認証処理群
func (f *Feedly) Auth() (AuthTokenResponse, error) {
	result := AuthTokenResponse{}
	code, err := f.getCode()
	if err != nil {
		return result, fmt.Errorf("Fail to getCode : %v", err)
	}
	result, err = f.getAccessToken(code)
	if err != nil {
		return result, fmt.Errorf("Fail to getAccessToken : %v", err)
	}
	return result, nil
}

func (f *Feedly) getCode() (string, error) {
	// キャッシュから読み込む
	file, e := os.Open(codeFile)
	defer file.Close()

	if os.IsNotExist(e) {
		f.OpenBrowser(f.createURI(authenticateURI) + "?client_id=" + clientID + "&redirect_uri=http://" + redirectURI + "&scope=" + scope + "&response_type=code&provider=" + provider + "&migrate=false")
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
		code := <-quit
		ioutil.WriteFile(codeFile, []byte(code), 0777)
		return code, nil
	} else if e == nil {
		text, err := ioutil.ReadAll(file)
		if err != nil {
			return "", fmt.Errorf("Unable to read code : %v", err)
		}
		return string(text), err
	}
	return "", nil
}

func (f *Feedly) getAccessToken(code string) (AuthTokenResponse, error) {
	at := &AuthTokenResponse{}

	// キャッシュから読み込む
	file, e := os.Open(tokensFile)
	defer file.Close()

	if os.IsNotExist(e) {
		_, err := f.request(
			"AUTH",
			accessTokenURI,
			at,
			url.Values{
				"client_id":     {clientID},
				"client_secret": {"0XP4XQ07VVMDWBKUHTJM4WUQ"},
				"grant_type":    {"authorization_code"},
				"redirect_uri":  {"http://www.feedly.com/feedly.html"},
				"code":          {code},
			},
		)
		if err != nil {
			return *at, fmt.Errorf("Unable to postform : %v", err)
		}
		// キャッシュ処理
		text, err := json.MarshalIndent(&at, "", " ")
		if err != nil {
			return *at, fmt.Errorf("Unable to convert authTokenResponse : %v", err)
		}
		ioutil.WriteFile(tokensFile, text, 0777)
	} else {
		if err := json.NewDecoder(file).Decode(&at); err != nil {
			return *at, fmt.Errorf("Unable to decode : %v", err)
		}
	}
	f.authToken = at
	return *at, nil
}
