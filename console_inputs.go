package main

import "fmt"

func main() {
	var name string

	fmt.Println("Enter username: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello,", name)
}
