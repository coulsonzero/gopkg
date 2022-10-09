package gopkg

import (
	"bytes"
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strings"
)

// version 2

type mysqlConf struct {
	user   string
	passwd string
	host   string
	port   string
	dbname string
}

// default config
var cfg = mysqlConf{
	user:   "root",
	passwd: "root",
	host:   "127.0.0.1",
	port:   "3306",
	dbname: "", // required
}

func ConfDSN(filepath string, mysqlConf ...string) (string, error) {
	// load config file by file type
	switch {
	case strings.HasSuffix(filepath, "ini"):
		loadIni(filepath, mysqlConf...)
	case strings.HasSuffix(filepath, "yml"):
		loadYml(filepath, mysqlConf...)
	case strings.HasSuffix(filepath, "env"):
		loadEnv(filepath, mysqlConf...)
	default:
		return "", errors.New("error: filepath is required")
	}

	// dbname
	if len(cfg.dbname) == 0 {
		return "", errors.New("error: dbname is required")
	}

	// dsn
	dsn := cfg.formatDSN()
	return dsn, nil
}

func contains(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), substr)
}

func loadIni(filepath string, mysqlConf ...string) {
	// load ini
	file, err := ini.Load(filepath)
	if err != nil {
		log.Fatal("error: Failed to load ini file")
	}

	for _, v := range mysqlConf {
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

func loadYml(filepath string, mysqlConf ...string) {
	viper.SetConfigType("yml")
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error: Failed to load yml file")
	}

	for _, v := range mysqlConf {
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

func loadEnv(filepath string, mysqlConf ...string) {
	if err := godotenv.Load(filepath); err != nil {
		log.Fatal("error: Failed to load env file")
	}

	for _, v := range mysqlConf {
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

func (cfg *mysqlConf) formatDSN() string {
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
		buf.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	}
	return buf.String()
}
