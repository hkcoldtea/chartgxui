// +build !debug

package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	chart "github.com/wcharczuk/go-chart"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func debug(format string, args ...interface{}) {}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

//convert GBK to UTF-8
func DecodeName(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.HZGB2312.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GetData(fname string) ([]time.Time, []float64) {
	var dates []time.Time
	var elapsed []float64

	switch fname {
	case "HSI":
		t2 := time.Now()
		t1 := t2.AddDate(-1, 0, 0)
		url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/download/%%5EHSI?period1=%d&period2=%d&interval=1d&events=history&includeAdjustedClose=true", t1.Unix(), t2.Unix())
		dates, elapsed = GetDataFromURL(url)
	default:
		dates, elapsed = GetDataFromFile(fname)
	}
	return dates, elapsed
}

func GetDataFromURL(url string) ([]time.Time, []float64) {
	var dates []time.Time
	var elapsed []float64

	resp, err := readCSVFromUrl(url)
	if err == nil {
		for i, record := range resp {
			// skip header
			if i == 0 {
				continue
			}
			ts := record[0]
			parsed, _ := time.Parse(chart.DefaultDateFormat, ts)
			dates = append(dates, parsed)
			closeP, _ := strconv.ParseFloat(record[4], 64)
			elapsed = append(elapsed, closeP)
		}
	}
	return dates, elapsed
}
