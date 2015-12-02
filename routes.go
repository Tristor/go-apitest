package main

import (
	"net/http"
)

//Route provides a simple way to build states that are compatible with
//mux.Router for defining all the parameters of the route I care about.
//There are additional parameters supported by mux.Router I do not declare.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a collection of Route objects to assist in simplifying building
//my mux.Router in router.go.
type Routes []Route

//routes is where I define all the routes I actually plan to provide in the
//API and which handlers should be used for them.  Expanding this and
//creating an http.HandlerFunc in handlers.go is most likely all that is
//necessary to extend the API.
var routes = Routes{
	Route{
		"GetUserByID",
		"GET",
		"/users/{uid}",
		GetUserByID,
	},
	Route{
		"CreateUser",
		"POST",
		"/users/",
		CreateUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/users/{uid}",
		UpdateUser,
	},
	Route{
		"RemoveUser",
		"DELETE",
		"/users/{uid}",
		DeleteUser,
	},
	Route{
		"GetGroupMembers",
		"GET",
		"/groups/{gid}",
		GetGroupMembers,
	},
	Route{
		"CreateGroup",
		"POST",
		"/groups/",
		CreateGroup,
	},
	Route{
		"UpdateGroup",
		"PUT",
		"/groups/{gid}",
		UpdateGroup,
	},
	Route{
		"DeleteGroup",
		"DELETE",
		"/groups/{gid}",
		DeleteGroup,
	},
}
