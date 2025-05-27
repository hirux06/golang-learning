// The mockup nodejs server can be found at: https://github.com/hirux06/golang-learning/tree/main/webserver
// That is where the requests are made

package main

import (
	"fmt"
	"io"
	"net/http"
)

const getUrl = "http://localhost:8000/get"

func main() {
	PerformGetReq(getUrl)
}
func PerformGetReq(getUrl string) {
	response, err :=  http.Get(getUrl)

	if(err != nil){
		fmt.Println("Error occurred while making GET request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Content: ", string(content))
}