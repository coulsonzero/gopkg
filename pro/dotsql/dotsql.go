package dotsql

type DotSql struct {
	queries map[string]string
}

func LoadFromFile(sqlFile string) (*DotSql, error)

func (d DotSql) LookupQuery(name string) (query string, err error)

func Get(match string)
