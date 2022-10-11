package pro

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	_ "unsafe" // for go:linkname
)

//go:linkname toTitle github.com/coulsonzero/gopkg/pro.ToTitle
// toTitle 每个单词首字母大写
// example: "hello, world!" -> "Hello, World!"
func toTitle(s string) string {
	return cases.Title(language.Und).String(s)
}
