package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	"log"
)

func main() {
	DemoReadSql()
}

func DemoReadSql() {
	sqlStr, err := pro.ReadSql("conf/user.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sqlStr)
}
