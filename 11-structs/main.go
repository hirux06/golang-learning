package main

import "fmt"

func main() {
	saran := User{"saran", "saran@go.dev", true, 19}
	fmt.Println(saran)
	fmt.Printf("Saran details are %+v\n", saran)
	fmt.Printf("Name is %v and email is %v\n", saran.Name, saran.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}