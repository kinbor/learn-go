package gsha512

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

//EncryptBytes Encrypt []byte
func EncryptBytes(data []byte) (encrypt string, err error) {
	h := sha512.New()
	if _, err = h.Write([]byte(data)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

//EncryptString Encrypt string
func EncryptString(data string) (encrypt string, err error) {
	return EncryptBytes([]byte(data))
}

//EncryptFile Encrypt File By filePath
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha512.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
