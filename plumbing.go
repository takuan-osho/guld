package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
)

func MakeHashObjectString(content string, filetype string) string {
	header := filetype + " " + strconv.Itoa(len(content)) + "\000"
	store := header + content
	hash := sha1.New()
	io.WriteString(hash, store)
	return hex.EncodeToString(hash.Sum(nil))
}
