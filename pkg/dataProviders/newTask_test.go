package dataproviders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	dataForNewTask = "insert into tasks (title) values ('watch Neon Genesis Evangelion')"
)

func TestNewTask(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForNewTask)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	newTask := NewTask{
		title:    "finish 4th lab",
		details:  "finish implementing commands",
		deadline: "2022-06-29",
	}
	task, err := dataProvider.NewTask(newTask)
	assert.Nil(t, err)
	assert.Equal(t, "finish 4th lab", task.Title)
	assert.Equal(t, "finish implementing commands", task.Details.String)
	assert.Equal(t, true, task.Details.Valid)
	assert.Equal(t, "2022-06-29", task.Deadline.Time.String()[:10])
	assert.Equal(t, true, task.Deadline.Valid)
	newTask = NewTask{
		title:    "go to sleep, pls",
		details:  "seriously, you need some rest",
		deadline: "",
	}
	task, err = dataProvider.NewTask(newTask)
	assert.Nil(t, err)
	assert.Equal(t, "go to sleep, pls", task.Title)
	assert.Equal(t, "seriously, you need some rest", task.Details.String)
	assert.Equal(t, true, task.Details.Valid)
	assert.Equal(t, false, task.Deadline.Valid)
	newTask = NewTask{
		title:    "eat a cherry",
		details:  "",
		deadline: "",
	}
	task, err = dataProvider.NewTask(newTask)
	assert.Nil(t, err)
	assert.Equal(t, "eat a cherry", task.Title)
	assert.Equal(t, false, task.Details.Valid)
	assert.Equal(t, false, task.Deadline.Valid)
	newTask = NewTask{
		title:    "",
		details:  "",
		deadline: "",
	}
	task, err = dataProvider.NewTask(newTask)
	assert.Nil(t, task)
	assert.Error(t, err)
}
