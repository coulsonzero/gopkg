package gopkg

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

func ReadSql(path string) (string, error) {
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
