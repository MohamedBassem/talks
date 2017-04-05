package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		fmt.Printf("Got: %v\n", name)
		fmt.Fprintf(w, "Hello %v\n", name)
	})

	http.ListenAndServe(":8080", nil)
}

// END OMIT
