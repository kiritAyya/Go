package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file,err := os.Create("test.txt")
	check(err)
	
	defer file.Close()

	file.WriteString("test")
	file.Sync()
}