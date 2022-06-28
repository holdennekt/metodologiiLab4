package dataproviders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	initQuery                = "drop table if exists tasks; create table tasks (id bigserial not null primary key, title varchar(50) not null, details varchar(200), deadline date, expired boolean not null default false, completed boolean not null default false, completed_at date)"
	listAllTasksTestQuery    = "insert into tasks (title, details, deadline) values ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-29'); insert into tasks (title, details, deadline, completed, completed_at) values ('make some tea', 'black tea with 2 spoons of sugar', '2022-06-28', true, '2022-06-28');"
	listActiveTasksTestQuery = "insert into tasks (title, details, deadline) values ('get beraly good at css', 'learn css a bit, even level \"not being disgusted while formatting small pet project\" will be fine', '2022-08-31'), ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-29'); insert into tasks (title) values ('watch Neon Genesis Evangelion'); insert into tasks (title, details, deadline, completed, completed_at) values ('make some tea', 'black tea with 2 spoons of sugar', '2022-06-28', true, '2022-06-28'), ('turn 19', 'yea, that`s it', '2022-04-24', true, '2022-04-23'); insert into tasks (title, details, deadline, expired) values ('read some books', 'science fiction like aisek azimov is fine, as well as rey breadberry', '2022-05-23', true), ('find a gift to the friend', 'maybe paper model constructor, maybe some arduino stuff', '2022-06-28', true);"
)

func TestListAllTasks(t *testing.T) {
	conn := Connection{
		Host:       "localhost",
		Port:       5432,
		User:       "nikitagryshchak",
		Password:   "",
		DbName:     "todo",
		DisableSSL: true,
	}
	dataProvider, err := NewDataProvider(conn)
	if err != nil {
		t.Fatalf("error while creating db: %v", err)
	}
	defer dataProvider.db.Close()

	_, err = dataProvider.db.Query(initQuery)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	_, err = dataProvider.db.Query(listAllTasksTestQuery)
	if err != nil {
		t.Fatalf("failed to fill db: %v", err)
	}
	res, err := dataProvider.ListAllTasks()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0].Id)
	assert.Equal(t, 2, res[1].Id)
}

func TestListActiveTasks(t *testing.T) {
	conn := Connection{
		Host:       "localhost",
		Port:       5432,
		User:       "nikitagryshchak",
		Password:   "",
		DbName:     "todo",
		DisableSSL: true,
	}
	dataProvider, err := NewDataProvider(conn)
	if err != nil {
		t.Fatalf("error while creating db: %v", err)
	}
	defer dataProvider.db.Close()
	_, err = dataProvider.db.Query(initQuery)
	if err != nil {
		t.Fatalf("failed to init db: %v", err)
	}
	_, err = dataProvider.db.Query(listActiveTasksTestQuery)
	if err != nil {
		t.Fatalf("failed to fill db: %v", err)
	}
	res, err := dataProvider.ListActiveTasks()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(res))
	assert.Equal(t, 2, res[0].Id)
	assert.Equal(t, 1, res[1].Id)
	assert.Equal(t, 3, res[2].Id)
}
