package files

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// IsExists 文件/目录是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(dirpath string) bool {
	// if fi, err := os.Stat(dirpath); err != nil || !fi.IsDir() {
	// 	panic("error:  cannot find the directory")
	// 	return false
	// }
	// return true
	fi, err := os.Stat(dirpath)
	return err != nil || !fi.IsDir()

}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// WalkRemoveFiles 移除置顶类型文件
var cnt int

func WalkRemoveFiles(path string, filetype string) {
	dir, _ := os.ReadDir(path)
	for _, file := range dir {
		// 拼接文件路径
		file_path := filepath.Join(path, file.Name())
		// 如果是目录递归，否则继续
		if file.IsDir() {
			WalkRemoveFiles(file_path, filetype)
		} else {
			// 如果文件以...为后缀
			if strings.HasSuffix(file_path, filetype) {
				cnt++
				fmt.Printf("%-4d: %s\n", cnt, file_path)
				if err := os.Remove(file_path); err != nil {
					return
				}
			}
		}
	}
}

func ReadFile(path string) (string, error) {
	if IsFile(path) {
		d, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(d), nil
	}
	return "", errors.New(path + "is not exists")
}
