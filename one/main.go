package main

import (
	"encoding/base64"
	"fmt"

	"github.com/kinbor/learn-gomod/one/gcrypto/gaes"
)

func main() {
	content := []byte("1234567890")
	key16 := []byte("1234567891234567")

	data, err := gaes.Encrypt(content, key16)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData := base64.StdEncoding.EncodeToString(data)
		fmt.Println(tmpData)
	}
}
