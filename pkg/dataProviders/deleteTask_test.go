package dataproviders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	dataForDeleteTask = "insert into tasks (title, details, deadline) values ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-29'); insert into tasks (title) values ('watch Neon Genesis Evangelion')"
)

func TestDeleteTask(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForDeleteTask)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	task, err := dataProvider.DeleteTask(2)
	assert.NotNil(t, task)
	assert.Nil(t, err)
	tasks, err := dataProvider.getTasks("select * from tasks where id=2")
	if err != nil {
		t.Fatalf("error while geting tasks: %v", err)
	}
	assert.Equal(t, 0, len(tasks))
}
