package main

import (
	"fmt"
	"Byometry/bioreader"
	"os"
	"encoding/csv"
	"log"
	"github.com/montanaflynn/stats"
)

type TwoTwelw struct {
	Height VarMeanMedian
	Weight VarMeanMedian
}

type VarMeanMedian struct {
	variability float64
	mean float64
	median float64
}

func main() {
	path := "/home/mrred/Загрузки/Rjabova_P_M.csv"
	res := bioreader.Constructor(ReadCsvFile(path))
	from2to12(res)
}
/*
func Initer(path string) {
	res := bioreader.Constructor(ReadCsvFile(path))
	from2to12(res)
}
*/
func from2to12(res *bioreader.Table) {
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