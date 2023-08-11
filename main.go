package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Prometheus SD API Server")

	// handle route using handler function
	http.HandleFunc("/targets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new Prometheus SD API Server")
	})

	// listen to port
	http.ListenAndServe(":5000", nil)
}
