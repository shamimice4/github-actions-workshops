package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define the function that handles incoming requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Everyone! Have a nice Day")
	}

	// Register the handler function to handle all requests to the root URL path
	http.HandleFunc("/", handler)

	// Start the HTTP server and listen for incoming requests on port 8080
	fmt.Println("HTTP Server Starting")
	http.ListenAndServe(":8080", nil)
}
