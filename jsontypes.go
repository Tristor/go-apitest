package main

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	UserID    string   `json:"userid"`
	Groups    []string `json:"groups"`
}

type Users []User

type Group struct {
	GroupName string `json:"name"`
}

type Groups []Group

type GroupWithMembers struct {
	GroupName    string   `json:"name"`
	GroupMembers []string `json:"members"`
}
