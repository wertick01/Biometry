package biostatics

import (
	"biometry/bioreader"
	"github.com/montanaflynn/stats"
  "math"
  "fmt"
)

type TwoTwelw struct {
	Height VarMeanMedian
	Weight VarMeanMedian
}

type VarMeanMedian struct {
	variability float64
	mean float64
	median float64
  q1 float64
  q3 float64
}
/*
func Initer(path string) {
	res := bioreader.Constructor(ReadCsvFile(path))
	from2to12(res)
}
*/
func From2to12(res *bioreader.Table) *TwoTwelw {
  var ms1, ms2 []float64
  for _, val := range res.Height {
    if val != 0.0 {
      ms1 = append(ms1, val)
    }
  }
  mean1, _ := stats.Mean(ms1)
	median1, _ := stats.Median(ms1)
	variance1, _ := stats.VarS(ms1)
	quartiles1, _ := stats.Quartile(ms1)

  for _, val := range res.Weight {
    if val != 0.0 {
      ms2 = append(ms2, val)
    }
  }
  mean2, _ := stats.Mean(ms2)
	median2, _ := stats.Median(ms2)
	variance2, _ := stats.VarS(ms2)
  quartiles2, _ := stats.Quartile(ms2)
  fmt.Println(len(res.Height), len(res.Weight))
  
  return &TwoTwelw{
    Height: VarMeanMedian{
      variability: variance1,
      mean: mean1,
      median: median1,
      q1: quartiles1.Q1,
      q3: quartiles1.Q3,
      },
    Weight: VarMeanMedian{
      variability: variance2,
      mean: mean2,
      median: median2,
      q1: quartiles2.Q1,
      q3: quartiles2.Q3,
    },
  }
}

func From13to18(res *bioreader.Table) {
  var ms1, ms2 []float64
  var ms3, ms4 []float64
  for _, val := range res.Height {
    if val != 0 {
      ms1 = append(ms1, val)
    }
  }
	quartiles1, _ := stats.Quartile(ms1)
  hiqr, _ := stats.InterQuartileRange(ms1)
  lheight25, uheight75 := quartiles1.Q1 - hiqr, quartiles1.Q3 + hiqr
  fmt.Println(lheight25, uheight75)
  for _, val := range res.Height {
    if val > uheight75 {
      ms3 = append(ms3, 0.0)
    } else if val < lheight25 {
      ms3 = append(ms3, 0.0)
    } else {
      ms3 = append(ms3, val)
    }
  }

  for _, val := range res.Weight {
    if val != 0 {
      ms2 = append(ms2, val)
    }
  }
  quartiles2, _ := stats.Quartile(ms2)
  wiqr, _ := stats.InterQuartileRange(ms2)
  lweight25, uweight75 := quartiles2.Q1 - wiqr, quartiles2.Q3 + wiqr
  fmt.Println(lweight25, uweight75)
    for _, val := range res.Weight {
    if val > uweight75 {
      ms4 = append(ms4, 0.0)
    } else if val < lweight25 {
      ms4 = append(ms4, 0.0)
    } else {
      ms4 = append(ms4, val)
    }
  }
  var bm1 []float64
  bmi := BMI(ms3, ms4)
  for _, val := range bmi {
    if val != 0 {
      bm1 = append(bm1, val)
    }
  }
  mean, median, variance := bmistat(bm1)
  fmt.Println(mean, median, variance)
}

func BMI(height, weight []float64) []float64 {
  var res []float64
  for i := 0; i < len(height); i ++ {
    if height[i] != 0 {
      res = append(res, weight[i]/math.Pow(height[i] * 0.01, 2))
    } else {
      res = append(res, 0.0)
    }
  }
  return res
}

func bmistat(mass []float64) (float64, float64, float64) {
  mean, _ := stats.Mean(mass)
	median, _ := stats.Median(mass)
	variance, _ := stats.VarS(mass)
  return mean, median, variance
}