package pro

// ConfDSN  read mysql config to dsn from config file
// @Return: dsn string
// @Params: At least one dbname parameter is required
// @FileTypes: The following file types are supported: ini yml env
func ConfDSN(filepath string, mysqlConfig ...string) (string, error)
