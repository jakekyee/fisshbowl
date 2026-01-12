package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// w.Header().Set("Access-Control-Allow-Origin", "https://jakeyee.com")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Allow everything! :)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass to next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	DB = InitDB()
	defer DB.Close()

	// Create ServeMux
	mux := http.NewServeMux()

	// Create a Huma API using the humago adapter
	api := humago.New(mux, huma.DefaultConfig("SSH Logger API", "1.0.0"))

	// Register routes
	RegisterRoutes(api)

	// Wrap the mux with CORS middleware
	handler := corsMiddleware(mux)

	fmt.Println("API running on :9992")
	log.Fatal(http.ListenAndServe(":9992", handler))
}
