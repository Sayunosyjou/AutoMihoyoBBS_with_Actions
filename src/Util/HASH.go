package Util

import (
	"crypto/sha256"
	"crypto/sha512"
)

func GetSha256(Data []byte) []byte {
	hash256 := sha256.New()
	hash256.Write(Data)
	return hash256.Sum(nil)
}

func GetSha512(Data []byte) []byte {
	hash512 := sha512.New()
	hash512.Write(Data)
	return hash512.Sum(nil)
}
