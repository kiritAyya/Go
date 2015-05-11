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
	file,err := os.Open("hello.go")
	check(err)
	defer file.Close()

	s1 := make([]byte, 5)
	n1,err := file.Read(s1)
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(s1))

	o2, err := file.Seek(6,0)
	check(err)
	s2 := make([]byte, 2)
	n2, err := file.Read(s2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(s2))
}