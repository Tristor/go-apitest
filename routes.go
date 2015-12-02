package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
