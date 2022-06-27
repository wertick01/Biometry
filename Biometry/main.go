package main

import (
	"Biometry/bioreader"
	"encoding/csv"
	"fmt"
	"log"
	"os"
    "github.com/montanaflynn/stats"
)

func main() {
	records := ReadCsvFile("/home/mrred/Загрузки/Rjabova_P_M.csv")
	res := bioreader.Constructor(records)
	a, _ := stats.Min(res.Height)
	fmt.Println(a)
}

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
