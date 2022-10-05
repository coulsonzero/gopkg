package test

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
)

type Student struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func TestEncode() string {
	s := &Student{
		Name:  "john",
		Email: "john@gmail.com",
		Age:   20,
	}
	res, err := gopkg.Encode(s)
	if err != nil {
		fmt.Println("error: failed to encode obj to json")
	}
	return res
	// {"name":"john","email":"john@gmail.com","age":20}
}

func TestDecode() {
	// res := `{"Name":"john","Email":"john@gmail.com","Age":20}`
	res := TestEncode()
	ret := gopkg.Decode([]byte(res), Student{})
	fmt.Println(ret)
	// &{john john@gmail.com 20}
}
