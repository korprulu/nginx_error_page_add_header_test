package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/auth_success" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.URL.Query().Has("add_header") {
		w.Header().Add("X-Custom-Header-Internal", "This is internal header")
	}

	w.WriteHeader(http.StatusForbidden)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
