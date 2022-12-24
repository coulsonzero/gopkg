package pro

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	_ "unsafe" // for go:linkname
)

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

const (
	defaultMatch string = "name"
	sep          string = " >> "
)

func Load(sqlFile string) (*ZeroSql, error) {
	file, err := os.Open(sqlFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var title []string
	var context []string
	current := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "--") {
			if len(current) != 0 {
				current = strings.TrimSpace(current)
				context = append(context, current)
				current = ""
			}
			// title = append(title, strings.Split(line, ": ")[1])
			// title = append(title, line)
			title = append(title, getByTag(line))
		} else {
			if len(current) > 0 {
				current = current + "\n"
			}
			current = current + line
		}
	}
	current = strings.TrimSpace(current)
	context = append(context, current)
	current = ""

	m := make(map[string]string)
	for i := 0; i < len(title); i++ {
		m[title[i]] = context[i]
	}

	res := &ZeroSql{
		queries:  m,
		matchTag: defaultMatch,
	}

	return res, nil
}

// PrintResult
//  @return print the map result in the console.
func (d *ZeroSql) PrintResult() {
	for k, v := range d.queries {
		fmt.Printf("key: %s\nvalue: %s\n", k, v)
	}
}

// SetTag
//  @Description: filter the result by match, only return the result of startsWith the match.
//  @param match like "name", "type", "tag" and so on.
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.SetTag("type")
//  zeroSql.LookupQuery("create-user")
func (d *ZeroSql) SetTag(match string) {
	d.matchTag = match
	m := make(map[string]string)
	for k, v := range d.queries {
		if strings.Split(k, sep)[0] == match {
			m[k] = v
		}
	}

	d.queries = m
}

func getByTag(line string) string {
	re := regexp.MustCompile("^\\s*--\\s*\\W?(\\w+)\\W?\\s*(\\S+)")
	// re := regexp.MustCompile("^\\s*--\\s*name:\\s*(\\S+)")

	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return ""
	}
	return matches[1] + sep + matches[2]
}

// LookupQuery
//  @Description: query the sql by the name commit
//  @param name like "create-user"
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.LookupQuery("create-user")	// default by "name"
//  if SetTag("tag"), then only return the result of startsWith the match "tag"
func (d *ZeroSql) LookupQuery(name string) (query string, err error) {
	key := d.matchTag + sep + name
	query, ok := d.queries[key]
	if !ok {
		err = fmt.Errorf("sql: '%s' could not be found", name)
	}

	return
}

// LookupQueryAny
//  @Usage
//  zeroSql, _ := Load(filepath)
//  zeroSql.LookupQueryAny("create-user")	// default by any match
func (d *ZeroSql) LookupQueryAny(name string) (query string, err error) {
	for k, v := range d.queries {
		key := strings.Split(k, sep)[1]
		if key == name {
			return v, nil
		}
	}
	return "", fmt.Errorf("sql: '%s' could not be found", name)
}

func (d ZeroSql) Get(match string) *QuerySql {
	d.matchTag = match
	m := make(map[string]string)
	for k, v := range d.queries {
		if strings.Split(k, sep)[0] == match {
			m[k] = v
		}
	}

	d.queries = m
	return &QuerySql{queries: m}
}

// QueryAny
//  @Usage
//  query, _ := zeroSql.Get("name").QueryAny("create-users-table")
//  fmt.Println(query)
func (d *QuerySql) QueryAny(name string) (query string, err error) {
	for k, v := range d.queries {
		key := strings.Split(k, sep)[1]
		if key == name {
			return v, nil
		}
	}
	return "", fmt.Errorf("sql: '%s' could not be found", name)
}
