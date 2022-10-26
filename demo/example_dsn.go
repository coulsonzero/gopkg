package main

import (
	"database/sql"
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	dsn, _ := pro.ConfDSN("conf/config.yml", "mysql.dbname")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error open mysql connect")
	}
	fmt.Println("Success to connect mysql")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Fatal("ping 失败")
	} else {
		fmt.Println("pong")
	}

	defer db.Close()
}
