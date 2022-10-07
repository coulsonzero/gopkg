package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
)

var arr = []string{"h", "e", "l", "l", "o"}

func main() {
	// fmt.Println(DemoStringBuild(arr))  // hello
	// fmt.Println(DemoStringBuffer(arr)) // hello
	fmt.Println(gopkg.ToTitle("hello, world")) // Hello, World
}

func DemoStringBuild(arr []string) string {
	return gopkg.StringBuild(arr)
}

func DemoStringBuffer(arr []string) string {
	return gopkg.StringBuffer(arr)
}
