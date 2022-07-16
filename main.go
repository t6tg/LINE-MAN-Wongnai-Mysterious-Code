package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	fmt.Println(reverse(string(sd)))
}

func reverse(s string) (whatIsIt string) {
	for _, v := range string(s) {
		whatIsIt = string(v) + whatIsIt
	}
	return string(whatIsIt)
}
