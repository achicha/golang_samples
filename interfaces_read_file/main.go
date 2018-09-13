package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	// get a file name from CMD args
	fileName := os.Args[1]

	// find executable DIR name
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// read file
	file, err := ioutil.ReadFile(path.Join(dir, fileName))
	if err != nil {
		log.Fatal(err)
	}

	// print output
	fmt.Println(string(file))
}
