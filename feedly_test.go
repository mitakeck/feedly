package feedly

import (
	"log"
	"net/url"
	"strings"
	"testing"
)

func TestDuplicateOption(t *testing.T) {
	f := Feedly{}
	f.Auth()

	// get markers
	markers, _ := f.MarkersCount()

	// get grobal all stream id
	var grobalStreamID string
	for _, marker := range markers.Unreadcounts {
		if strings.HasSuffix(marker.ID, "global.all") {
			grobalStreamID = marker.ID
			break
		}
	}
	log.Println("stream id : ", grobalStreamID)
	// get grobal all stream content
	_, err := f.StreamID(grobalStreamID, url.Values{
		"streamId": {"dup"}, // error : key duplicate
	})

	if err == nil {
		t.Errorf("duplicate option")
	}
}
