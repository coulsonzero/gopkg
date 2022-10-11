package pro

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	_ "unsafe" // for go:linkname
)

//go:linkname readFileSql github.com/coulsonzero/gopkg/pro.ReadSql
// readFileSql read sql string from sql file
func readFileSql(path string) (string, error) {
	sqlFile, err := filepath.Abs(path)
	if err != nil {
		return "", errors.New("error: file path error ")
	}

	file, err := ioutil.ReadFile(sqlFile)

	if err != nil {
		return "", errors.New("error: read sql file error ")
	}
	return string(file), nil
}
