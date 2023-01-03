package pro

import (
	"reflect"
	"unsafe"
	_ "unsafe" // for go:linkname
)

//go:linkname isEqual github.com/coulsonzero/gopkg/pro/slices.Equal
//go:linkname isEqual github.com/coulsonzero/gopkg/pro.IsEqual
// isEqual
// 判断两个对象是否相等
func isEqual(x any, y any) bool {
	return reflect.DeepEqual(x, y)
}

// TypeOf
// return 数据类型
func TypeOf(e any) any {
	return reflect.TypeOf(e)
}

// SizeOf
// return 字节大小
func SizeOf(e any) any {
	return unsafe.Sizeof(e)
}

// GetStructTag
// return struct tag info
func GetStructTag(obj interface{}, field string, tags ...string) (res []string) {
	t := reflect.TypeOf(obj)
	if field, ok := t.FieldByName(field); ok {
		for _, v := range tags {
			res = append(res, field.Tag.Get(v))
		}
	}
	return res
}
