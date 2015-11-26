package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"GetUserByID",
		"GET",
		"/users",
		GetUserByID,
	},
	Route{
		"CreateUser",
		"POST",
		"/users",
		CreateUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/users",
		UpdateUser,
	},
	Route{
		"RemoveUser",
		"DELETE",
		"/users",
		DeleteUser,
	},
	Route{
		"GetGroupMembers",
		"GET",
		"/groups",
		GetGroupMembers,
	},
	Route{
		"CreateGroup",
		"POST",
		"/groups",
		CreateGroup,
	},
	Route{
		"UpdateGroup",
		"PUT",
		"/groups",
		UpdateGroup,
	},
	Route{
		"DeleteGroup",
		"DELETE",
		"/groups",
		DeleteGroup,
	},
}
