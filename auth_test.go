package feedly

import "testing"

func TestAuth(t *testing.T) {
	f := Feedly{}
	token, err := f.Auth()
	if err != nil {
		t.Error("Error Feedly.Auth()")
	}
	if *token.AccessToken == "" {
		t.Error("Error Feedly.Auth() response is invalid")
	}
}
