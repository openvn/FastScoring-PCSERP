package main

import (
	"encoding/json"
	"github.com/openvn/FastScoring-PCSERP/auth"
	"net/http"
)

func GetCourseList(w http.ResponseWriter, r *http.Request) {
	if err := auth.ValidToken(r.Header.Get("Token")); err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		json.NewEncoder(w).Encode(
			&ErrorResponse{http.StatusForbidden, err.Error()})
	}

	courseLst := database.GetCourseList("asdasd")
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(courseLst)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	if err := auth.ValidToken(r.Header.Get("Token")); err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		json.NewEncoder(w).Encode(
			&ErrorResponse{http.StatusForbidden, err.Error()})
	}

	course := database.GetCourse("asdasd")
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(&course)
}
