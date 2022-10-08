package gopkg

import "reflect"

// type any = interface{}

func IsEqual(x any, y any) bool {
	return reflect.DeepEqual(x, y)
}
