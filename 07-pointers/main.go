package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int;
	// fmt.Println("Value of pointer is: ", ptr)

	myNum := 23

	var ptr  = &myNum

	fmt.Println("Value of actual pointer is: ", *ptr)
	fmt.Println("Value of actual pointer is: ", ptr)

	*ptr = *ptr * 2;
	fmt.Println("New value is: ",myNum)

}