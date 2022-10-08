package main

import (
	"fmt"
)

func Sum() int {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum
}
