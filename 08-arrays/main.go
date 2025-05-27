package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays Conecpt")

	var fruitList [4]string;

	fruitList[0] = "Apple";
	fruitList[1] = "Banana";
	fruitList[3] = "Peach"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))


	var vegList = [3]string{"potato", "tomato", "beans"}

	fmt.Println(vegList)
	fmt.Println(len(vegList))
}