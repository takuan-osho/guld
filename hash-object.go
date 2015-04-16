package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

var CmdHashObjectFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "t",
		Usage: "object type",
	},
	cli.BoolFlag{
		Name:  "w",
		Usage: "write the object into the object database",
	},
	cli.StringFlag{
		Name:  "stdin",
		Usage: "read the object from stdin",
	},
}

func WriteHashObject(hashObject *HashObject, repositoryPath string) {
	dirPath := filepath.Join(repositoryPath, "objects", hashObject.Dir)
	filePath := filepath.Join(dirPath, hashObject.Name)

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filePath, hashObject.CompressedContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func HashObjectAction(c *cli.Context) {
	var objectType, objectContent string

	if c.String("t") == "" {
		objectType = "blob"
	} else {
		objectType = c.String("t")
	}

	if c.String("stdin") != "" {
		objectContent = c.String("stdin")
	} else {
		if len(c.Args()) <= 0 {
			log.Fatalf("Input file you want to compress.")
		}
		fileName := c.Args()[0]
		byteContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		objectContent = string(byteContent)
	}

	object := NewHashObject(objectContent, objectType)

	if c.Bool("w") {
		WriteHashObject(object, BaseDir)
	}

	fmt.Println(object.ID)
}
