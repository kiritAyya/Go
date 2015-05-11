package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dir, err := os.Open(".")
	check(err)
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	check(err)
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}