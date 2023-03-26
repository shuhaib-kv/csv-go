package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("annual-enterprise-survey-2021-financial-year-provisional-csv.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data []map[string]string
	headers := records[0]
	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, value := range record {
			if i < len(headers) {
				row[headers[i]] = value
			}
		}
		data = append(data, row)
	}

	for i, row := range data {
		if i == 0 {
			for key := range row {
				fmt.Printf("%-20s", key)
			}
			fmt.Println()
		}
		for _, value := range row {
			fmt.Printf("%-20s", value)
		}
		fmt.Println()
	}
}
