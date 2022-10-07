package main

import (
	"fmt"
	"github.com/coulsonzero/gopkg"
	"log"
)

var (
	envArr = []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	iniArr = []string{"Db_username", "Db_password", "Db_host", "Db_port", "Db_name"}
	ymlArr = []string{"mysql.username", "mysql.password", "mysql.host", "mysql.port", "mysql.dbname"}
	// DSN    = "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	DemoConfigYml()
	DemoConfigInt()
	DemoConfigEnv()
}

func DemoConfigYml() {
	// 初始化yml配置
	dsn, err := gopkg.ConfigYml("conf/config.yml", ymlArr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true, // would be `user`
	// 	},
	// })
	//
	// if err != nil {
	// 	log.Fatal("Error: Failed to create a connection to database")
	// }
	//
	// log.Fatalf("Success connect to mysql. %v", db)
}

func DemoConfigInt() {
	dsn, err := gopkg.ConfigInt("conf/config.ini", iniArr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dsn)
}

func DemoConfigEnv() {
	dsn, err := gopkg.ConfigEnv("conf/qqw.env", envArr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dsn)
}
