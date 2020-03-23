package gsha512

import (
	"fmt"
	"testing"
)

func TestEncryptBytes(t *testing.T) {
	sha512Value, _ := EncryptBytes([]byte("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(sha512Value)
}

func TestEncryptString(t *testing.T) {
	sha512Value, _ := EncryptString("1234567890")
	fmt.Println(sha512Value)
}

func TestEncryptFile(t *testing.T) {
	sha512Value, _ := EncryptFile("sha512File.txt")
	fmt.Println(sha512Value)
}
