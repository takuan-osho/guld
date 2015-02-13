package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
)

type HashObject struct {
	Content           string
	Type              string
	Header            string
	ContentWithHeader string
	ID                string
	Dir               string
	Name              string
	CompressedContent []byte
}

func NewHashObject(content string, filetype string) *HashObject {
	hashObject := HashObject{Content: content, Type: filetype}
	hashObject.Header = filetype + " " + strconv.Itoa(len(content)) + "\000"
	hashObject.ContentWithHeader = hashObject.Header + hashObject.Content

	hash := sha1.New()
	io.WriteString(hash, hashObject.ContentWithHeader)
	hashObject.ID = hex.EncodeToString(hash.Sum(nil))
	hashObject.Dir = hashObject.ID[:2]
	hashObject.Name = hashObject.ID[2:]

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(hashObject.ContentWithHeader))
	w.Close()
	hashObject.CompressedContent = b.Bytes()

	return &hashObject
}
