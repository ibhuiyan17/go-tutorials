package main

import (
	"fmt"
	"net/http"
)

/*
	This is a handler. Specific signature needed to implement
	the http.Handler interface that works with calling it in main
*/
func hello(w http.ResponseWriter, req *http.Request) { // handler returns "hello"
	fmt.Println("received a request")

	fmt.Fprint(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) { // handler returns req headers
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	/* Registers the /hello endpoint that is imlemented by
	the hello function handler */
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":5050", nil) // listen on port 5050 on default serveMux
}
