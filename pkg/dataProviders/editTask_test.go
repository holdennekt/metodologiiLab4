package dataproviders

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	dataForEditTask      = "insert into tasks (title, details, deadline) values ('get beraly good at css', 'learn css a bit, even level \"not being disgusted while formatting small pet project\" will be fine', '2022-08-31'), ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-28'); insert into tasks (title) values ('watch Neon Genesis Evangelion'); insert into tasks (title, details, deadline, completed, completed_at) values ('make some tea', 'black tea with 2 spoons of sugar', '2022-06-28', true, '2022-06-28')"
	dataForMarkCompleted = "insert into tasks (title, details, deadline) values ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-28'); insert into tasks (title) values ('watch Neon Genesis Evangelion')"
)

func TestEditTask(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForEditTask)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	task, err := dataProvider.EditTask(1, "become css senior", "basicly, you need to be the css god", "2022-07-01")
	assert.Nil(t, err)
	assert.Equal(t, "become css senior", task.Title)
	assert.Equal(t, "basicly, you need to be the css god", task.Details.String)
	assert.Equal(t, true, task.Details.Valid)
	assert.Equal(t, "2022-07-01", task.Deadline.Time.String()[:10])
	assert.Equal(t, true, task.Deadline.Valid)
	task, err = dataProvider.EditTask(2, "", "please, finish finally this fking lab", "2022-06-30")
	assert.Nil(t, err)
	assert.Equal(t, "finish 4 lab", task.Title)
	assert.Equal(t, "please, finish finally this fking lab", task.Details.String)
	assert.Equal(t, true, task.Details.Valid)
	assert.Equal(t, "2022-06-30", task.Deadline.Time.String()[:10])
	assert.Equal(t, true, task.Deadline.Valid)
	task, err = dataProvider.EditTask(3, "", "", "2022-09-01")
	assert.Nil(t, err)
	assert.Equal(t, "watch Neon Genesis Evangelion", task.Title)
	assert.Equal(t, "", task.Details.String)
	assert.Equal(t, false, task.Details.Valid)
	assert.Equal(t, "2022-09-01", task.Deadline.Time.String()[:10])
	assert.Equal(t, true, task.Deadline.Valid)
	task, err = dataProvider.EditTask(4, "", "", "")
	assert.Nil(t, err)
	assert.Equal(t, "make some tea", task.Title)
	assert.Equal(t, "black tea with 2 spoons of sugar", task.Details.String)
	assert.Equal(t, true, task.Details.Valid)
	assert.Equal(t, "2022-06-28", task.Deadline.Time.String()[:10])
	assert.Equal(t, true, task.Deadline.Valid)
}

func TestMarkCompleted(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForEditTask)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	err = dataProvider.MarkCompleted(2)
	assert.Nil(t, err)
	tasks, err := dataProvider.ListAllTasks()
	if err != nil {
		t.Fatalf("error while listing all tasks: %v", err)
	}
	assert.Equal(t, false, tasks[0].Completed)
	assert.Equal(t, true, tasks[1].Completed)
	assert.Equal(t, time.Now().String()[:10], tasks[1].CompletedAt.Time.String()[:10])
	assert.Equal(t, true, tasks[1].CompletedAt.Valid)
}
