package main

import (
	"fmt"
	"os"
)

func CreateTerraformFiles(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}
	}
	for _, file := range []string{"main.tf", "variables.tf", "outputs.tf", "README.md"} {
		if _, err := os.Create(fmt.Sprintf("%s/%s", path, file)); err != nil {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("path parameter missing")
		os.Exit(0)
	}

	path := os.Args[1]
	CreateTerraformFiles(path)
}
