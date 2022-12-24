package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/runtime/unsafe/pro"
)

type student struct {
	name string `json:"name" level:"12"`
	age  int    `json:"age"`
}

func main() {
	s := student{}
	res := pro.GetStructTag(s, "name", "json", "level")
	fmt.Println(res)
}
