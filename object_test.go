package main

import (
	"testing"
)

func TestNewHashObject(t *testing.T) {
	content := "what is up, doc?"
	filetype := "blob"
	hashObject := NewHashObject(content, filetype)
	if hashObject.Type != filetype {
		t.Errorf("Assertion Error: File type of the object is not expected value.")
	}
	if hashObject.Content != content {
		t.Errorf("Assertion Error: Content of the object is not expected value.")
	}
	if hashObject.Header != "blob 16\000" {
		t.Errorf("Assertion Error: Header of the object is not expected value.")
	}
	if hashObject.ContentWithHeader != "blob 16\000what is up, doc?" {
		t.Errorf("Assertion Error: Content with header is not expected value.")
	}
	if hashObject.ID != "bd9dbf5aae1a3862dd1526723246b20206e5fc37" {
		t.Errorf("Assertion Error: Sha1 hashed content of the object is not expected value.")
	}
	if hashObject.Dir != "bd" {
		t.Errorf("Assertion Error: Directory name where the object should be is not expected value.")
	}
	if hashObject.Name != "9dbf5aae1a3862dd1526723246b20206e5fc37" {
		t.Errorf("Assertion Error: File name of the object is not expected value.")
	}
}
