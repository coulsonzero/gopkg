package csv

/*
import "github.com/coulsonzero/gopkg/pro/csv"

func main() {
	res := csv.ReadCsv("./conf/test.csv")
	csv.PrintCsv(res)
	csv.WriteCsv("./conf/res.csv", res)
}
*/

// ReadCsv 读取csv文件
func ReadCsv(filename string) [][]string

// WriteCsv 写入csv文件
func WriteCsv(filename string, data [][]string)

// PrintCsv 打印读取的csv表结果
func PrintCsv(records [][]string)
