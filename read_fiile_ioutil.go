package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error){
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dat,err := ioutil.ReadFile("hello.go")
	check(err)
	fmt.Println(string(dat))
}