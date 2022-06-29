package commands

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarkTaskComleted(t *testing.T) {
	repo := NewTestRepo()

	output := bytes.NewBuffer([]byte{})
	etc := NewCompleteTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "1"})
	assert.Nil(t, err)
	assert.Equal(t, 5, len(repo.tasks))
	assert.Equal(t, fmt.Sprintf("Task completed succesfully at %s!\n", time.Now().String()[:10]), output.String())
	assert.Equal(t, true, repo.tasks[3].Expired)
}

func TestMarkTaskComletedWithUnexistingId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "10"})
	assert.Error(t, err)
}

func TestMarkTaskComletedWithIdBelowOne(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "-7"})
	assert.Error(t, err)
}

func TestMarkTaskComletedWithoutId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{})
	assert.Error(t, err)
}
