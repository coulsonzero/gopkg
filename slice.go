package gopkg

import "fmt"

// make: 主动扩容, 返回新的slice, 原来的slice不变
// copy: 拷贝短的slice元素到另一个长的slice中, 自动截取
// append: 自动扩容
// var arr = [...]int{1, 2, 4: 5, 6}          // len:  6, cap:  6, array: [1 2 0 0 5 6]
// var arr = [...]int{'a': 1}                 // len: 98, cap: 98, array: [..., 1]

type sl interface {
	int | int64 | float64 | string | bool
}

func SlicePrint[T sl](slice []T) {
	fmt.Printf("len: %d, cap: %d, slice: %v \n", len(slice), cap(slice), slice)
}

func SliceInsert[T sl](slice []T, index int, value T) []T {
	return append(slice[:index], append([]T{value}, slice[index:]...)...)
}

func sliceInsert(slice []int, index int, value int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func SliceRemove[T sl](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func SliceReverse[T sl](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SliceContains[T sl](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}