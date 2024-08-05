package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// middleware , helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("Welcome to API with golang")

	// create router
	router := mux.NewRouter()

	// seeding 5 course with random data

	courses = append(courses, Course{CourseId: "1", CourseName: "Golang", CoursePrice: 100, Author: &Author{Fullname: "Gorilla", Website: "gorilla.com"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{Fullname: "Guido", Website: "guido.com"}})

	courses = append(courses, Course{CourseId: "3", CourseName: "Java", CoursePrice: 300, Author: &Author{Fullname: "James", Website: "james.com"}})

	courses = append(courses, Course{CourseId: "4", CourseName: "C++", CoursePrice: 400, Author: &Author{Fullname: "Bjarne", Website: "bjarne.com"}})

	courses = append(courses, Course{CourseId: "5", CourseName: "JavaScript", CoursePrice: 500, Author: &Author{Fullname: "Brendan", Website: "brendan.com"}})

	// Routing

	// Home Route
	router.HandleFunc("/", serveHome).Methods("GET")

	// Get All Courses Route
	router.HandleFunc("/courses", getAllCourses).Methods("GET")

	// Get One Course Route
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")

	// Create One Course Route
	router.HandleFunc("/courses", createOneCourse).Methods("POST")

	// Update One Course Route
	router.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")

	// Delete One Course Route
	router.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to port
	log.Fatal(http.ListenAndServe(":4000", router))

}

// Controllers

//Serve Home Route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API with golang</h1>"))
}

// Get All Courses Route

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get One Course Route

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)

	// loop over courses and find one with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course not found")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		json.NewEncoder(w).Encode("Please send a request body")
		return

	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		// http.Error(w, "No Data Provided", 400)
		json.NewEncoder(w).Encode("No Data Provided")
		return
	}

	// Checking if courseid and coursename exists

	for _, c := range courses {
		if c.CourseId == course.CourseId || c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	// generate random id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into course
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// updateOneCourse updates a course with a given ID.
// The function assumes that the ID is present in the request parameters.
// The course data is expected in the request body, and it is decoded into a variable of type Course.
// The function removes the old course from the courses slice and appends the new course to the slice.
// The function returns the updated course as a JSON response.

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	// Extract the ID from the request parameters
	params := mux.Vars(r)

	// Loop through the courses slice and find the course with the given ID
	for index, course := range courses {
		// If the course ID matches the requested ID, remove the course from the slice
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			// Decode the new course data from the request body
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)

			// Assign the requested ID to the updated course
			updatedCourse.CourseId = params["id"]

			// Append the updated course to the courses slice
			courses = append(courses, updatedCourse)

			// Return the updated course as a JSON response
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	// Extract the ID from the request parameters
	params := mux.Vars(r)

	// Loop through the courses slice and find the course with the given ID
	for index, course := range courses {
		// If the course ID matches the requested ID, remove the course from the slice
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// Return a success message as a JSON response
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	// If the course with the given ID is not found, return a 404 error
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course not found")
}
