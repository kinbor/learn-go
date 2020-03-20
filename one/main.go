package main

import (
	"fmt"
	"encoding/base64"
	
	"github.com/kinbor/learn-gomod/one/gcrypto"
)

func main()  {
	content:=[]byte("1234567890")
	key_16 := []byte("1234567891234567")

	data, err := gaes.Encrypt(content, key_16)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData :=base64.StdEncoding.EncodeToString(data)
		fmt.Println(tmpData)
	}
}