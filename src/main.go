package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Label struct {
	Job      string `json:"job"`
	Targests string `json:"targets"`
}

var Labels []Label

func main() {
	fmt.Println("Prometheus SD API Server")

	Labels = []Label{
		Label{Job: "mysql",
			Targests: "01.mydomain.com:9112"},
		Label{Job: "mysql",
			Targests: "02.mydomain.com:9112"},
		Label{Job: "mysql",
			Targests: "03.mydomain.com:9112"},
	}

	// handle route using handler function
	http.HandleFunc("/targets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint targets: returnAllTargets")
		json.NewEncoder(w).Encode(Labels)
	})

	// listen to port
	http.ListenAndServe(":5000", nil)
}
