/*
This module is kind of hackish, and maybe even slightly terrible.
That said, it does some useful things in a persistence layer, even if it is
only a mock.  I've enforced uniqueness, enforced foreign key constraints
guaranteeing that a group cannot contain members which do not exist as User
objects and that a User cannot be a member of a group which does not exist.

I'm sure there's a better way to do some of these things, but all I can say
is it makes me thankful that databases exist, since this would have been
prettier in SQL.
*/
package main

import (
	"errors"
	"fmt"
)

var users Users
var groups Groups

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

func DBCreateUser(u User) (User, error) {
	f := DBFindUser(u.UserID)
	if f.UserID == u.UserID {
		err := errors.New("Cannot create a duplicate user")
		return f, err
	}
	users = append(users, u)
	return u, nil
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
	userid := u.UserID
	err := DBDeleteUser(userid)
	if err != nil {
		return u, fmt.Errorf("Could not find user %s to update", userid)
	}
	u, err = DBCreateUser(u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func DBCreateGroup(gwm GroupWithMembers) (GroupWithMembers, error) {
	f := DBFindGroup(gwm.GroupName)
	if f.GroupName == gwm.GroupName {
		err := errors.New("Cannot create a duplicate group")
		return gwm, err
	}
	gwm, err := DBAddGroupMembers(gwm)
	if err != nil {
		return gwm, err
	}
	groups = append(groups, Group{gwm.GroupName})
	return gwm, nil
}

func DBFindGroup(groupname string) Group {
	for _, g := range groups {
		if g.GroupName == groupname {
			return g
		}
	}
	// return an empty Group object if not found
	return Group{}
}

func DBDeleteGroup(groupname string) error {
	for i, g := range groups {
		if g.GroupName == groupname {
			groups = append(groups[:i], groups[i+1:]...)
			err := DBDeleteGroupMembers(groupname)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("Could not find group %s to delete", groupname)
}

func DBUpdateGroup(gwm GroupWithMembers) (GroupWithMembers, error) {
	err := DBDeleteGroup(gwm.GroupName)
	if err != nil {
		return gwm, fmt.Errorf("Could not find group %s to update", gwm.GroupName)
	}
	gwm, err = DBCreateGroup(gwm)
	if err != nil {
		return gwm, err
	}
	return gwm, nil
}

func DBCheckGroups(u User) error {
	for _, g := range u.Groups {
		f := DBFindGroup(g)
		if f.GroupName != g {
			return fmt.Errorf("The group %s does not exist. Groups must be created before they can be added to users", g)
		}
	}
	return nil
}

func DBGetGroupMembers(g Group) GroupWithMembers {
	var gwm GroupWithMembers
	gwm.GroupName = g.GroupName
	for _, u := range users {
		for _, m := range u.Groups {
			if m == gwm.GroupName {
				gwm.GroupMembers = append(gwm.GroupMembers, u.UserID)
			}
		}
	}
	//If a group has no members, then the slice is empty, which is okay.
	return gwm
}

func DBAddGroupMembers(gwm GroupWithMembers) (GroupWithMembers, error) {
	for _, u := range gwm.GroupMembers {
		var UserHasGroup bool = false
		fu := DBFindUser(u)
		if fu.UserID != u {
			err := errors.New("Cannot create a group with a non-existent member")
			return gwm, err
		}
		for _, g := range fu.Groups {
			if g == gwm.GroupName {
				UserHasGroup = true
				break
			}
		}
		if UserHasGroup == true {
			continue
		}
		fu.Groups = append(fu.Groups, gwm.GroupName)
		fu, err := DBUpdateUser(fu)
		if err != nil {
			return gwm, err
		}
	}
	return gwm, nil
}

func DBDeleteGroupMembers(groupname string) error {
	gwm := DBGetGroupMembers(Group{groupname})
	for _, u := range gwm.GroupMembers {
		fu := DBFindUser(u)
		if fu.UserID != u {
			err := errors.New("Something terrible has happened and the database is inconsistent.")
			return err
		}
		for i, g := range fu.Groups {
			if g == gwm.GroupName {
				fu.Groups = append(fu.Groups[:i], fu.Groups[i+1:]...)
			}
		}
		fu, err := DBUpdateUser(fu)
		if err != nil {
			return err
		}
	}
	return nil
}
