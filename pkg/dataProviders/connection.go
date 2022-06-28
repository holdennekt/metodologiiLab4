package dataproviders

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq"
)

type DataProvider struct {
	db *sql.DB
}

func (dp *DataProvider) UpdateState() error {
	tasks, err := dp.getTasks("select * from tasks where expired=false and completed=false and deadline is not null")
	if err != nil {
		return err
	}
	for _, task := range tasks {
		year, month, day := time.Now().Date()
		loc, _ := time.LoadLocation("UTC")
		date := time.Date(year, month, day, 0, 0, 0, 0, loc)
		fmt.Println("id:", task.Id, "deadline:", task.Deadline.Time, "now:", date)
		if date.After(task.Deadline.Time) {
			err = dp.markExpired(task.Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (dp *DataProvider) markExpired(id int) error {
	_, err := dp.db.Exec(fmt.Sprintf("update tasks set expired=true where id=%d", id))
	if err != nil {
		return err
	}
	return nil
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
