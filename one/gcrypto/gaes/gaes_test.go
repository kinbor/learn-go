package gaes

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var (
	content16       = []byte("abcdefghijklmnopqrstuvwxyz")
	content16Base64 = "P3VXDkiBMxOljQgmMNFSQL4sIhx9CFmy4Q+J0F2bfXc="

	iv       = []byte("1234567890123456")
	key16    = []byte("1234567890123456")
	key24    = []byte("123456789009876543211234")
	key32    = []byte("12345678900987654321123456789001")
	keyErr16 = []byte("1234567890123456abc")
	keyErr24 = []byte("123456789009876543211234abc")
	keyErr32 = []byte("12345678900987654321123456789001abc")
)

func TestEncrypt(t *testing.T) {
	data, err := Encrypt(content16, key16)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData := base64.StdEncoding.EncodeToString(data)
		if content16Base64 == tmpData {
			fmt.Println("package:gaes test Encrypt(content16,key16):OK!")
		} else {
			fmt.Println("package:gaes test Encrypt(content16,key16):Fail!" + tmpData)
		}
	}
}
