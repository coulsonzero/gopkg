package testing

import (
	"github.com/coulsonzero/gopkg/pro"
	"reflect"
	"testing"
)

func TestSortInts(t *testing.T) {
	intArr := []int{2, 5, 1, 4, 3, 6}

	reversedArr := make([]int, 0)
	sortedArr := make([]int, 0)

	// reverse sort
	pro.SortInt(intArr, true)
	reversedArr = append(reversedArr, intArr...)

	// sort
	pro.SortInt(intArr)
	sortedArr = append(sortedArr, intArr...)

	if !reflect.DeepEqual(sortedArr, intArr) {
		t.Errorf("intArr: %v \nreversedArr: %v \nsortedArr: %v \n", intArr, reversedArr, sortedArr)
	}
	// fmt.Printf("intArr: %v \nreversedArr: %v \nsortedArr: %v \n", intArr, reversedArr, sortedArr)
}
