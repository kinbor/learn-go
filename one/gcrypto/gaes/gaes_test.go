package gaes

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

var (
	content16       = "abcdefghijklmnopqrstuvwxyz"
	content16Base64 = "PuQnJJSpi6RbG015hNI/kv0pMvymNOwBRNgUTodfcDo="

	iv       = []byte("1234567890123456")
	key16    = []byte("1234567890123456")
	key24    = []byte("123456789009876543211234")
	key32    = []byte("12345678900987654321123456789001")
	keyErr16 = []byte("1234567890123456abc")
	keyErr24 = []byte("123456789009876543211234abc")
	keyErr32 = []byte("12345678900987654321123456789001abc")
)

func TestGenKeyIV(t *testing.T) {
	rkey, riv := GenKeyIV(16)
	fmt.Println("Random Key Length：" + strconv.Itoa(len(rkey)))
	fmt.Println("Random Key String：" + hex.EncodeToString(rkey))
	fmt.Println("Random IV Length：" + strconv.Itoa(len(riv)))
	fmt.Println("Random IV String：" + hex.EncodeToString(riv))
	//key16 = rkey
	//iv = riv
}

func TestEncrypt(t *testing.T) {
	data, err := Encrypt([]byte(content16), key16)
	//data, err := Encrypt([]byte(content16), key16, iv)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData := base64.StdEncoding.EncodeToString(data)
		if content16Base64 == tmpData {
			fmt.Println("package:gaes test Encrypt(content16,key16):OK!" + tmpData)
		} else {
			fmt.Println("package:gaes test Encrypt(content16,key16):Fail!" + tmpData)
		}
		//content16Base64=tmpData
	}
}

func TestDecrypt(t *testing.T) {
	contentBytes, _ := base64.StdEncoding.DecodeString(content16Base64)
	data, err := Decrypt(contentBytes, key16)
	//data, err := Decrypt(contentBytes, key16, iv)
	if err != nil {
		fmt.Println(err)
	} else {
		tmpData := string(data)
		if content16 == tmpData {
			fmt.Println("package:gaes test Decrypt(content16,key16):OK!" + tmpData)
		} else {
			fmt.Println("package:gaes test Decrypt(content16,key16):Fail!" + tmpData)
		}
	}
}
