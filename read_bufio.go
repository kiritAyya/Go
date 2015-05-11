package main

import (
	"bufio"
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
	file, err := os.Open("hello.go")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	b4, err := reader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}