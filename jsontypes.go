package main

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    string `json:"userid"`
	Groups    Groups `json:"groups"`
}

type Users []User

type Group struct {
	GroupName    string `json:"name"`
	GroupMembers Users  `json:"members"`
}

type Groups []Group
