package gopkg

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// ConfigEnv
// Deprecated: Use ConfDSN with options instead.
func ConfigEnv(filepath string, envArr []string) (string, error) {
	if envArr == nil || len(envArr) != 5 {
		return "", errors.New("error: 配置数据不能为空 或 配置数量不全 ")
	}

	if err := godotenv.Load(filepath); err != nil {
		return "", errors.New("error: Failed to load env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv(envArr[0]),
		os.Getenv(envArr[1]),
		os.Getenv(envArr[2]),
		os.Getenv(envArr[3]),
		os.Getenv(envArr[4]),
	)

	return dsn, nil
}
