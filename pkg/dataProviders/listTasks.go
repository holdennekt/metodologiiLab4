package dataproviders

func (dp *DataProvider) getTasks(sqlQuery string) ([]*Task, error) {
	rows, err := dp.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task, 0)
	for rows.Next() {
		var t Task
		err := rows.Scan(
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
		tasks = append(tasks, &t)
	}
	return tasks, nil
}

func (dp *DataProvider) ListAllTasks() ([]*Task, error) {
	return dp.getTasks("select * from tasks")
}

func (dp *DataProvider) ListActiveTasks() ([]*Task, error) {
	return dp.getTasks("select * from tasks where completed=false and expired=false order by deadline asc")
}

func (dp *DataProvider) ListExpiredTasks() ([]*Task, error) {
	return dp.getTasks("select * from tasks where expired=true order by deadline asc")
}
