package main

import (
	"github.com/openvn/FastScoring-PCSERP/db"
	"log"
	"net/http"
)

var database db.DBContext = &db.TestDBContext{}

func main() {
	http.HandleFunc("/GetCourseList", GetCourseList)
	http.HandleFunc("/GetCourse", GetCourse)

	log.Fatal(http.ListenAndServe(":12345", nil))
}
