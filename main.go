package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := b64.StdEncoding.DecodeString(secret)
	whatIsIt = reverse(sd)

	fmt.Println(whatIsIt)
}

func reverse(b []byte) string {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}
