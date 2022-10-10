package main

import (
	"github.com/coulsonzero/gopkg/arrays"
)

func main() {
	slice := []int{2, 5, 1, 3}
	slice2 := []float64{2.1, 5.1, 1.3, 3.6}

	arrays.SlicePrint(slice)

	slice = arrays.SliceInsert(slice, 2, 9) // [2 5 9 1 3]
	arrays.SlicePrint(slice)

	slice = arrays.SliceRemove(slice, 2) // [2 5 1 3]
	arrays.SlicePrint(slice)

	arrays.SlicePrint(slice2)

	sliceReverse := arrays.SliceReverse(slice)
	arrays.SlicePrint(sliceReverse)

}
