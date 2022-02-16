package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

func readCsvFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ','
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records, err
}

func GetDataFromFile(fname string) ([]time.Time, []float64) {
	var dates []time.Time
	var elapsed []float64

	resp, err := readCsvFile(fname)
	if err == nil {
		for i, record := range resp {
			// skip header
			if i == 0 {
				continue
			}
			ts := record[0]
			parsed, err := time.Parse(chart.DefaultDateFormat, ts)
			if err == nil {
				closeP, err := strconv.ParseFloat(record[4], 64)
				if err == nil {
					dates = append(dates, parsed)
					elapsed = append(elapsed, closeP)
				}
			}
		}
		if len(resp)-1 > len(elapsed) {
			log.Println("Number of invalid records =", len(resp)-1-len(elapsed))
		}
	}
	return dates, elapsed
}
