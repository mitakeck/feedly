package feedly

import (
	"net/url"
	"testing"
)

func TestDuplicateOption(t *testing.T) {
	f := Feedly{}
	opt := url.Values{
		"key": {"value"},
	}
	err := f.setOption(&url.Values{
		"key": {"dup value"},
	}, opt)

	if err == nil {
		t.Errorf("duplicate option")
	}
}
