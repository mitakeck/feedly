package feedly

import (
	"log"
	"testing"
)

func TestCategory(t *testing.T) {
	f := Feedly{}
	_, errauth := f.Auth()
	if errauth != nil {
		log.Print(errauth)
		t.Error("erro Feedly.Auth()")
	}
	_, err := f.Categories()
	if err != nil {
		log.Print(err)
		t.Error("error Feedly.Categories()")
	}
}
