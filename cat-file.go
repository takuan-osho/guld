package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"regexp"
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
	re := regexp.MustCompile("(blob|tree|commit|tag)[\t\n\v\f\r ]([0-9]*)\\000(.*)")
	group := re.FindStringSubmatch(object.ContentWithHeader)
	object.Type = group[1]
	object.Content = group[3]
	object.Header = group[1] + " " + group[2] + "\000"

	hash := sha1.New()
	io.WriteString(hash, object.ContentWithHeader)
	object.ID = hex.EncodeToString(hash.Sum(nil))
	object.Dir = object.ID[:2]
	object.Name = object.ID[2:]
	return &object
}
