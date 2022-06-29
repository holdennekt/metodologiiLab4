package dataproviders

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	initQuery          = "drop table if exists tasks; create table tasks (id bigserial not null primary key, title varchar(50) not null, details varchar(200), deadline date, expired boolean not null default false, completed boolean not null default false, completed_at date)"
	dataForMarkExpired = "insert into tasks (title, details, deadline) values ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-29'); insert into tasks (title) values ('watch Neon Genesis Evangelion')"
	dataForUpdateState = "insert into tasks (title, details, deadline) values ('get beraly good at css', 'learn css a bit, even level \"not being disgusted while formatting small pet project\" will be fine', '2022-08-31'), ('finish 4 lab', 'yes, 4th wall was crushed', '2022-06-28'); insert into tasks (title) values ('watch Neon Genesis Evangelion'); insert into tasks (title, details, deadline, completed, completed_at) values ('make some tea', 'black tea with 2 spoons of sugar', '2022-06-28', true, '2022-06-28'), ('turn 19', 'yea, that`s it', '2022-04-24', true, '2022-04-23'); insert into tasks (title, details, deadline) values ('read some books', 'science fiction like aisek azimov is fine, as well as rey breadberry', '2022-05-23'), ('find a gift to the friend', 'maybe paper model constructor, maybe some arduino stuff', '2022-06-28')"
)

func getTestDataProvider(query string) (*DataProvider, error) {
	conn := Connection{
		Host:       "localhost",
		Port:       5432,
		User:       "nikitagryshchak",
		Password:   "",
		DbName:     "todotest",
		DisableSSL: true,
	}
	dataProvider, err := NewDataProvider(conn)
	if err != nil {
		return nil, fmt.Errorf("error while creating db: %v", err)
	}
	_, err = dataProvider.db.Exec(initQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %v", err)
	}
	_, err = dataProvider.db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fill db: %v", err)
	}
	return dataProvider, nil
}

func TestNewDataProvider(t *testing.T) {
	conn := Connection{
		Host:       "localhost",
		Port:       5432,
		User:       "nikitagryshchak",
		Password:   "",
		DbName:     "todotest",
		DisableSSL: true,
	}
	dataProvider, err := NewDataProvider(conn)
	if err != nil {
		t.Fatalf("error while creating db: %v", err)
	}
	defer dataProvider.db.Close()
	if err := dataProvider.db.Ping(); err != nil {
		t.Fatalf("error while pinging db: %v", err)
	}
}

func TestMarkExpired(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForMarkExpired)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	err = dataProvider.markExpired(1)
	assert.Nil(t, err)
	tasks, err := dataProvider.ListAllTasks()
	if err != nil {
		t.Fatalf("error while listing all tasks: %v", err)
	}
	fmt.Println(tasks[0], tasks[1])
	assert.Equal(t, true, tasks[0].Expired)
	assert.Equal(t, false, tasks[1].Expired)
}

func TestUpdateState(t *testing.T) {
	dataProvider, err := getTestDataProvider(dataForUpdateState)
	if err != nil {
		t.Fatal(err)
	}
	defer dataProvider.db.Close()
	err = dataProvider.UpdateState()
	assert.Nil(t, err)
	tasks, err := dataProvider.ListAllTasks()
	if err != nil {
		t.Fatalf("failed to list all tasks: %v", err)
	}
	assert.Equal(t, false, tasks[0].Expired)
	assert.Equal(t, true, tasks[1].Expired)
	assert.Equal(t, false, tasks[2].Expired)
	assert.Equal(t, false, tasks[3].Expired)
	assert.Equal(t, false, tasks[4].Expired)
	assert.Equal(t, true, tasks[5].Expired)
	assert.Equal(t, true, tasks[6].Expired)
}
