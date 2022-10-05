package gopkg

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func ConfigEnv(envArr []string) (string, error) {
	if envArr == nil || len(envArr) != 4 {
		return "", errors.New("error: env配置数据不能为空 或 配置数量不全 ")
	}

	if err := godotenv.Load(); err != nil {
		return "", errors.New("error: Failed to load env file")
	}

	dbUser := os.Getenv(envArr[0])
	dbPass := os.Getenv(envArr[1])
	dbHost := os.Getenv(envArr[2])
	dbName := os.Getenv(envArr[3])

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbName,
	)

	return dsn, nil
}

// func ConfigEnvTest() {
// 	testEnvArr := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_NAME"}
//
// 	dsn, _ := ConfigEnv(testEnvArr)
// 	fmt.Println(dsn)
//
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
