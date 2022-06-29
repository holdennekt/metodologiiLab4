package dataproviders

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int
	Title       string
	Details     sql.NullString
	Deadline    sql.NullTime
	Expired     bool
	Completed   bool
	CompletedAt sql.NullTime
}

type Repository interface {
	ListAllTasks() ([]*Task, error)
	ListActiveTasks() ([]*Task, error)
	ListExpiredTasks() ([]*Task, error)
	NewTask(title, details, deadline string) (*Task, error)
	EditTask(id int, title, details, deadline string) (*Task, error)
	MarkCompleted(id int) (*time.Time, error)
	DeleteTask(id int) (*Task, error)
	UpdateState() error
}

type DataProvider struct {
	db *sql.DB
}

func NewDataProvider(conn Connection) (*DataProvider, error) {
	db, err := conn.Open()
	if err != nil {
		return nil, err
	}
	return &DataProvider{db}, nil
}

func (dp *DataProvider) UpdateState() error {
	return nil
}
