package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"

	"github.com/codegangsta/cli"
)

var CmdCatFileFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "t",
		Usage: "show object type",
	},
	cli.BoolFlag{
		Name:  "s",
		Usage: "show object size",
	},
	cli.BoolFlag{
		Name:  "p",
		Usage: "pretty-print object's content",
	},
}

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

func CatFileAction(c *cli.Context) {
	object := CatFile(filepath.Join(ObjectDir, c.Args()[0][:2], c.Args()[0][2:]))

	if c.Bool("t") {
		fmt.Println(object.Type)
	}

	if c.Bool("p") {
		fmt.Println(object.Content)
	}
}
