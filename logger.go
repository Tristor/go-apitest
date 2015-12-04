package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logfile *os.File
var logger *log.Logger

func init() {
	logfile, err := os.OpenFile("go-apitest.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger = log.New(logfile, "apitest:", log.Lshortfile|log.LstdFlags)
}

//Logger takes in an http.Handler and a name and returns a crafted
//http.Handler.  This provides us a simple wrapper that can be used
//anywhere an http.Handler is expected to add simple request logging.
//This is currently only used when generating our mux.Router.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
