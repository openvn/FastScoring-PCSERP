package main

import (
	"fmt"
	"github.com/openvn/FastScoring-PCSERP/db"
	"log"
	"net/http"
)

var database db.DBContext = &db.TestDBContext{}

type ErrorResponse struct {
	Code   int
	Detail string
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Code: %d %s", e.Code, e.Detail)
}

func main() {
	http.HandleFunc("/GetCourseList", GetCourseList)
	http.HandleFunc("/GetCourse", GetCourse)
	http.HandleFunc("/Login", Login)

	log.Fatal(http.ListenAndServe(":12345", nil))
}
