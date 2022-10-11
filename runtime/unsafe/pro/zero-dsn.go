// version 2:
// need update go version to 1.18
// Author: CoulsonZero
// Date: 2022-10-09

package pro

import (
	"bytes"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)

type mysqlConfig struct {
	user   string
	passwd string
	host   string
	port   string
	dbname string
}

// default config
var cfg = mysqlConfig{
	user:   "root",
	passwd: "root",
	host:   "127.0.0.1",
	port:   "3306",
	dbname: "", // required
}

const params = "?charset=utf8mb4&parseTime=True&loc=Local"

func contains(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), substr)
}

func loadIni(filepath string, mysqlConfig ...string) {
	file, err := ini.Load(filepath)
	if err != nil {
		logger("error: Failed to load ini file")
	}

	for _, v := range mysqlConfig {
		if contains(v, "user") {
			cfg.user = file.Section("mysql").Key(v).String()
		} else if contains(v, "pass") {
			cfg.passwd = file.Section("mysql").Key(v).String()
		} else if contains(v, "host") {
			cfg.host = file.Section("mysql").Key(v).String()
		} else if contains(v, "port") {
			cfg.port = file.Section("mysql").Key(v).String()
		} else if contains(v, "name") && !contains(v, "user") {
			cfg.dbname = file.Section("mysql").Key(v).String()
		}
	}
}

func loadYml(filepath string, mysqlConfig ...string) {
	viper.SetConfigType("yml")
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		logger("error: Failed to load yml file")
	}

	for _, v := range mysqlConfig {
		switch {
		case contains(v, "user"):
			cfg.user = viper.GetString(v)
		case contains(v, "pass"):
			cfg.passwd = viper.GetString(v)
		case contains(v, "host"):
			cfg.host = viper.GetString(v)
		case contains(v, "port"):
			cfg.port = viper.GetString(v)
		case contains(v, "name") && !contains(v, "user"):
			cfg.dbname = viper.GetString(v)
		}
	}
}

func loadEnv(filepath string, mysqlConfig ...string) {
	if err := godotenv.Load(filepath); err != nil {
		logger("error: Failed to load env file")
	}

	for _, v := range mysqlConfig {
		switch {
		case contains(v, "user"):
			cfg.user = os.Getenv(v)
		case contains(v, "pass"):
			cfg.passwd = os.Getenv(v)
		case contains(v, "host"):
			cfg.host = os.Getenv(v)
		case contains(v, "port"):
			cfg.port = os.Getenv(v)
		case contains(v, "name") && !contains(v, "user"):
			cfg.dbname = os.Getenv(v)
		}
	}
}

func (cfg *mysqlConfig) formatDSN() string {
	var buf bytes.Buffer
	if len(cfg.user) > 0 {
		buf.WriteString(cfg.user)
		if len(cfg.passwd) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.passwd)
		}
		buf.WriteString("@tcp(")
	}
	if len(cfg.host) > 0 {
		buf.WriteByte('(')
		buf.WriteString(cfg.host)
		if len(cfg.port) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.port)
		}
		buf.WriteByte(')')
		buf.WriteByte('/')
		buf.WriteString(cfg.dbname)
		buf.WriteString(params)
	}
	return buf.String()
}

func (cfg *mysqlConfig) isValid() {
	switch {
	case len(cfg.user) == 0:
		logger("error: mysql user length cannot be zero")
	case len(cfg.passwd) == 0:
		logger("error: mysql password length cannot be zero")
	case len(cfg.host) == 0:
		logger("error: mysql host length cannot be zero")
	case len(cfg.port) == 0:
		logger("error: mysql port length cannot be zero")
	case len(cfg.dbname) == 0:
		logger("error: mysql dbname length cannot be zero")
	default:
		logger("All mysql config is valid")
	}
}
