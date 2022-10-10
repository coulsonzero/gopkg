package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
)

var arr = []string{"h", "e", "l", "l", "o"}

func main() {
	// fmt.Println(DemoStringBuild(arr))  // hello
	// fmt.Println(DemoStringBuffer(arr)) // hello
	fmt.Println(pro.ToTitle("hello, world")) // Hello, World
}
