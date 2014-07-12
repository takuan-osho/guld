package main

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	BaseDirBaseName = "guld"
	BaseDir         = "." + BaseDirBaseName
)

func Init() {
	err := os.Mkdir(BaseDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(BaseDir)
	if err != nil {
		log.Fatal(err)
	}

	initFiles := []string{"HEAD", "config", "description", "index"}
	for _, fileName := range initFiles {
		err = ioutil.WriteFile(fileName, nil, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	initDirs := []string{"branches", "hooks", "info", "objects", "refs"}
	for _, dirName := range initDirs {
		err = os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	initSubDirs := []string{"objects/info", "objects/pack", "refs/head", "refs/tags"}
	for _, subDirName := range initSubDirs {
		err = os.MkdirAll(subDirName, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	Init()
}
