package testing

import (
	"github.com/coulsonzero/gopkg/pro"
	"reflect"
	"testing"
)

type Student struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func TestJson(t *testing.T) {
	s := &Student{
		Name:  "john",
		Email: "john@gmail.com",
		Age:   20,
	}
	// obj -> json
	res, err := pro.Encode(s)
	if err != nil {
		return
	}

	// json -> obj
	s2 := &Student{}
	pro.Decode([]byte(res), s2)

	if !reflect.DeepEqual(s, s2) {
		t.Errorf("s1: %q, s2: %q", *s, *s2)
	}
}
