package main

import (
  "fmt"
  "github.com/tealeg/xlsx"
)

//https://godoc.org/github.com/tealeg/xlsx#example-Row-ReadStruct

func main() {
  //create a new xlsx file and write a struct
  //in a new row
  f := xlsx.NewFile()
  // sheet, _ := f.AddSheet("TestSheet")
  // row := sheet.AddRow()



  err := f.Save("test1.xlsx")

  if err != nil {
    fmt.Println(err)
  }
}
