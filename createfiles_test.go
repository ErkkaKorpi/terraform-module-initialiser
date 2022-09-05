package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateTerraformFiles(t *testing.T) {
	CreateTerraformFiles("test")

	files, err := ioutil.ReadDir("test")
	if err != nil {
		log.Fatal(err)
	}

	if len(files) != 3 {
		t.Error("Want 3, got: ", len(files))
	}
	t.Cleanup(func() {
		if err := os.RemoveAll("test"); err != nil {
			log.Fatal(err)
		}
	})
}
