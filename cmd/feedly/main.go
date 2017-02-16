package main

import (
	"fmt"
	"strings"

	"github.com/mitakeck/feedly"
)

func main() {

	// Auth
	feedly := feedly.Feedly{}
	feedly.Auth()

	// get profile
	profile, _ := feedly.Profile()
	fmt.Println("Auth : " + profile.Email)

	// get markers
	markers, _ := feedly.MarkersCount()

	// get grobal all stream id
	var grobalStreamID string
	for _, marker := range markers.Unreadcounts {
		if strings.HasSuffix(marker.ID, "global.all") {
			grobalStreamID = marker.ID
			break
		}
	}

	// get grobal all stream content
	streams, _ := feedly.StreamContent(grobalStreamID)

	// show feed content
	for _, item := range streams.Items {
		fmt.Printf("[%s](%s)\n", item.Title, item.Alternate[0].Href)
	}

	// download OPML
	feedly.OPML("opml.xml")
}
