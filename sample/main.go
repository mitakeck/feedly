package main

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/mitakeck/feedly-go/feedly"
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

	// // カテゴリ取得
	// profile, err := feedly.MarkersCount()
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// pp.Print(profile)

	// ----

	// 全件の stream 取得
	stream, err := feedly.StreamContent("user/103c35b9-edc0-449f-94da-800ff8b483a9/category/IT全般")
	if err != nil {
		log.Print(err)
		return
	}
	pp.Print(stream)

}
