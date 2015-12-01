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

## Test Procedure


## Implementation Notes

While the assignment clearly stated that I shouldn't get hung up on persistence issues, I did a little bit.  My mock persistence layer is a bit more complex than I wanted, but it achieves all of the basic requirements to make the API sane.

As I had discussed on the phone with the recruiter, I'm not really much of a software developer.  I am able to read/write code, and I understand concepts, but I'm much more focused on operations/infrastructure type problems typically.  That said, while what I write works, it's probably not conforming to style best practices or idiomatic Go.  If I were doing this in a team setting on a regular basis, I would either learn expected style as I go from code reviews or expect some sort of style guidelines that the team follows.

## Contact Information

Cell: 210-213-7249  
Email: tristor@gmail.com
