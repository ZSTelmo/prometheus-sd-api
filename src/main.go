package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Prometheus SD API Server")

	ReadFile()
	// handle route using handler function
	http.HandleFunc("/targets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint targets: returnAllTargets")
		json.NewEncoder(w).Encode(ReadFile())
	})

	http.ListenAndServe(":5000", nil)
}
