package main

import (
	"errors"
	"fmt"
)

var users Users
var groups Groups

//DBFindUser takes a string and returns a User.  It searches through
//the Users object which comprises part of the mock database.  If its
//unable to find a user, it returns an empty object instead.  This is
//not an error condition in all cases, so is handled by the caller.
func DBFindUser(userid string) User {
	for _, u := range users {
		if u.UserID == userid {
			return u
		}
	}
	// return an empty User object if not found
	return User{}
}

//DBCreateUser takes a User and returns a User object and an
//error value.  If everything goes okay, the User returned should match
//the user passed in.  If it finds that a user already exists in the mock
//database, it will return the User object that was there, rather than
//what was passed.  This enables the caller some additional insight into
//the result if it chooses to use it.
func DBCreateUser(u User) (User, error) {
	f := DBFindUser(u.UserID)
	if f.UserID == u.UserID {
		err := errors.New("Cannot create a duplicate user")
		return f, err
	}
	users = append(users, u)
	return u, nil
}

//DBDeleteUser takes a string and returns an error.  The only error
//condition is if the user passed in doesn't exist.  This should be handled
//by the caller. Since we generate GroupWithMembers objects on the fly
//we are able to maintain internal consistency for User.Groups and
//GroupWithMembers.GroupMembers simply by deleting the User object from
//the mock database.
func DBDeleteUser(userid string) error {
	for i, u := range users {
		if u.UserID == userid {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find user %s to delete", userid)
}

//DBUpdateUser takes a User and returns a User and an error.
//Our update is pretty hacky, but workable.  It simply calls DBDeleteUser
//followed by DBCreateUser with some error checking.  The User object
//returned should always be the one passed in.
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

//DBCreateGroup takes a GroupWithMembers and returns a GroupWithMembers
//and an error.  You'll see we call DBAddGroupMembers to enforce
//consistency in the mock database if a GroupWithMembers.GroupMembers
//value contains users on creation, since that state is stored in the
//User objects.
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

//DBFindGroup takes a string and returns a Group.  If it's unable
//to find the group it returns an empty Group object.  This is not
//necessarily an error and should be handled by the caller.
func DBFindGroup(groupname string) Group {
	for _, g := range groups {
		if g.GroupName == groupname {
			return g
		}
	}
	// return an empty Group object if not found
	return Group{}
}

//DBDeleteGroup takes a string and returns an error.  Unlike deleting
//a user, deleting a group may have other error conditions and can return
//them.  The caller should be aware and handling possible errors.  This
//primarily relies on the DBDeleteGroupMembers function to clean up state
//in stored User objects so that there is no inconsistencies in the mock
//database.  Pretty hackish, but hey, it works.  Makes you yearn for SQL.
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

//DBUpdateGroup takes a GroupWithMembers and returns a GroupWithMembers
//and an error.  The update process for groups is basically the same as
//for users.  We simply delete and it and recreate it.  This is a pretty
//expensive operation due to how I'm enforcing consistency in the mock
//database, but for our purposes here should suffice.  If performance
//mattered, this would be an obvious choice for refactoring.
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

//DBCheckGroups takes a User and returns an error.  This function is used
//to simply verify that a User object does not contain any groups which
//don't exist.  This is helpful during User creation and updates.
func DBCheckGroups(u User) error {
	for _, g := range u.Groups {
		f := DBFindGroup(g)
		if f.GroupName != g {
			return fmt.Errorf("The group %s does not exist. Groups must be created before they can be added to users", g)
		}
	}
	return nil
}

//DBGetGroupMembers takes a Group and returns a GroupWithMembers.  This
//is the real workhorse behind how I'm enforcing the foreign constraints
//between GroupWithMembers.GroupMembers and User.Groups.  All of the state
//is stored in the mock database in the User objects and then a
//GroupWithMembers is generated on the fly with this function whenever
//it is needed.  This function is used extensively both within the database
//and within the HTTP handlers. Due to loop nesting this isn't very fast
//but it's better than the alternative implementations I could think of for
//enforcing consistency.
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

//DBAddGroupMembers takes a GroupWithMembers and returns a GroupWithMembers
//and an error.  The purpose of this function is to support the API
//requirement that you can manipulate group memberships with a PUT to
///groups/{groupname}.  This first verifies that all members provided are
//existing users, and then updates all relevant User objects in the mock
//database to contain the gwm.GroupName in their User.Groups.  Again, due
//to loop nesting, this is not as performant as it could be, but it provides
//a pretty strong consistency guarantee for our mock database.
func DBAddGroupMembers(gwm GroupWithMembers) (GroupWithMembers, error) {
	for _, u := range gwm.GroupMembers {
		var UserHasGroup = false
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

//DBDeleteGroupMembers takes a string and returns an error.  This is
//essentially the reverse of DBAddGroupMembers.  It finds all the User
//objects which contain a Group.GroupName in their User.Groups matching
//the provided string and remove it.  One of the errors I check for here
//should never happen.  It would only be possible if somehow the database
//were manipulate outside the bounds of the HTTP handlers.  It's pretty
//much a panic() situation, but is being passed to the caller for them to
//handle it.
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
