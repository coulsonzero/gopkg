package main

import (
	"github.com/coulsonzero/gopkg"
)

func main() {
	slice := []int{2, 5, 1, 3}
	slice2 := []float64{2.1, 5.1, 1.3, 3.6}

	gopkg.SlicePrint(slice)

	slice = gopkg.SliceInsert(slice, 2, 9) // [2 5 9 1 3]
	gopkg.SlicePrint(slice)

	slice = gopkg.SliceRemove(slice, 2) // [2 5 1 3]
	gopkg.SlicePrint(slice)

	gopkg.SlicePrint(slice2)

	sliceReverse := gopkg.SliceReverse(slice)
	gopkg.SlicePrint(sliceReverse)

}
