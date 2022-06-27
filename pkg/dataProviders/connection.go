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
	userInfo := fmt.Sprintf("%s:%s", c.User, c.Password)
	host := fmt.Sprintf("%s:%d", c.Host, c.Port)
	url := fmt.Sprintf("postgresql://%s@%s/%s?sslmode=disable", userInfo, host, c.DbName)
	return url
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
