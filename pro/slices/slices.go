package slices

type sl interface {
	int | int64 | float64 | string | bool
}

func Contains[T sl](array []T, val T) bool

func Index[T sl](array []T) int

func Print[T sl](array ...[]T)

func Insert[T sl](array []T, index int, value T) []T

func Delete[T sl](array []T, index int) []T

func Reverse[T sl](array []T) []T

func IsEqual(x any, y any) bool

func SortStr(array []string, reverse ...bool)

// SortInt to sort an integer Array in place
// example sort   : SortInt(nums) or SortInt(nums, false)
// example reverse: SortInt(nums, true)
func SortInt(nums []int, reverse ...bool)
