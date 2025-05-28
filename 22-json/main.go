package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJSON()
	DecodeJSON()
}

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJSON() {

	availableCourses := []courses{
		{"Golang Bootcamp", 299, "udemy.com", "abc123", []string{"go", "backend"}},
		{"Angular Bootcamp", 299, "udemy.com", "abc123", []string{"angular", "frontend", "js"}},
		{"Golang Bootcamp", 299, "udemy.com", "abc123", nil},
	}

	// Package this data as JSON
	// finalJson, err := json.Marshal(availableCourses)
	finallJson, err := json.MarshalIndent(availableCourses,"","\t") // Indent the json based on the third parameter

	if(err != nil){
		panic(err)
	}

	fmt.Printf("JSON data: %s\n", finallJson)
}

func DecodeJSON()  {
	jsonDataFromDB := []byte(`
		{
			"coursename": "Golang Bootcamp",
			"Price": 299,
			"website": "udemy.com",
			"tags": ["go", "backend"]
        }
	`)

	var availableCourses courses;

	checkValid := json.Valid(jsonDataFromDB)

	if(checkValid){
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromDB, &availableCourses)
		fmt.Printf("%#v\n", availableCourses)
	}else {
		fmt.Printf("JSON is not valid")
	}

	// some cases where we just want to add some data to key value

	var myCourseData map[string]interface{}

	json.Unmarshal(jsonDataFromDB, &myCourseData)
	fmt.Printf("%#v\n", myCourseData)

	for k, v := range myCourseData{
		fmt.Printf("Key is %v and value is: %v and type is: %T\n",k, v, v )
	}
}