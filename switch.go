package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)

	switch age {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fallthrough
	case 9:
		fallthrough

	case 10:
		fallthrough

	case 11:
		fallthrough

	case 12:
		{
			fmt.Println(name, "is a kid.")
		}

	case 13:
		fallthrough
	case 14:
		fallthrough
	case 15:
		fallthrough
	case 16:
		fallthrough
	case 17:
		fallthrough
	case 18:
		fallthrough

	case 19:
		{
			fmt.Println(name, "is a teenager.")
		}

	case 20:
		{
			fmt.Println(name, "is an adult.")
		}

	default:
		{
			fmt.Println(name, ", what are you?!?!")
		}
	}
}
