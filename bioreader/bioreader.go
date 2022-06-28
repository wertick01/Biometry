package bioreader

import (
  "os"
  "encoding/csv"
	//"fmt"
	"log"
	"strconv"
	"time"
	//"reflect"
)

type Table struct {
	IID       []int
	FID       []int
	BFM       []int
	MDD       []int
	LDDD      []int
	CLBP      []int
	Sex       []string
	PA        []string
	Ethnicity []string
	Birthday  []time.Time //!
	Visit     []time.Time //!
	Weight    []float64
	Height    []float64
}

type VMM interface {
  
}

func Constructor(mass [][]string) *Table {
	//var table *Table
	mass = mass[1:]
	IID, FID := FIID(mass)
	Weight, Height := HeightWeight(mass)
	BFM, MDD, LDDD, CLBP := BFMMDD(mass)
	Sex, PA, Ethnicity := SexPAEthn(mass)
	Birthday, Visit := Dater(mass)
	tb := &Table{
		IID:       IID,
		FID:       FID,
		Height:    Height,
		Weight:    Weight,
		BFM:       BFM,
		MDD:       MDD,
		LDDD:      LDDD,
		CLBP:      CLBP,
		Sex:       Sex,
		PA:        PA,
		Ethnicity: Ethnicity,
		Birthday:  Birthday,
		Visit:     Visit,
	}
	return tb
}

func Dater(mass [][]string) ([]time.Time, []time.Time) {
	var Birthday, Visit []time.Time
	var t time.Time
	for _, val := range mass {
		i, err := time.Parse("2006-01-02", val[3])
		if err != nil {
			t, _ = time.Parse("0000-00-00", "0000-00-00")
			Birthday = append(Birthday, t)
		} else {
			Birthday = append(Birthday, i)
		}

		i, err = time.Parse("2006-01-02", val[9])
		if err != nil {
			t, _ = time.Parse("0000-00-00", "0000-00-00")
			Visit = append(Visit, t)
		} else {
			Visit = append(Visit, i)
		}

	}
	return Birthday, Visit
}

func SexPAEthn(mass [][]string) ([]string, []string, []string) {
	var Sex, PA, Ethnicity []string
	for _, val := range mass {
		Sex = append(Sex, val[2])
		PA = append(PA, val[7])
		Ethnicity = append(Ethnicity, val[8])
	}
	return Sex, PA, Ethnicity
}

func BFMMDD(mass [][]string) ([]int, []int, []int, []int) {
	var BFM, MDD, LDDD, CLBP []int
	for _, val := range mass {
		i, err := strconv.Atoi(val[6])
		check(err)
		BFM = append(BFM, i)

		i, err = strconv.Atoi(val[10])
		check(err)
		MDD = append(MDD, i)

		i, err = strconv.Atoi(val[11])
		check(err)
		LDDD = append(LDDD, i)

		i, err = strconv.Atoi(val[12])
		check(err)
		CLBP = append(MDD, i)
	}

	return BFM, MDD, LDDD, CLBP
}

func FIID(mass [][]string) ([]int, []int) {
	var IID, FID []int
	for _, val := range mass {
		i, err := strconv.Atoi(val[0])
		check(err)
		IID = append(IID, i)

		i, err = strconv.Atoi(val[1])
		check(err)
		FID = append(FID, i)
	}
	return IID, FID
}

func HeightWeight(mass [][]string) ([]float64, []float64) {
	var Height, Weight []float64
	for _, val := range mass {
		if val[4] != "" {
			j, err := strconv.ParseFloat(val[4], 64)
			check(err)
			Height = append(Height, j)
		} else {
			Height = append(Height, 0.0)
		}
		if val[5] != "" {
			i, err := strconv.ParseFloat(val[5], 64)
			check(err)
			Weight = append(Weight, i)
		} else {
			Weight = append(Weight, 0.0)
		}
	}
	return Height, Weight
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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
