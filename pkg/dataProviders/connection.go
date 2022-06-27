package dataproviders

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DataProvider struct {
	db sql.DB
}

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
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

func NewDataProvider(conn Connection) (*DataProvider, error) {
	url := conn.String()
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &DataProvider{db: *db}, nil
}

type Task struct {
}
