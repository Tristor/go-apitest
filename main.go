package main

import (
	"github.com/dougblack/sleepy"
	"net/http"
	"net/url"
)

type User struct {
	firstName string   `json:"first_name"`
	lastName  string   `json:"last_name"`
	userID    string   `json:"userid"`
	groups    []*Group `json:"groups"`
}

type Group struct {
	groupName    string  `json:"name"`
	groupMembers []*User `json:"group_members"`
}
