package dataproviders

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

type Connection struct {
	Host       string
	Port       int
	User       string
	Password   string
	DbName     string
	DisableSSL bool
}

func (c *Connection) String() string {
	dbUrl := &url.URL{
		Scheme: "postgres",
		Host:   c.Host,
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.DbName,
	}
	if c.DisableSSL {
		dbUrl.RawQuery = url.Values{
			"sslmode": []string{"disable"},
		}.Encode()
	}
	return dbUrl.String()
}

func (c *Connection) Open() (*sql.DB, error) {
	return sql.Open("postgres", c.String())
}
