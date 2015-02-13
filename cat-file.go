package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"io/ioutil"
	"log"
)

func CatFile(filePath string) *HashObject {
	object := HashObject{}
	compressedContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	object.CompressedContent = compressedContent

	b := bytes.NewReader(compressedContent)

	r, err := zlib.NewReader(b)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer

	io.Copy(&out, r)
	r.Close()

	object.ContentWithHeader = string(out.Bytes()[:])
	return &object
}
