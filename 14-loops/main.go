package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for d :=0; d<len(days); d++{
		fmt.Println(days[d])
	}

	for d :=range days{
		fmt.Println(days[d])
	}

	for index, day := range days{
		fmt.Printf("Index is %v and day is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10{

		if rougueValue == 2{
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++;
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++;
	}

	lco:
		fmt.Println("Jumping here!!")
}