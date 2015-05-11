package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file,err := os.Open("hello.go")
	check(err)
	defer file.Close()

	stat,err := file.Stat()
	check(err)

	bs := make([]byte, stat.Size())
	_,err = file.Read(bs)
	check(err)

	str := string(bs)
	fmt.Println(str)
}