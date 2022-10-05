package gopkg

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
)

func ConfigInt(filename string, envArr []string) (string, error) {
	if envArr == nil || len(envArr) != 4 {
		return "", errors.New("Error: env配置数据不能为空 或 配置数量不全 !")
	}

	file, err := ini.Load(filename)
	if err != nil {
		return "", errors.New("Failed to load ini file")
	}

	dbUser := file.Section("mysql").Key(envArr[0]).String()
	dbPass := file.Section("mysql").Key(envArr[1]).String()
	dbHost := file.Section("mysql").Key(envArr[2]).String()
	dbName := file.Section("mysql").Key(envArr[3]).String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)
	return dsn, nil

}
