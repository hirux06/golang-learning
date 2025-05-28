// The mockup nodejs server can be found at: https://github.com/hirux06/golang-learning/tree/main/webserver
// That is where the requests are made

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const getUrl = "http://localhost:8000/get"
const postUrl = "http://localhost:8000/post"
const formUrl = "http://localhost:8000/postform"

func main() {
	// PerformGetReq(getUrl)
	// PerformPostReq(postUrl)
	PerformPostFormReq(formUrl)
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

func PerformPostReq(postUrl string) {

	// Fake JSON Payload
	requestBody := strings.NewReader(`
		{
			"coursename": "Learn Golang",
			"price": "0",
			"platform": "www.youtube.com"
		}
	`)

	response, err := http.Post(postUrl, "application/json", requestBody)

	if(err != nil){
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormReq(formUrl string) {
	data := url.Values{}

	data.Add("firstname", "saran")
	data.Add("lastname", "hiruthik")
	data.Add("email","saran@go.dev")

	response, err := http.PostForm(formUrl, data)

	if(err != nil){
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}