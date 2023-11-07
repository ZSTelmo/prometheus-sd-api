package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Label ...
type Label struct {
	Job      string `json:"job"`
	Targets string `json:"targets"`
}

// ReadFile ...
func ReadFile() []Label {
	file, err := os.Open("db.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var labels []Label
	err = json.Unmarshal(data, &labels)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	return labels

}
