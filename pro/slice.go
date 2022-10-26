package pro

type sl interface {
	int | int64 | float64 | string | bool
}

func SlicePrint[T sl](slice ...[]T)

func SliceInsert[T sl](slice []T, index int, value T) []T

func SliceRemove[T sl](slice []T, index int) []T

func SliceRemoveElems[T sl](slice []T, index int) []T

func SliceContains[T sl](array []T, val T) bool

func SliceReverse[T sl](s []T) []T

func IsEqual(x any, y any) bool

func SortStr(s []string, reverse ...bool)

// SortInt to sort an integer Array in place
// example sort   : SortInt(nums) or SortInt(nums, false)
// example reverse: SortInt(nums, true)
func SortInt(nums []int, reverse ...bool)
