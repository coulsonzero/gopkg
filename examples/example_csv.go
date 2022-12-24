package main

import "github.com/coulsonzero/gopkg/pro/csv"

func main() {
	res := csv.ReadCsv("./conf/test.csv")
	csv.PrintCsv(res)
	csv.WriteCsv("./conf/res.csv", res)
}
