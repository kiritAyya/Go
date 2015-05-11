package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	resp, err := http.Get("http://www.google.com/robots.txt")
	check(err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	body := make([]byte, 100000)

	body, err = ioutil.ReadAll(resp.Body)
	check(err)

	for _, v := range body {
		fmt.Printf("%s", string(v))
	}
}
