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

	if len(files) != 4 {
		t.Error("Want 4, got: ", len(files))
	}
	t.Cleanup(func() {
		if err := os.RemoveAll("test"); err != nil {
			log.Fatal(err)
		}
	})
}
