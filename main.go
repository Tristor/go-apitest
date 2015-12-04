/*
This is just an implementation of a simple JSON based REST API for
a programming interview assignment for Planet Labs.  Details about
the assignment, including an API specification can be found in the
provided code_test.md which is in this repository. Further general
documentation can be found in README.md in this repository.
*/
package main

import (
	"net/http"
)

func main() {
	router := NewRouter()
	logger.Fatal(http.ListenAndServe(":8090", router))
}
