package testing

import (
	"database/sql"
	"fmt"
	"github.com/coulsonzero/gopkg/pro"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestMysqlConn(t *testing.T) {
	dsn, _ := pro.ConfDSN("conf/config.yml", "mysql.dbname")
	// dsn := "root:root@tcp(localhost:3306)/go_study?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Errorf("Error open mysql connect: %q", dsn)
	}
	fmt.Println("Success to connect mysql")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		t.Error("ping 失败")
	}
	// fmt.Println("pong")

	defer db.Close()
}
