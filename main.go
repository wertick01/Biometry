package main

import (
	"biometry/bioreader"
	"biometry/biostatics"
	"fmt"
)

func main() {
	//records := ReadCsvFile("datas/Rjabova_P_M.csv")
	res := bioreader.Constructor(bioreader.ReadCsvFile("datas/Rjabova_P_M.csv"))
  tb := biostatics.From2to12(res)
  fmt.Println(tb)
  biostatics.From13to18(res)
}

