package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	"log"
)

var testcases = map[string][]string{
	"conf/config.yml": {"mysql.dbname", "mysql.username", "mysql.password", "mysql.host", "mysql.port"},
	"conf/config.ini": {"Db_username", "Db_password", "Db_host", "Db_port", "Db_name"},
	"conf/.env":       {"db_name"},
}

func main() {
	pro.Now()
	fmt.Println(pro.MD5Encode("admin123"))

	TestDSN()

}

func TestDSN() {
	for fp, tc := range testcases {
		orig, err := pro.ConfDSN(fp, tc...)
		if err != nil {
			log.Fatal("error: failed to load confDSN!")
		}
		want := "root:root@tcp((127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
		if orig != want {
			log.Printf("error: \norig: %s; \nwant: %s \n", orig, want)
		}
	}
	log.Println("All test success!")
}
