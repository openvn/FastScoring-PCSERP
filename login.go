package main

import (
	"encoding/json"
	"github.com/openvn/FastScoring-PCSERP/auth"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var c auth.Credential

	err := json.NewDecoder(r.Body).Decode(&c)
	r.Body.Close()

	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			&ErrorResponse{http.StatusBadRequest, err.Error()})
		return
	}

	token, err := auth.Login(c)
	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
		json.NewEncoder(w).Encode(
			&ErrorResponse{http.StatusForbidden, err.Error()})
		return
	}

	w.Header().Add("Token", token)
}
