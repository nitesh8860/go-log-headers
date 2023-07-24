package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func logRequestHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request headers in JSON format
		headers := make(map[string][]string)
		for key, values := range r.Header {
			headers[key] = values
		}

		data, err := json.MarshalIndent(headers, "", "  ")
		if err != nil {
			log.Printf("Error marshaling headers to JSON: %s", err)
		} else {
			log.Printf("Request Headers:\n%s\n", string(data))
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Set the log output to stdout
	log.SetOutput(os.Stdout)

	// Create a new server mux and attach the logRequestHeaders middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is the Go server!"))
	})

	// Wrap the main handler with the logRequestHeaders middleware
	handler := logRequestHeaders(mux)

	// Start the server on port 8082
	port := ":8082"
	log.Printf("Starting server on port %s\n", port)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
