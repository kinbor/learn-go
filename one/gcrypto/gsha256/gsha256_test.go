package gsha256

import (
	"fmt"
	"testing"
)

func TestEncryptBytes(t *testing.T) {
	sha256Value, _ := EncryptBytes([]byte("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(sha256Value)
}

func TestEncryptString(t *testing.T) {
	sha256Value, _ := EncryptString("1234567890")
	fmt.Println(sha256Value)
}

func TestEncryptFile(t *testing.T) {
	sha256Value, _ := EncryptFile("sha256File.txt")
	fmt.Println(sha256Value)
}
