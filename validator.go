package main

import (
	"errors"
	"net/http"
	"regexp"
)

var ValidPath = regexp.MustCompile("^/(users|groups)/([a-zA-Z]+)$")

func ValidatePath(w http.ResponseWriter, r *http.Request) (string, error) {
	m := ValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid ID. IDs must only contain alpha characters.")
	}
	return m[2], nil
}
