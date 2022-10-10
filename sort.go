package gopkg

import (
	"github.com/zhangyunhao116/pdqsort"
	"sort"
)

func SortString(s []string) {
	sort.Strings(s)
}

// SortStringReverse
// 9 34 30 3
func SortStringReverse(s []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
}

// SortStringReverse2
// 9 34 3 30
func SortStringReverse2(s []string) {
	sort.Slice(s, func(i, j int) bool {
		return s[i]+s[j] >= s[j]+s[i]
	})
}

func SortInt(nums []int) {
	pdqsort.Slice(nums)
	// sort.Ints(nums)
	// sort.Sort(sort.IntSlice(nums))
}

func SortIntReverse(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
}
