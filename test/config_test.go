package main

import (
	"github.com/coulsonzero/gopkg"
	"testing"
)

func TestConfigEnv(t *testing.T) {
	testEnvArr := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	// testEnvArr := []string{"DB_USER", "DB_PASSWORD", "DB_NAME"}
	// testEnvArr := []string{"DB_NAME"}

	dsn, err := gopkg.ConfigEnv("./conf/.env", testEnvArr)
	if err != nil {
		return
	}

	if dsn != "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local" {
		t.Errorf("error: not equal: %q", dsn)
	}
}
