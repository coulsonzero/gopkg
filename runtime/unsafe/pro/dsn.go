package pro

import (
	"strings"
	_ "unsafe" // for go:linkname
)

//go:linkname confDsn github.com/coulsonzero/gopkg/pro.ConfDSN
// @Return: a string dsn
// @Params: At least one dbname parameter is required
// @FileTypes: The following file types are supported: ini yml env
func confDsn(filepath string, mysqlConfig ...string) (string, error) {
	// load config file by file type
	switch {
	case strings.HasSuffix(filepath, "ini"):
		loadIni(filepath, mysqlConfig...)
	case strings.HasSuffix(filepath, "yml"):
		loadYml(filepath, mysqlConfig...)
	case strings.HasSuffix(filepath, "env"):
		loadEnv(filepath, mysqlConfig...)
	default:
		logger("error: filepath is required")
	}

	// check all mysql params is valid
	// dbname especially
	cfg.isValid()

	// dsn
	dsn := cfg.formatDSN()
	return dsn, nil
}
