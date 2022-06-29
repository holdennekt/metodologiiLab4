package dataproviders

import "time"

// TODO: implement db access functions:
// EditTask

func (dp *DataProvider) EditTask(id int, title, details, deadline string) (*Task, error) {
	return nil, nil
}

func (dp *DataProvider) MarkCompleted(id int) (*time.Time, error) {
	return nil, nil
}
