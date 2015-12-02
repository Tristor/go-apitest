/*
This is just an implementation of a simple JSON based REST API for
a programming interview assignment for Planet Labs.  Details about
the assignment, including an API specification can be found in the
provided code_test.md which is in this repository. Further general
documentation can be found in README.md in this repository.
*/
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logfile, err := os.Create("go-apitest.log")
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	//Send all logging to the logfile.  This should capture logs
	//from logger.go too.
	log.SetOutput(logfile)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8090", router))
}
