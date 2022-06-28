package dataproviders

import (
	"fmt"
	"strings"
)

// TODO: implement db access functions:
// NewTask

type NewTask struct {
	title, details, deadline string
}

func (nt *NewTask) toMap() map[string]string {
	return map[string]string{
		"title":    nt.title,
		"details":  nt.details,
		"deadline": nt.deadline,
	}
}

func (dp *DataProvider) NewTask(newT NewTask) (*Task, error) {
	cols, vals := make([]string, 0), make([]string, 0)
	for key, val := range newT.toMap() {
		if val != "" {
			cols = append(cols, key)
			vals = append(vals, fmt.Sprintf("'%s'", val))
		}
	}
	colsStr := strings.Join(cols, ",")
	valsStr := strings.Join(vals, ",")
	query := fmt.Sprintf("insert into tasks (%s) values (%s) returning *", colsStr, valsStr)
	row := dp.db.QueryRow(query)
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
