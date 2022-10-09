package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
)

type student struct {
	name string `json:"name" level:"12"`
	age  int    `json:"age"`
}

func main() {
	s := student{}
	res := gopkg.GetStructTag(s, "name", "json", "level")
	fmt.Println(res)
}
