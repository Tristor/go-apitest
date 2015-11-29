package main

import "fmt"

var users Users

// Generate seed data

func init() {
}

func DBFindUser(userid string) User {
	for _, u := range users {
		if u.UserID == userid {
			return u
		}
	}
	// return an empty User object if not found
	return User{}
}

func DBCreateUser(u User) User {
	users = append(users, u)
	return u
}

func DBDeleteUser(userid string) error {
	for i, u := range users {
		if u.UserID == userid {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find user %s to delete", userid)
}

func DBUpdateUser(u User) (User, error) {
	userid = u.UserID
	err := DBDeleteUser(userid)
	if err != nil {
		return u, fmt.Errorf("Could not find user %s to update", userid)
	}
	u = DBCreateUser(u)
	return u, err
}
