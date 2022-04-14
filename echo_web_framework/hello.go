package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Web Server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World\n")
		io.WriteString(w, "This is pretty amazing!\n")
		io.WriteString(w, "I like to write code!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
