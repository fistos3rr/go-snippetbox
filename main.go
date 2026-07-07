package main

import (
	"log"
	"net/http"
)

// Function home is a handler which writes a byte slice
// containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// http.NewServeMux() function to initialize a servemux, then
	// register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// http.ListenAndServe() function to start a new web server.
	// Pass two params: TCP network address to listen on
	// and the servemux we created. If it returns error, we log it
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
