package main

import (
	"fmt"
	"log"
	"net/http"  // Go's standard library includes a powerful package net/http for building HTTP clients and servers. currently we are building server
)

// read from  main function

// helloHandler handles HTTP requests sent to the "/hello" endpoint.
// It checks if the request method is GET and responds with "Hello!".
// If the method is not GET, it returns a 404 error.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 'r' is a pointer to the HTTP request received from the client.
	// 'w' is an interface used to construct and send the HTTP response back to the client.

	if r.Method != "GET" {
		// If the request method is not GET, respond with a 404 Not Found and an error message.
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	// Send a simple "Hello!" message as a response if the request method is GET.
	// fmt.Fprintf writes formatted output to the response writer 'w'.
	fmt.Fprintf(w, "Hello !")
}


func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse error  %v ", err)
		return
	}

	fmt.Fprintf(w, "POST request complete \n")
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Fprintf(w, "username %s \n", username)
	fmt.Fprintf(w, "password %s\n", password)
	log.Printf("Received form: username=%s, password=%s", username, password)

}


func main() {
	// Create a file server to serve static files (like HTML, CSS, JS) from the "./static" directory.
	// http.Dir is used to specify the local directory to be served.
	// * this will NOT start server, return http.Handler to routing
	fileserver := http.FileServer(http.Dir("./static"))

	// Map the root URL path "/" to the file server.
	// This means accessing "http://localhost:8080/" will serve the files from "./static".
	http.Handle("/", fileserver)

	// Map the "/hello" route to the helloHandler function.
	// When a request is made to "http://localhost:8080/hello", helloHandler will handle it.
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/form", formHandler)

	// Print a message indicating the server is starting on port 8080.
	fmt.Printf("Starting server that connects to port 8080\n")

	// * Start the HTTP server on port 8080. return Error, Routing not valid here.
	// The second argument is 'nil', meaning the default ServeMux (multiplexer) will be used to route requests.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Log the error and stop the program if the server fails to start.
		log.Fatal(err)
	}
}
