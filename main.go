package main

import "fmt"

func main() {
	// 認証開始
	code, err := getCode()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Code: %s\n", code)
	token, err := getAccessToken(code)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("AccessToken : %s\n", token.AccessToken)

	// // marker counts を取得してみる
	// m, err := getMarkersCount(token.AccessToken)
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Print(m)
}
