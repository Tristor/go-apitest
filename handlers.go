package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

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
	}

	u, err := DBCreateUser(user)
	if err != nil {
		ThrowClientError(w, r, err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		panic(err)
	}
	if err := DBDeleteUser(userid); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
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
	}
	if user.UserID != userid {
		uiderr := errors.New("The UserID of your JSON object does not match the UserID in the URL Path.  User IDs cannot be changed and are unique.")
		ThrowClientError(w, r, uiderr)
	}
	u, err := DBUpdateUser(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidatePath(w, r)
	if err != nil {
		ThrowClientError(w, r, err)
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
}

func GetGroupMembers(w http.ResponseWriter, r *http.Request) {}

func CreateGroup(w http.ResponseWriter, r *http.Request) {}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {}
