package main

import (
	"fmt"
	"net/http"

	"github.com/hirux06/mongoap/router"
)

func main() {
	fmt.Println("Server is getting started..")
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000...")
}