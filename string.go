package gopkg

import (
	"bytes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func StringBuild(arr []string) string {
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(v)
	}
	return sb.String()
}

func StringBuffer(arr []string) string {
	buf := new(bytes.Buffer)
	for _, v := range arr {
		buf.Write([]byte(v))
	}
	return buf.String()
}

// ToTitle 每个单词首字母大写
// example: "hello, world!" -> "Hello, World!"
func ToTitle(s string) string {
	return cases.Title(language.Und).String(s)
}
