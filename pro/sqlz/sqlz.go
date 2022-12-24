package sqlz

/**
 *
 * zeroSql, _ := Load("query.sql")
 *
 *	// zeroSql.PrintResult()
 *	// println(zeroSql.matchTag)
 *  // zeroSql.SetTag("type")
 *
 *  =========== first: default tag by name ===========
 *
 *	// query, _ := zeroSql.LookupQuery("create-users-table")
 *	// query, _ := zeroSql.LookupQuery("create-user")
 * 	// query, _ := zeroSql.LookupQuery("find-users-by-email")
 *	// query, _ := zeroSql.LookupQuery("drop-users-table")
 *
 *  =========== second ===========
 *	query, err := zeroSql.LookupQueryAny("drop-users-table")
 *	if err != nil {
 *		log.Fatal(err)
 *	}
 *	fmt.Println(query)	// DROP TABLE users;
 *
 *  =========== third ===========
 *  println(zeroSql.Get("type").QueryAny("find-one-user-by-email"))
 *	println(zeroSql.Get("type").QueryAny("find-users-by-email"))
 *	println(zeroSql.Get("tag").QueryAny("drop-users-table"))
 * 	query, _ := zeroSql.Get("name").QueryAny("create-users-table")
 *	fmt.Println(query)
 */

type ZeroSql struct {
	queries  map[string]string
	matchTag string
}

type QuerySql struct {
	queries map[string]string
}

/* query.sql looks like:
-- name: create-users-table
CREATE TABLE users (
id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
name VARCHAR(255),
email VARCHAR(255)
);

-- name: create-user
INSERT INTO users (name, email) VALUES(?, ?);

-- type: find-users-by-email
SELECT id,name,email FROM users WHERE email = ?;

-- @type find-one-user-by-email
SELECT id,name,email FROM users WHERE email = ? LIMIT 1;

-- @tag drop-users-table
DROP TABLE users;
*/

func Load(sqlFile string) (*ZeroSql, error)

// PrintResult
//  @return print the map result in the console.
func (d *ZeroSql) PrintResult()

// SetTag
//  @Description: filter the result by match, only return the result of startsWith the match.
//  @param match like "name", "type", "tag" and so on.
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.SetTag("type")
//  zeroSql.LookupQuery("create-user")
func (d *ZeroSql) SetTag(match string)

func (d ZeroSql) Get(match string) *QuerySql

// LookupQuery
//  @Description: query the sql by the name commit
//  @param name like "create-user"
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.LookupQuery("create-user")	// default by "name"
//  if SetTag("tag"), then only return the result of startsWith the match "tag"
func (d *ZeroSql) LookupQuery(name string) (query string, err error)

// LookupQueryAny
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.LookupQueryAny("create-user")	// default by any match
func (d *ZeroSql) LookupQueryAny(name string) (query string, err error)

// QueryAny
//  @Usage
//  query, _ := zeroSql.Get("name").QueryAny("create-users-table")
//  fmt.Println(query)
func (d *QuerySql) QueryAny(name string) (query string, err error)
