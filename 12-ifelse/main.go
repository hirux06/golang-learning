package main

import "fmt"

func main() {
	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Excatly 10 Users"
	}

	fmt.Println(result)
	fmt.Println(9%2)
}