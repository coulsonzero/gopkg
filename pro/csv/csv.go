package csv

/**
 *  res := readCsv("test.csv")
 *	printCsv(res)
 *	writeCsv("./src/conf/res.csv", res)
 */

// ReadCsv 读取csv文件
func ReadCsv(filename string) [][]string

// WriteCsv 写入csv文件
func WriteCsv(filename string, data [][]string)

// PrintCsv 打印读取的csv表结果
func PrintCsv(records [][]string)
