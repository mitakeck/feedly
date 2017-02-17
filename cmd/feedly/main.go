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
	var unreadCount int
	for _, marker := range markers.Unreadcounts {
		if strings.HasSuffix(marker.ID, "global.all") {
			grobalStreamID = marker.ID
			unreadCount = marker.Count
			break
		}
	}

	// get grobal all stream content
	streams, _ := feedly.StreamContent(grobalStreamID, feedly.Count(unreadCount), feedly.Ranked("newest"))

	// show feed content
	for _, item := range streams.Items {
		fmt.Printf("[%s](%s)\n", item.Title, item.Alternate[0].Href)
	}

	// download OPML
	// feedly.OPML("opml.xml")
}
