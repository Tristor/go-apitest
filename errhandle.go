package main

import (
	"encoding/json"
	"net/http"
)

func ThrowClientError(w http.ResponseWriter, r *http.Request, e error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

func ThrowBadEntity(w http.ResponseWriter, r *http.Request, e error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

func ThrowNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}
