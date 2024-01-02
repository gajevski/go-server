package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		fmt.Println("Server running")
		fmt.Println("Hello, you've requested:", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
