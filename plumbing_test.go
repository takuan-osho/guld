package main

import "testing"

func TestMakeHashObjectString(t *testing.T) {
	content := "what is up, doc?"
	blobString := MakeHashObjectString(content, "blob")
	if blobString != "bd9dbf5aae1a3862dd1526723246b20206e5fc37" {
		t.Errorf("Error occurred when generating strings of blob object.")
	}
}
