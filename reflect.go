package gopkg

import "reflect"

type any = interface{}

func Equal(x any, y any) bool {
	return reflect.DeepEqual(x, y)
}
