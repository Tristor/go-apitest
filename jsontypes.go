package main

//User is a simple data structure to assist in marshalling and
//unmarshalling data to JSON for our API.
type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	UserID    string   `json:"userid"`
	Groups    []string `json:"groups"`
}

//Users is a slice of User objects.  The primary use is to store
//all of our users in the mock database.
type Users []User

//Group is defined by its name as its only required value. As such
//we store the name only in its basic form.
type Group struct {
	GroupName string `json:"name"`
}

//Groups is a slice of Group objects.  The primary use is to store
//all of our groups in the mock database.
type Groups []Group

//GroupWithMembers is present to help us marshal an appropriate JSON
//response for GET and PUT calls to the API.  The GroupMembers value
//is generated on the fly as needed within the mock database.
type GroupWithMembers struct {
	GroupName    string   `json:"name"`
	GroupMembers []string `json:"members"`
}
