package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	whatIsIt = reverseConcat(sd)
	fmt.Println(string(whatIsIt))
}

func reverseConcat(sd []byte) string {
	var reversed string

	for i := len(sd) - 1; i >= 0; i-- {
		reversed += string(sd[i])
	}

	return reversed
}
