package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Scapper ...
type Scapper struct {
	JobName string `json:"job_name"`
	StaticConfigs StaticConfigs `json:"static_configs"`
	ScrapeInterval string `json:"scrape_interval,omitempty"`
	ScrapeTimeout string `json:"scrape_timeout,omitempty"`
}

// StaticConfigs ...
type StaticConfigs struct {
	Targets []string `json:"targets"`
	Labels Labels `json:"labels,omitempty"`
}

// Labels ...
type Labels struct {
	Alias string `json:"alias,omitempty"`
}

func main() {
	fmt.Println("Prometheus SD API Server")

	ReadFile()
	// handle route using handler function
	http.HandleFunc("/targets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint targets: returnAllTargets")
		json.NewEncoder(w).Encode(ReadFile())
	})

	http.HandleFunc("/register-targets", registerTargets)

	http.ListenAndServe(":5000", nil)
}

func registerTargets(w http.ResponseWriter, r *http.Request) {
	var requestBody Scapper

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var jobs []Scapper
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jobs = append(jobs, requestBody)

	file.Seek(0, 0)
	file.Truncate(0)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}