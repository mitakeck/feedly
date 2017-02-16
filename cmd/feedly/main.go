package main

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/mitakeck/feedly"
)

func main() {
	feedly := feedly.Feedly{}

	// Feedly 認証
	token, err := feedly.Auth()
	if err != nil {
		log.Print(err)
		return
	}
	pp.Println(token)

	// ------

	// Search : golang
	search, err := feedly.Search("golang")
	if err != nil {
		log.Print(err)
		return
	}
	pp.Print(search)

	// ------

	// Download OPML
	feedly.OPML("opml.xml")

}
