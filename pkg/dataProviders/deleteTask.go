package dataproviders

import "fmt"

func (dp *DataProvider) DeleteTask(id int) (*Task, error) {
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
	_, err = dp.db.Exec(fmt.Sprintf("delete from tasks where id=%d", id))
	if err != nil {
		return nil, err
	}
	return &t, nil
}
