package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
)

func main() {
	intArr := []int{2, 5, 1, 4, 3, 6}
	// pro.SortInt(intArr)
	pro.SortInt(intArr, true)
	// pro.SortInt(intArr, false)
	fmt.Println(intArr)
	// Sum(2)
	// Sum(2, true)
	// Sum(2, false)
	// Sum(2, false, true)
}
