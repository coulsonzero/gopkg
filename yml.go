package gopkg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// ConfigYml Deprecated
func ConfigYml(filepath string, ymlArr []string) (string, error) {
	if ymlArr == nil || len(ymlArr) != 5 {
		log.Fatal("error: 配置数据不能为空 或 配置数量不全 ")
	}

	viper.SetConfigType("yml")
	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error: Failed to load yml file")
	}

	dbUser := viper.GetString(ymlArr[0])
	dbPass := viper.GetString(ymlArr[1])
	dbHost := viper.GetString(ymlArr[2])
	dbPort := viper.GetString(ymlArr[3])
	dbName := viper.GetString(ymlArr[4])

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	return dsn, nil
}
