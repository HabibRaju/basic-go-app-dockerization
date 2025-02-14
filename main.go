package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"status": "ok tres"}`)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}
