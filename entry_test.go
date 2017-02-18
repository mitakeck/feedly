package feedly

import (
	"log"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestEntry(t *testing.T) {
	f := Feedly{}
	f.Auth()

	// get markers
	markers, _ := f.MarkersCount()

	// get grobal all stream id
	var grobalStreamID string
	var unreadCount int
	for _, marker := range markers.Unreadcounts {
		if strings.HasSuffix(marker.ID, "global.all") {
			grobalStreamID = marker.ID
			unreadCount = marker.Count
			break
		}
	}

	// get grobal all stream content
	streamIDs, _ := f.StreamID(grobalStreamID, url.Values{
		"count": {strconv.Itoa(unreadCount)},
	})

	if len(streamIDs.Ids) > 0 {
		_, err := f.Entry(streamIDs.Ids[0])
		log.Println(err)
		if err != nil {
			t.Error("test Feedly.Entry()")
		}
	}

}
