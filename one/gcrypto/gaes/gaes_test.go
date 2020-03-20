package gaes_test

import (
	"testing"
	"encoding/base64"
	"fmt"

	"github.com/kinbor/learn-gomod/one/gcrypto/gaes"
)

var (
	content_16			= []byte("abcdefghijklmnopqrstuvwxyz")
	content_16_base64	= "P3VXDkiBMxOljQgmMNFSQL4sIhx9CFmy4Q+J0F2bfXc="

	iv				= []byte("1234567890123456")
	key_16			= []byte("1234567890123456")
	key_24			= []byte("123456789009876543211234")
	key_32			= []byte("12345678900987654321123456789001")
	key_err_16		= []byte("1234567890123456abc")
	key_err_24		= []byte("123456789009876543211234abc")
	key_err_32		= []byte("12345678900987654321123456789001abc")
)

func TestEncrypt(t *testing.T) {
	data, err :=gaes.Encrypt(content_16, key_16)
	if err !=nil {
		fmt.Println(err)
	} else {
		tmpData :=base64.StdEncoding.EncodeToString(data)
		if content_16_base64 ==tmpData  {
			fmt.Println("package:gaes test Encrypt(content_16,key_16):OK!")
		} else {
			fmt.Println("package:gaes test Encrypt(content_16,key_16):Fail!"+tmpData)
		}
	}
}