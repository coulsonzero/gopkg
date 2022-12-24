package pro

import (
	"fmt"
	_ "unsafe" // for go:linkname
)

// make: 主动分配容量
// copy: 拷贝后者的slice元素到前者的slice中, 自动截断，长度和容量不变
// append: 末尾新增一个或多个元素, 长度超过容量时自动扩容

// var arr = [...]int{1, 2, 4: 5, 6}          // len:  6, cap:  6, array: [1 2 0 0 5 6]
// var arr = [...]int{'a': 1}                 // len: 98, cap: 98, array: [..., 1]

type sl interface {
	int | int64 | float64 | string | bool
}

//go:linkname slice_print github.com/coulsonzero/gopkg/pro.SlicePrint
func slice_print[T sl](slice ...[]T) {
	// fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
	for _, v := range slice {
		fmt.Printf("len: %d, cap: %d, slice: %v \n", len(v), cap(v), v)
	}
}

//go:linkname slice_insert github.com/coulsonzero/gopkg/pro.SliceInsert
// 插入元素
func slice_insert[T sl](slice []T, index int, value T) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

func slice_insert2(slice []int, index int, value int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

//go:linkname slice_remove github.com/coulsonzero/gopkg/pro.SliceRemove
// 删除单个元素
func slice_remove[T sl](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

//go:linkname slice_contains github.com/coulsonzero/gopkg/pro.SliceContains
// 判断是否包含目标元素
func slice_contains[T sl](array []T, val T) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}
	return false
}

//go:linkname slice_reverse github.com/coulsonzero/gopkg/pro.SliceReverse
// 反转
func slice_reverse[T sl](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
