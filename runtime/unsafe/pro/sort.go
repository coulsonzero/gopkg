package pro

import (
	"sort"
	_ "unsafe" // for go:linkname

	"github.com/zhangyunhao116/pdqsort"
)

type sortSlice interface {
	sortStr()
	sortInt()
}

//go:linkname sortStr github.com/coulsonzero/gopkg/pro/slices.SortStr
//go:linkname sortStr github.com/coulsonzero/gopkg/pro.SortStr
// sortStr 字符串数组排序
func sortStr(s []string, reverse ...bool) {
	if len(reverse) > 1 {
		logger("error: reverse param max len only is 1 !")
		return
	}

	if len(reverse) == 0 || !reverse[0] {
		sort.Strings(s)
	} else {
		// 9 34 30 3
		sort.Sort(sort.Reverse(sort.StringSlice(s)))
		// 9 34 3 30
		// sort.Slice(s, func(i, j int) bool {
		//     return s[i]+s[j] > s[j]+s[i]
		// })
	}
}

//go:linkname sortInt github.com/coulsonzero/gopkg/pro/slices.SortInt
//go:linkname sortInt github.com/coulsonzero/gopkg/pro.SortInt
// sortInt 整数数组排序
// To sort an integer Array in place
// example sort: SortInt(nums) or SortInt(nums, false)
// example reverse: SortInt(nums, true)
func sortInt(nums []int, reverse ...bool) {
	if len(reverse) > 1 {
		logger("error: reverse param max len only is 1 !")
		return
	}

	if len(reverse) == 0 || !reverse[0] {
		pdqsort.Slice(nums)
		// sort.Ints(nums)
		// sort.Sort(sort.IntSlice(nums))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	}
}
