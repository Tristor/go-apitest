package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

//CreateUser implements the http.HandlerFunc interface to process POST
//requests to the API per the specification.  It handles all returned
//errors by responding with appropriate HTTP errors and returns a valid
//Location: header along with an HTTP 201 and the created User if it
//succeeds.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		ThrowBadEntity(w, r, err)
		return
	}
	if err := DBCheckGroups(user); err != nil {
		ThrowClientError(w, r, err)
		return
	}
	u, err := DBCreateUser(user)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", r.URL.Path+u.UserID)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

//DeleteUser implements the http.HandlerFunc interface to process DELETE
//requests to the API.  We return a 404 rather than a 400 if
//DBDeleteUser fails because we have awareness that DBDeleteUser only
//returns an error if it cannot find the user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	if err := DBDeleteUser(userid); err != nil {
		ThrowNotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

//UpdateUser implements the http.HandlerFunc interface to process PUT
//requests to the API.  It also provides additional error handling
//prior to calling DBUpdateUser.  Just as in CreateUser I call
//DBCheckGroups here so that I can catch its error conditions early.  As
//per the spec, this returns a 404 if the user doesn't exist.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		ThrowBadEntity(w, r, err)
		return
	}
	if user.UserID != userid {
		uiderr := errors.New("The UserID of your JSON object does not match the UserID in the URL Path.  User IDs cannot be changed and are unique.")
		ThrowClientError(w, r, uiderr)
		return
	}
	u := DBFindUser(userid)
	if u.UserID != userid {
		ThrowNotFound(w, r)
		return
	}
	if err := DBCheckGroups(user); err != nil {
		ThrowClientError(w, r, err)
		return
	}
	u, err = DBUpdateUser(user)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

//GetUserByID implements the http.HandlerFunc interface to process GET
//requests to the API. It's mostly a wrapper for DBFindUser. If the
//User exists it is returned, otherwise it returns an HTTP 404.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	u := DBFindUser(userid)
	if u.UserID == userid {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(u); err != nil {
			panic(err)
		}
		return
	}
	ThrowNotFound(w, r)
}

//GetGroupMembers implements the http.HandlerFunc interface to process GET
//requests to the API.  If a Group exists it requests a GroupWithMembers
//to be generated and returns it.  Otherwise it returns an HTTP 404.
func GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	groupname, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	g := DBFindGroup(groupname)
	if g.GroupName == groupname {
		gwm := DBGetGroupMembers(g)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(gwm); err != nil {
			panic(err)
		}
		return
	}
	ThrowNotFound(w, r)
}

//CreateGroup implements the http.HandlerFunc interface to process POST
//requests to the API.  It expects a GroupWithMembers but can also handle
//a Group being passed as the JSON object due to the flexibility of the
//json.Unmarshal function.  It returns a GroupWithMembers if it succeeds
//otherwise it returns an appropriate HTTP error.
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var gwm GroupWithMembers
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &gwm); err != nil {
		ThrowBadEntity(w, r, err)
		return
	}
	g, err := DBCreateGroup(gwm)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", r.URL.Path+g.GroupName)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}
}

//UpdateGroup implements the http.HandlerFunc interface to process PUT
//requests to the API.  It expects a GroupWithMembers in JSON and performs
//additional error-checking before calling DBUpdateGroup. Per the spec if
//the group doesn't exist it returns an HTTP 404.  It returns otherwise
//appropriate errors or a GroupWithMembers on success.
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	groupname, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	var gwm GroupWithMembers
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &gwm); err != nil {
		ThrowBadEntity(w, r, err)
		return
	}
	if gwm.GroupName != groupname {
		uiderr := errors.New("The name string in your JSON object does not match the group name provided in the URL")
		ThrowClientError(w, r, uiderr)
		return
	}
	if g := DBFindGroup(groupname); g.GroupName != groupname {
		ThrowNotFound(w, r)
		return
	}
	gwm, err = DBUpdateGroup(gwm)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(gwm); err != nil {
		panic(err)
	}
}

//DeleteGroup implements the http.HandlerFunc interface to process DELETE
//requests to the API.  Unlike deleting users, deleting groups has more
//possible failure conditions, and those are handled here as client errors.
//Otherwise, a 204 is returned with no body.
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupname, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
		return
	}
	if err := DBDeleteGroup(groupname); err != nil {
		ThrowClientError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
