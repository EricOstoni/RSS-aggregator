// helper for json

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error handler
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	// JSON Serialization (convert data to json )
	dat, err := json.Marshal(payload)

	// Error handling
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	//Setting Response Headers (info the client that the response will be in JSON)
	w.Header().Add("Content-Type", "application/json")
	//Set status code on resonse 200 OK
	w.WriteHeader(code)
	// Write JSON data as the response
	w.Write(dat)

}
