package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// START OMIT
type GreetingRequest struct {
	Name string
}

type GreetingResponse struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Read Body
		var greq GreetingRequest
		err := json.NewDecoder(req.Body).Decode(&greq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		// Send response
		resp := GreetingResponse{
			Message: fmt.Sprintf("Hello %v", greq.Name),
		}
		json.NewEncoder(w).Encode(resp)
	})
	http.ListenAndServe(":8080", nil)
}

// END OMIT
