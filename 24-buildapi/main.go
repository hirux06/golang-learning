package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB

var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName  == ""
	return c.CourseName  == ""
}

func main() {
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice: 299, Author: &Author{FullName: "Saran", Website: "youtube.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Python", CoursePrice: 599, Author: &Author{FullName: "Hiruthik", Website: "udemy.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}


// controllers

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to Home Route</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	
	for _, course := range courses{
		if(course.CourseId == params["id"]){
			json.NewEncoder(w).Encode(course)
			return;
		}
	}

	json.NewEncoder(w).Encode("No course found with given id: "+ params["id"])
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, val := range courses{
		if course.CourseName  == val.CourseName{
			json.NewEncoder(w).Encode("Course Name already exists")
			return
		}
	}

	// genrate unique courseid
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, get id, remove, add with id

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}


func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, get id, delete

	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}