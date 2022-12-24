package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
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
	res, err := pro.Encode(s)
	if err != nil {
		fmt.Println("error: failed to encode obj to json")
	}
	return res
	// {"name":"john","email":"john@gmail.com","age":20}
}

func TestDecode() {
	// res := `{"Name":"john","Email":"john@gmail.com","Age":20}`
	res := TestEncode()
	ret := pro.Decode([]byte(res), Student{})
	fmt.Println(ret)
	// &{john john@gmail.com 20}
}
