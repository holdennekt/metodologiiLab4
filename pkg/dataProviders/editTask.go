package dataproviders

import (
	"fmt"
	"time"
)

func (dp *DataProvider) EditTask(id int, title, details, deadline string) (*Task, error) {
	if title != "" {
		_, err := dp.db.Exec(fmt.Sprintf("update tasks set title='%s' where id=%d", title, id))
		if err != nil {
			return nil, err
		}
	}
	if details != "" {
		_, err := dp.db.Exec(fmt.Sprintf("update tasks set details='%s' where id=%d", details, id))
		if err != nil {
			return nil, err
		}
	}
	if deadline != "" {
		_, err := dp.db.Exec(fmt.Sprintf("update tasks set deadline='%s' where id=%d", deadline, id))
		if err != nil {
			return nil, err
		}
	}
	row := dp.db.QueryRow(fmt.Sprintf("select * from tasks where id=%d", id))
	var t Task
	err := row.Scan(
		&t.Id,
		&t.Title,
		&t.Details,
		&t.Deadline,
		&t.Expired,
		&t.Completed,
		&t.CompletedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (dp *DataProvider) MarkCompleted(id int) error {
	date := time.Now().String()[:10]
	queryF := "update tasks set completed=true, completed_at='%s' where id=%d"
	_, err := dp.db.Exec(fmt.Sprintf(queryF, date, id))
	if err != nil {
		return nil
	}
	return nil
}
