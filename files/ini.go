package gopkg

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
)

// ConfigInt
// Deprecated: Use ConfDSN with options instead.
func ConfigInt(filepath string, envArr []string) (string, error) {
	if len(envArr) != 5 {
		return "", errors.New("error: env配置数据不能为空 或 配置数量不全 ")
	}

	file, err := ini.Load(filepath)
	if err != nil {
		return "", errors.New("failed to load ini file")
	}

	dbUser := file.Section("mysql").Key(envArr[0]).String()
	dbPass := file.Section("mysql").Key(envArr[1]).String()
	dbHost := file.Section("mysql").Key(envArr[2]).String()
	dbPort := file.Section("mysql").Key(envArr[3]).String()
	dbName := file.Section("mysql").Key(envArr[4]).String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	return dsn, nil
}
