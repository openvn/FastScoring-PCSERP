package main

import (
	"encoding/json"
	"net/http"
)

func GetCourseList(w http.ResponseWriter, r *http.Request) {
	courseLst := database.GetCourseList("asdasd")
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(courseLst)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	course := database.GetCourse("asdasd")
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&course)
}
