package gaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"

	//"github.com/kinbor/learn-gomod/one/gcrypto/gtest"
)

var (
	IVDefaultValue = "www.xlsbook.com!"
)

func Encrypt(plainText []byte, key []byte, iv ...[]byte)([]byte, error){
	//gtest.PrintText()
	return EncryptCBC(plainText, key, iv...)
}

func Decrypt(cipherText []byte, key []byte, iv ...[]byte)([]byte, error){
	return DecryptCBC(cipherText, key, iv...)
}


func EncryptCBC(plainText []byte,key []byte, iv ...[]byte)([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize :=block.BlockSize()
	plainText =PKCS5Padding(plainText, blockSize)
	ivValue :=([]byte)(nil)
	if len(iv)>0{
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}

	blockMode := cipher.NewCBCEncrypter(block, ivValue)
	cipherText :=make([]byte,len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return cipherText, nil
}

func DecryptCBC(cipherText []byte, key []byte, iv ...[]byte)([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}

	blockSize:=block.BlockSize()
	if len(cipherText) < blockSize {
		return nil,errors.New("cipherText too short")
	}

	ivValue :=([]byte)(nil)
	if len(iv) >0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}

	if len(cipherText)%blockSize !=0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plainText,cipherText)
	plainText,err2 := PKC5UnPadding(plainText, blockSize)	
	if err2 != nil {
		return nil, err2
	}
	return plainText, nil
}


func PKCS5Padding(src []byte, blockSize int)[]byte{
	padding :=blockSize - len(src)%blockSize
	padText :=bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKC5UnPadding(src []byte, blockSize int)([]byte, error){
	length := len(src)
	if blockSize <=0 {
		return nil,errors.New("invalid blockSize")
	}

	if length%blockSize !=0 || length==0 {
		return nil,errors.New("invalid data Length")
	}

	unpadding := int(src[length-1])
	if unpadding > blockSize || unpadding == 0 {
		return nil,errors.New("invalid padding")
	}

	padding := src[length-unpadding:]
	for i:=0; i<unpadding; i++ {
		if padding[i] != byte(unpadding) {
			return nil,errors.New("invalid padding")
		}
	}
	
	return src[:(length-unpadding)],nil
}

func EncryptCFB(plainText []byte, key []byte, padding *int, iv ...[]byte)([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainText, *padding = ZeroPadding(plainText, blockSize)
	ivValue :=([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue =[]byte(IVDefaultValue)
	}

	stream := cipher.NewCFBEncrypter(block, ivValue)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

func DecryptCFB(cipherText []byte, key []byte, unpadding int, iv ...[]byte)([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(IVDefaultValue)
	}
	stream := cipher.NewCFBDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	plainText = ZeroUnPadding(plainText, unpadding)
	return plainText, nil
}

func ZeroPadding(cipherText []byte, blockSize int)([]byte, int){
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(cipherText, padText...), padding
}

func ZeroUnPadding(plainText []byte, unpadding int)[]byte{
	length := len(plainText)
	return plainText[:(length - unpadding)]
}