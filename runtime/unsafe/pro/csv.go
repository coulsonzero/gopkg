package pro

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	_ "unsafe" // for go:linkname
)

/**
 *  res := readCsv("./src/conf/test.csv")
 *	printCsv(res)
 *	writeCsv("./src/conf/res.csv", res)
 */

//go:linkname readCsv github.com/coulsonzero/gopkg/pro/csv.ReadCsv
func readCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error: open file")
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("error: read csv file recodes")
	}

	return records
}

//go:linkname printCsv github.com/coulsonzero/gopkg/pro/csv.PrintCsv
func printCsv(records [][]string) {
	for _, record := range records {
		fmt.Printf("%v\n", record)
	}
}

//go:linkname writeCsv github.com/coulsonzero/gopkg/pro/csv.WriteCsv
func writeCsv(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("error: create file")
	}
	w := csv.NewWriter(file)
	err = w.WriteAll(data)
	if err != nil {
		log.Fatal("error: write csv file")
	}
	w.Flush()
}
