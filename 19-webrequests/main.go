package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.anapioficeandfire.com/api/characters"
func main() {
	fmt.Println("Learning Web Requets in Golang")

	response, err := http.Get(url)

	if(err != nil){
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n",response)
		
	defer response.Body.Close()


	databytes, err := io.ReadAll(response.Body)

	if(err != nil){
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}