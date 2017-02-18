package feedly

import (
	"log"
	"testing"
)

func TestCategory(t *testing.T) {
	f := Feedly{}
	f.Auth()
	_, err := f.Categories()
	if err != nil {
		log.Print(err)
		t.Error("error when fetch category")
	}
}
