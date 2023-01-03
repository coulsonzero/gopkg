package main

import "github.com/coulsonzero/gopkg/pro/slices"

func main() {
	nums := []int{1, 2, 3}
	nums2 := []int{1, 2, 3}
	println(slices.Equal(nums, nums2))
	// println(slices.Contains(nums, 1))
	// println(slices.Contains(nums, 2))
}
