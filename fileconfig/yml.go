package fileconfig

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

func ConfigYml(filename string, envArr []string) (string, error) {
	if envArr == nil || len(envArr) != 4 {
		return "", errors.New("Error: env配置数据不能为空 或 配置数量不全 !")
	}

	viper.SetConfigType("yml")
	viper.SetConfigFile(filename)

	if err := viper.ReadInConfig(); err != nil {
		return "", errors.New("Error: Failed to load yml file")
	}

	dbUser := viper.GetString(envArr[0])
	dbPass := viper.GetString(envArr[1])
	dbHost := viper.GetString(envArr[2])
	dbName := viper.GetString(envArr[3])

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)
	return dsn, nil
}

// func TestConfig() {
// 	// 初始化yml配置
// 	testEnvArr := []string{"mysql.username", "mysql.password", "mysql.host", "mysql.database"}
// 	dsn, err := ConfigYml("./config.yml", testEnvArr)
// 	if err != nil {
// 		log.Fatal("Error: load yml config")
// 	}
// 	fmt.Println(dsn)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			SingularTable: true, // would be `user`
// 		},
// 	})
//
// 	if err != nil {
// 		log.Fatal("Error: Failed to create a connection to database")
// 	}
//
// 	log.Fatalf("Success connect to mysql. %v", db)
// }
