package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCatFile(t *testing.T) {
	content := `ホールは孔のマッチかっこうたちがセロを引きあげ窓たまし。それからまた生意気だたとして写真ございな。

下手ましでんんはならではテープの生意気たちのままへもがぶがぶ気の毒たまして、そこじゃマッチへ置くれ気ましう。

見つめすぎきみは虎が白いでて一番の孔の音げで出第一手団の話に笑っがいうまし。`
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

	if testObject.Content != object.Content {
		t.Errorf("Couldn't read object content")
	}

	if testObject.Type != object.Type {
		t.Errorf("Couldn't read object type")
	}

	if testObject.Header != object.Header {
		t.Errorf("Couldn't read object header")
	}

	if testObject.ContentWithHeader != object.ContentWithHeader {
		t.Errorf("Couldn't read object content with header")
	}

	if testObject.ID != object.ID {
		t.Errorf("Couldn't read object id")
	}

	if testObject.Dir != object.Dir {
		t.Errorf("Couldn't read object directory value")
	}

	if testObject.Name != object.Name {
		t.Errorf("Couldn't read object name")
	}

	if !reflect.DeepEqual(testObject.CompressedContent, object.CompressedContent) {
		t.Errorf("Couldn't read object compressed content")
	}

	os.Remove(objectPath)
	os.Remove(objectDir)
	os.Remove(tempDir)
}
