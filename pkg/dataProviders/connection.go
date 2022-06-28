package dataproviders

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

type DataProvider struct {
	db *sql.DB
}

func NewDataProvider(conn Connection) (*DataProvider, error) {
	url := conn.String()
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &DataProvider{db}, nil
}

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

type Task struct {
	Id          int
	Title       string
	Details     sql.NullString
	Deadline    sql.NullTime
	Expired     bool
	Completed   bool
	CompletedAt sql.NullTime
}
