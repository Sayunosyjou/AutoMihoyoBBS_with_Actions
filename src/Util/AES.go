package Util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const AESKeyLength = 32
const AESNonceLength = 12

type aesStruct struct {
	Token []byte
	Nonce []byte
	Key   []byte
}

func (receiver *aesStruct) SetNonce(Nonce []byte) {
	receiver.Nonce = Nonce
}

func (receiver *aesStruct) SetKey(Key []byte) {
	receiver.Key = Key
}

func (receiver *aesStruct) SetToken(Token []byte) {
	receiver.Token = Token
}

func (receiver *aesStruct) getRandData(Length int) ([]byte, error) {
	data := make([]byte, Length)
	_, err := io.ReadFull(rand.Reader, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (receiver *aesStruct) NewNonce() error {
	var err error
	receiver.Nonce, err = receiver.getRandData(AESNonceLength)
	return err
}

func (receiver *aesStruct) NewKey() error {
	var err error
	receiver.Key, err = receiver.getRandData(AESKeyLength)
	return err
}

func (receiver *aesStruct) Encrypt(Plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(receiver.Key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, receiver.Nonce, Plaintext, receiver.Token)

	return ciphertext, nil
}

func (receiver *aesStruct) EncryptToBase64(Plaintext []byte) (string, error) {
	Ciphertext, err := receiver.Encrypt(Plaintext)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(Ciphertext), nil
}

func (receiver *aesStruct) Decrypt(Ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(receiver.Key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	Plaintext, err := aesgcm.Open(nil, receiver.Nonce, Ciphertext, receiver.Token)
	if err != nil {
		return nil, err
	}

	return Plaintext, nil
}

func (receiver *aesStruct) DecryptFromBase64(Base64Ciphertext string) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(Base64Ciphertext)
	if err != nil {
		return nil, err
	}

	plaintext, err := receiver.Decrypt(ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func NewAES() aesStruct {
	var AES aesStruct
	return AES
}
