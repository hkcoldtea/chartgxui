package main

import (
	"encoding/csv"
	"log"
	"os"
)

func writeCsvFile(data [][]string) {
	file, err := os.Create("records.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(data)
	log.Println("records saved")
}
