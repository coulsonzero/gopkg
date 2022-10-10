package test

import (
	"github.com/coulsonzero/gopkg"
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
	res, err := gopkg.Encode(s)
	if err != nil {
		return
	}

	// json -> obj
	s2 := &Student{}
	gopkg.Decode([]byte(res), s2)

	if !reflect.DeepEqual(s, s2) {
		t.Errorf("s1: %q, s2: %q", *s, *s2)
	}
}
