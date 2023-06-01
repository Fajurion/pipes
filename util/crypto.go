package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// encrypt encrypts text with key using AES-256 (credit to https://stackoverflow.com/questions/18817336/encrypting-a-string-with-aes-and-base64 with slight modifications)
func EncryptAES(block cipher.Block, text []byte) ([]byte, error) {

	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

// decrypt decryptes text with key using AES-256 (credit to https://stackoverflow.com/questions/18817336/encrypting-a-string-with-aes-and-base64 with slight modifications)
func DecryptAES(block cipher.Block, text []byte) ([]byte, error) {

	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}

	return data, nil
}
