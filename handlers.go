package main

import (
	"io"
	"io/ioutil"
	"json"
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	u, err := DBCreateUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidPath(w, r)
	if err != nil {
		panic(err)
	}
	if err := DBDeleteUser(userid); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userid, err := ValidPath(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if user.UserID != userid {
		uiderr := errors.New("The UserID of your JSON object does not match the UserID in the URL Path.  User IDs cannot be changed and are unique.")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(uiderr); err != nil {
			panic(err)
		}
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
	userid, err := ValidPath(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
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
