package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("---------- Files Lecture ----------")

	content := "This will go into the file"

	file, err := os.Create("./myNewFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNillError(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNillError(err)
	fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./myNewFile.txt")
}


func readFile(fileName string)  {
	dataByte, err := os.ReadFile(fileName)
	checkNillError(err)

	fmt.Println("Text data inside is: ", string(dataByte))
}

func checkNillError(err error)  {
	if err != nil {
		panic(err)
	}
}