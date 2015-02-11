package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func GetStringsForHashObject(content string, filetype string) (contentWithHeader, sha1String string) {
	header := filetype + " " + strconv.Itoa(len(content)) + "\000"
	contentWithHeader = header + content

	hash := sha1.New()
	io.WriteString(hash, contentWithHeader)
	sha1String = hex.EncodeToString(hash.Sum(nil))

	return
}

func WriteHashObject(content string, filetype string, repositoryPath string) {
	contentWithHeader, sha1String := GetStringsForHashObject(content, filetype)
	dirPath := repositoryPath + "/objects/" + sha1String[:2]
	filePath := dirPath + "/" + sha1String[2:]

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("Couldn't make directory for blob objects.")
	}

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte(contentWithHeader))
	w.Close()

	err = ioutil.WriteFile(filePath, b.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Couldn't make a blob object.")
	}
}
