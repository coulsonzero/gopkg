package test

import (
	"github.com/coulsonzero/gopkg/arrays"
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{2, 5, 1, 3}
	newSlice := []int{2, 5, 9, 1, 3}

	slice = arrays.SliceInsert(slice, 2, 9)
	if !reflect.DeepEqual(slice, newSlice) {
		t.Errorf("error: not equal: slice: %q, newSlice: %q", slice, newSlice)
	}
}
