package pro

import (
	"github.com/zhangyunhao116/pdqsort"
	"sort"
	_ "unsafe" // for go:linkname
)

type sortSlice interface {
	sortStr()
	sortInt()
}

//go:linkname sortStr github.com/coulsonzero/gopkg/pro.SortStr
// sortStr 字符串数组排序
func sortStr(s []string, reverse bool) {
	if reverse {
		// 9 34 30 3
		sort.Sort(sort.Reverse(sort.StringSlice(s)))
		// 9 34 3 30
		// sort.Slice(s, func(i, j int) bool {return s[i]+s[j] >= s[j]+s[i]})
	} else {
		sort.Strings(s)
	}
}

//go:linkname sortInt github.com/coulsonzero/gopkg/pro.SortInt
// sortInt 整数数组排序
func sortInt(nums []int, reverse bool) {
	if reverse {
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	} else {
		pdqsort.Slice(nums)
		// sort.Ints(nums)
		// sort.Sort(sort.IntSlice(nums))
	}
}
