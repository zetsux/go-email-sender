package main

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filePath string) (rows [][]string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err = reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return rows, nil
}
