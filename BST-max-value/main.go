package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Req struct {
	Data [][]int `json:"data"`
}

// O(n^2)
func findMaxPathSum(data [][]int) int {
	// if no data just return 0
	if len(data) == 0 {
		return 0
	}
	// start from second row
	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			data[i][j] += max(data[i+1][j], data[i+1][j+1])
		}
	}

	return data[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDataFormJson() (Req, error) {
	var jsonData Req
	file, err := os.Open("hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return jsonData, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return jsonData, err
	}
	return jsonData, nil
}

func main() {
	req, err := getDataFormJson()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Maximum path: %v", findMaxPathSum(req.Data))
}
