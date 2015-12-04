# go-apitest

## Synopsis
This is just an implementation of an API for a programming interview assignment in Google Golang

The assignment is documented in this repository as code_test.md

This project assumes you have a valid Golang environment setup on your
local system before use.

## Installation and Usage

```bash
$ go get github.com/tduzan/go-apitest
$ $GOPATH/bin/go-apitest
```

To interact with the API use http://localhost:8090/

API matches the specifications in the assignment.

All exported functions and types are documented inline in go doc
compatible format.  Please use `go doc -cmd` in the repository to see a
list of symbols and the package description.

## Test Procedure

I didn't write unit tests because I was already a bit behind on this
assignment.  I didn't get nearly as much time to work on this as I'd
hoped over the holiday.  Instead I manually tested using curl from the
command line.

POST to /users/

```bash
curl -H "Content-Type: application/json" -H "Accept: application/json"
-X POST -i -d '{"first_name":"john", "last_name":"smith",
"userid":"jsmith", "groups":[]}' http://localhost:8090/users/
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Location: /users/jsmith
Date: Fri, 04 Dec 2015 03:59:28 GMT
Content-Length: 72

{"first_name":"john","last_name":"smith","userid":"jsmith","groups":[]}
```

POST to /groups/

```bash
curl -H "Content-Type: application/json" -H "Accept: application/json"
-i -X POST -d '{"name":"admin", "members":["jsmith"]}'
http://localhost:8090/groups/
HTTP/1.1 201 Created
Content-Type: application/json; charset=UTF-8
Location: /groups/admin
Date: Fri, 04 Dec 2015 04:01:14 GMT
Content-Length: 38

{"name":"admin","members":["jsmith"]}
```

GET from /users/jsmith
```bash
curl -H "Accept: application/json" -i -X GET
http://localhost:8090/users/jsmith
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:02:40 GMT
Content-Length: 79

{"first_name":"john","last_name":"smith","userid":"jsmith","groups":["admin"]}
```

GET from /users/notexist
```bash
curl -H "Accept: application/json" -i -X GET
http://localhost:8090/users/notexist
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:03:15 GMT
Content-Length: 0
```
GET from /groups/admin
```bash
curl -H "Accept: application/json" -i -X GET
http://localhost:8090/groups/admin
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:03:44 GMT
Content-Length: 38

{"name":"admin","members":["jsmith"]}
```

GET from /groups/notexist

```bash
curl -H "Accept: application/json" -i -X GET
http://localhost:8090/groups/notexist
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:04:10 GMT
Content-Length: 0
```

I then created a second user Jane Doe and a second group, user.

PUT to /users/jsmith
```bash
curl -X PUT -H "Content-Type: application/json" -H "Accept:
application/json" -i -d '{"first_name":"John", "last_name":"Smith",
"userid":"jsmith", "groups":["user"]}'
http://localhost:8090/users/jsmith
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:09:14 GMT
Content-Length: 78

{"first_name":"John","last_name":"Smith","userid":"jsmith","groups":["user"]}
```

PUT to /users/notexist (doesn't match JSON object)
```bash
curl -X PUT -H "Content-Type: application/json" -H "Accept:
application/json" -i -d '{"first_name":"John", "last_name":"Smith",
"userid":"jsmith", "groups":["user"]}'
http://localhost:8090/users/notexist
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:09:54 GMT
Content-Length: 3
```

PUT to /users/notexist (matches JSON object)
```bash
curl -X PUT -H "Content-Type: application/json" -H "Accept:
application/json" -i -d '{"first_name":"John", "last_name":"Smith",
"userid":"notexist", "groups":["user"]}'
http://localhost:8090/users/notexist
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:10:40 GMT
Content-Length: 0
```

PUT to /groups/admin (adding jsmith back to admin)
```bash
curl -X PUT -H "Content-Type: application/json" -H "Accept:
application/json" -i -d '{"name":"admin", "members":["jsmith"]}'
http://localhost:8090/groups/admin
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:11:58 GMT
Content-Length: 38

{"name":"admin","members":["jsmith"]}

curl -X GET -H "Accept: application/json" -i
http://localhost:8090/users/jsmith
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:12:20 GMT
Content-Length: 86

{"first_name":"John","last_name":"Smith","userid":"jsmith","groups":["user","admin"]}
```

DELETE to /groups/user
```bash
curl -X DELETE -i http://localhost:8090/groups/user
HTTP/1.1 204 No Content
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:13:19 GMT

curl -X GET -H "Accept: application/json" -i
http://localhost:8090/users/jsmith
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:13:34 GMT
Content-Length: 79

{"first_name":"John","last_name":"Smith","userid":"jsmith","groups":["admin"]}
```

DELETE to /user/jsmith
```bash
curl -X DELETE -i http://localhost:8090/users/jsmith
HTTP/1.1 204 No Content
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:14:28 GMT

curl -X GET -H "Accept: application/json" -i
http://localhost:8090/groups/admin
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Fri, 04 Dec 2015 04:14:52 GMT
Content-Length: 32

{"name":"admin","members":null}
```

That's POST, PUT, GET, and DELETE covered for both groups and users and
it appears to meet the specified requirements while maintain
consistency for group membership between user and group objects.


## Implementation Notes

While the assignment clearly stated that I shouldn't get hung up on
persistence issues, I did a little bit.  My mock persistence layer is a
bit more complex than I wanted, but it achieves all of the basic
requirements to make the API sane. I could probably have implemented it
better by using a map instead of a slice, which would have eliminated
some nested loops, but I was trying to keep the implementation simple to
understand/debug vs trying to be elegant, considering how hackish it
already was.

As I had discussed on the phone with the recruiter, I'm not really much of a software developer.
I am able to read/write code, and I understand concepts, but I'm much more focused on
operations/infrastructure type problems typically.  That said, while what I write works,
it's probably not conforming to style best practices or idiomatic Go.  If I were doing
this in a team setting on a regular basis, I would either learn expected style
as I go from code reviews or expect some sort of style guidelines that the team follows.

## Contact Information

Cell: 210-213-7249  
Email: tristor@gmail.com
