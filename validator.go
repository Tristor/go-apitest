package main

import (
	"errors"
	"net/http"
	"regexp"
)

//ValidPath invokes regexp.MustCompile which causes a fatal error
//if the regexp fails.  This is the basis for enforcing the requirement
//that User.UserID and Group.GroupName must be alpha only by checking it
//in the URL path on everything but POST.  I don't currently enforce
//this elsewhere, but may do so eventually.
var ValidPath = regexp.MustCompile("^/(users|groups)/([a-zA-Z]+)$")

//ValidatePath takes w and r and returns a string and an error.
//Using the ValidPath regexp, it verifies that the identifier provided
//in the URL as part of the API request is valid.  Once it has done so
//it returns the identifier and any errors to the caller.  Used by our
//HTTP handlers to get the User.UserID and Group.GroupName to operate on
//from the URL parameters.
func ValidatePath(w http.ResponseWriter, r *http.Request) (string, error) {
	m := ValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid ID. IDs must only contain alpha characters.")
	}
	return m[2], nil
}
