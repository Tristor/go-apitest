package main

import (
	"encoding/json"
	"net/http"
)

//ThrowClientError takes in an error object and returns an HTTP 400 with
//the message encoded to a JSON object. This is used often in the handlers
//so was standardized here.
func ThrowClientError(w http.ResponseWriter, r *http.Request, e error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

//ThrowBadEntity takes in an error object and returns an HTTP 422 with
//the message encoded to a JSON object.  Is most often used when the
//client sends an object we were unable to json.Unmarshal.
func ThrowBadEntity(w http.ResponseWriter, r *http.Request, e error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		panic(err)
	}
}

//ThrowNotFound immediately returns an HTTP 404 and then returns.
//Used when a DBFindUser or DBFindGroup operation fails on a GET call.
func ThrowNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}
