package gsha1

import (
	"fmt"
	"testing"
)

func TestEncryptBytes(t *testing.T) {
	sha1Value, _ := EncryptBytes([]byte("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(sha1Value)
}

func TestEncryptString(t *testing.T) {
	sha1Value, _ := EncryptString("1234567890")
	fmt.Println(sha1Value)
}

func TestEncryptFile(t *testing.T) {
	sha1Value, _ := EncryptFile("sha1File.txt")
	fmt.Println(sha1Value)
}
