package commands

import (
	"database/sql"
	"fmt"
	"sort"
	"time"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type testRepo struct {
	tasks []*dataproviders.Task
}

func (tr *testRepo) filter(test func(*dataproviders.Task) bool) []*dataproviders.Task {
	filtered := make([]*dataproviders.Task, 0)
	for _, task := range tr.tasks {
		if test(task) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func NewTestRepo() *testRepo {
	nullTime := time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))
	task1 := &dataproviders.Task{
		Id:          1,
		Title:       "test task 1",
		Details:     sql.NullString{String: "this is an active task", Valid: true},
		Deadline:    sql.NullTime{Time: time.Date(2022, time.Month(9), 1, 0, 0, 0, 0, time.FixedZone("UTC", 0)), Valid: true},
		Expired:     false,
		Completed:   false,
		CompletedAt: sql.NullTime{Time: nullTime},
	}
	task2 := &dataproviders.Task{
		Id:          2,
		Title:       "test task 2",
		Details:     sql.NullString{String: "this is an expired task", Valid: true},
		Deadline:    sql.NullTime{Time: time.Date(2022, time.Month(6), 22, 0, 0, 0, 0, time.FixedZone("UTC", 0)), Valid: true},
		Expired:     true,
		Completed:   false,
		CompletedAt: sql.NullTime{Time: nullTime},
	}
	task3 := &dataproviders.Task{
		Id:          3,
		Title:       "test task 3",
		Details:     sql.NullString{String: "this is a completed task", Valid: true},
		Deadline:    sql.NullTime{Valid: false},
		Expired:     false,
		Completed:   true,
		CompletedAt: sql.NullTime{Time: time.Date(2022, time.Month(6), 28, 0, 0, 0, 0, time.FixedZone("UTC", 0)), Valid: true},
	}
	task4 := &dataproviders.Task{
		Id:          4,
		Title:       "test task 4",
		Details:     sql.NullString{String: "this task has been expired at 2022-06-28, but not marked yet", Valid: true},
		Deadline:    sql.NullTime{Time: time.Date(2022, time.Month(6), 28, 0, 0, 0, 0, time.FixedZone("UTC", 0)), Valid: true},
		Expired:     false,
		Completed:   false,
		CompletedAt: sql.NullTime{Time: nullTime},
	}
	task5 := &dataproviders.Task{
		Id:          5,
		Title:       "test task 5",
		Details:     sql.NullString{String: "one more active task", Valid: true},
		Deadline:    sql.NullTime{Time: time.Date(2022, time.Month(8), 1, 0, 0, 0, 0, time.FixedZone("UTC", 0)), Valid: true},
		Expired:     false,
		Completed:   false,
		CompletedAt: sql.NullTime{Time: nullTime},
	}
	return &testRepo{tasks: []*dataproviders.Task{task1, task2, task3, task4, task5}}
}

func (tr *testRepo) ListAllTasks() ([]*dataproviders.Task, error) {
	return tr.tasks, nil
}

func (tr *testRepo) ListActiveTasks() ([]*dataproviders.Task, error) {
	filtered := tr.filter(func(task *dataproviders.Task) bool {
		return !task.Completed && !task.Expired
	})
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Deadline.Time.Before(filtered[j].Deadline.Time)
	})
	return filtered, nil
}

func (tr *testRepo) ListExpiredTasks() ([]*dataproviders.Task, error) {
	filtered := tr.filter(func(task *dataproviders.Task) bool {
		return task.Expired
	})
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Deadline.Time.Before(filtered[j].Deadline.Time)
	})
	return filtered, nil
}

func (tr *testRepo) NewTask(title, details, deadline string) (*dataproviders.Task, error) {
	nullTime := time.Date(1, time.Month(1), 1, 0, 0, 0, 0, time.FixedZone("UTC", 0))
	var sqlTime sql.NullTime
	if deadline != "" {
		parsed, err := time.Parse("2006-01-02", deadline)
		if err != nil {
			return nil, err
		}
		sqlTime = sql.NullTime{Time: parsed, Valid: true}
	} else {
		sqlTime = sql.NullTime{Time: nullTime, Valid: false}
	}
	task := &dataproviders.Task{
		Id:          len(tr.tasks) + 1,
		Title:       title,
		Details:     sql.NullString{String: details, Valid: details != ""},
		Deadline:    sqlTime,
		Expired:     false,
		Completed:   false,
		CompletedAt: sql.NullTime{Time: nullTime, Valid: false},
	}
	tr.tasks = append(tr.tasks, task)
	return task, nil
}

func (tr *testRepo) EditTask(id int, title, details, deadline string) (*dataproviders.Task, error) {
	var task *dataproviders.Task = nil
	for _, t := range tr.tasks {
		if t.Id == id {
			task = t
		}
	}
	if task == nil {
		return nil, fmt.Errorf("no task with id: %d", id)
	}
	if title != "" {
		task.Title = title
	}
	if details != "" {
		task.Details = sql.NullString{String: details, Valid: true}
	}
	if deadline != "" {
		parsed, err := time.Parse("2006-01-02", deadline)
		if err != nil {
			return nil, err
		}
		task.Deadline = sql.NullTime{Time: parsed, Valid: true}
	}
	return task, nil
}

func (tr *testRepo) MarkCompleted(id int) (*time.Time, error) {
	var task *dataproviders.Task = nil
	for _, t := range tr.tasks {
		if t.Id == id {
			task = t
		}
	}
	if task == nil {
		return nil, fmt.Errorf("no task with id: %d", id)
	}
	task.Completed = true
	now := time.Now()
	task.CompletedAt = sql.NullTime{Time: now, Valid: true}
	return &now, nil
}

func (tr *testRepo) DeleteTask(id int) (*dataproviders.Task, error) {
	var task *dataproviders.Task = nil
	var taskIndex int = -1
	for i, t := range tr.tasks {
		if t.Id == id {
			task = t
			taskIndex = i
		}
	}
	if task == nil {
		return nil, fmt.Errorf("no task with id: %d", id)
	}
	tr.tasks = append(tr.tasks[:taskIndex], tr.tasks[taskIndex+1:]...)
	return task, nil
}

func (tr *testRepo) UpdateState() error {
	for _, task := range tr.tasks {
		if task.Expired || task.Completed || !task.Deadline.Valid {
			continue
		}
		year, month, day := time.Now().Date()
		loc := time.FixedZone("UTC", 0)
		date := time.Date(year, month, day, 0, 0, 0, 0, loc)
		if date.After(task.Deadline.Time) {
			task.Expired = true
		}
	}
	return nil
}
