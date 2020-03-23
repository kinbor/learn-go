package gmd5

import (
	"fmt"
	"testing"
)

func TestEncryptBytes(t *testing.T) {
	md5Value, _ := EncryptBytes([]byte("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(md5Value)
}

func TestEncryptString(t *testing.T) {
	md5Value, _ := EncryptString("1234567890")
	fmt.Println(md5Value)
}

func TestEncryptFile(t *testing.T) {
	md5Value, _ := EncryptFile("md5File.txt")
	fmt.Println(md5Value)
}
