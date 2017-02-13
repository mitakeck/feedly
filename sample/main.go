package main

import (
	"fmt"
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
	fmt.Printf("AccessToken : %s\n", *token.AccessToken)

	// ------

	// カテゴリ取得
	profile, err := feedly.Categories()
	if err != nil {
		log.Print(err)
		return
	}
	pp.Print(profile)
}
