package main

import "fmt"

const LoginToken string = "blah-blah-blah" // Public variable

func main() {
	var username string = "Saran Hiruthk M"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)


	fmt.Println("------------------------------------")


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	fmt.Println("------------------------------------")
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)
	
	fmt.Println("------------------------------------")
	var smallFloat float32 = 255.67898798
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// default values and some aliases
	
	fmt.Println("------------------------------------")
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000  // but this cannot declared outside of functions
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("%T", LoginToken)

}