package main

import "testing"

func TestGetStringsForHashObject(t *testing.T) {
	content := "what is up, doc?"
	contentWithHeader, blobString := GetStringsForHashObject(content, "blob")
	if contentWithHeader != "blob 16\000what is up, doc?" {
		t.Errorf("Error occurred when getting content with header.")
	}
	if blobString != "bd9dbf5aae1a3862dd1526723246b20206e5fc37" {
		t.Errorf("Error occurred when getting strings for blob object.")
	}
}
