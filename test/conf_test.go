package test

var (
	envArr = []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	ymlArr = []string{"mysql.username", "mysql.password", "mysql.host", "mysql.port", "mysql.dbname"}
	DSN    = "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
)

// func TestConfigEnv(t *testing.T) {
// 	dsn, err := gopkg.ConfigEnv("conf/.env", envArr)
// 	if err != nil {
// 		// return
// 		t.Error(err)
// 	}
//
// 	if dsn != DSN {
// 		t.Errorf("error: not equal: %q", dsn)
// 	}
// }

// func TestConfigYml(t *testing.T) {
// 	dsn, err := gopkg.ConfigYml("conf/config.yml", ymlArr)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if dsn != DSN {
// 		t.Errorf("error: not equal: %q", dsn)
// 	}
// }
