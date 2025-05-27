package main

import "fmt"

func main() {
	saran := User{"saran", "saran@go.dev", false, 19}
	fmt.Println(saran)
	saran.GetStatus()
	saran.NewMail()
		fmt.Println(saran)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	if(u.Status){
		fmt.Println("Active")
	}else {
		fmt.Println("Inactive")
	}
}

func (u User) NewMail() {
	u.Email = "new@go.dev"
	fmt.Printf("Email of user is: %v", u.Email)
}