package main

import "fmt"

func main() {
	fmt.Println("Inside Main Function")
	greet()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, proResultMsg := proAdded(3, 5,1, 4, 6, 7)
	fmt.Println("Result is: ", proResult, proResultMsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdded(values ...int) (int, string) {
	total := 0

	for _, val := range values{
		total += val;
	}

	return total, "Hi, from proResult func()"
}

func greet(){
	fmt.Println("Vanakkam from greet fn.()")
}
