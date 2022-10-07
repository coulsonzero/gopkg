package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
	"log"
)

func main() {
	DemoReadSql()
}

func DemoReadSql() {
	sqlStr, err := gopkg.ReadSql("conf/user.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sqlStr)
}
