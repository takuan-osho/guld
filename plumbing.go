package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
)

func GetStringsForHashObject(content string, filetype string) (contentWithHeader, hashObjectString string) {
	header := filetype + " " + strconv.Itoa(len(content)) + "\000"
	contentWithHeader = header + content

	hash := sha1.New()
	io.WriteString(hash, contentWithHeader)
	hashObjectString = hex.EncodeToString(hash.Sum(nil))

	return
}
