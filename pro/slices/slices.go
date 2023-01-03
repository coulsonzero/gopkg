package slices

type S interface {
	int | int64 | float64 | string | bool
}

func Contains[T S](array []T, v T) bool

func Index[T S](array []T, v T) int

func Print[T S](array ...[]T)

func Insert[T S](array []T, i int, v ...T) []T

func Delete[T S](array []T, index int) []T

func Reverse[T S](array []T) []T

func Equal(x any, y any) bool

func SortStr(array []string, reverse ...bool)

// SortInt to sort an integer Array in place
// example sort   : SortInt(nums) or SortInt(nums, false)
// example reverse: SortInt(nums, true)
func SortInt(nums []int, reverse ...bool)
