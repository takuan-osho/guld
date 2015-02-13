package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestCatFile(t *testing.T) {
	content := "what is up, doc?"
	filetype := "blob"
	object := NewHashObject(content, filetype)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tempDir, err := ioutil.TempDir(currentDir, "temp")
	if err != nil {
		log.Fatal(err)
	}

	objectDir := filepath.Join(tempDir, object.Dir)
	err = os.MkdirAll(objectDir, 0755)
	if err != nil {
		os.Remove(tempDir)
		log.Fatal(err)
	}

	objectPath := filepath.Join(objectDir, object.Name)
	err = ioutil.WriteFile(objectPath, object.CompressedContent, 0644)
	if err != nil {
		os.Remove(objectDir)
		os.Remove(tempDir)
		t.Errorf("Couldn't write temp file.")
	}

	testObject := CatFile(objectPath)
	if testObject.ContentWithHeader != object.ContentWithHeader {
		os.Remove(objectPath)
		os.Remove(objectDir)
		os.Remove(tempDir)
		t.Errorf("Couldn't read object compressed content")
	}

	os.Remove(objectPath)
	os.Remove(objectDir)
	os.Remove(tempDir)
}
