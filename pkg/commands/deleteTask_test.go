package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteTaskSuccessfull(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	dtc := NewDeleteTaskCommand(output, repo)
	err := dtc.Run([]string{"-id", "2"})
	assert.Nil(t, err)
	assert.Equal(t, 4, len(repo.tasks))
	assert.Equal(t, "Task deleted:\nid: 2, title: test task 2\ndetails: this is an expired task\ndeadline: 22 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n", output.String())
	assert.Equal(t, true, repo.tasks[2].Expired)
}

func TestDeleteTaskWithoutId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	dtc := NewDeleteTaskCommand(output, repo)
	err := dtc.Run([]string{})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}

func TestDeleteTaskWithIdBelowOne(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	dtc := NewDeleteTaskCommand(output, repo)
	err := dtc.Run([]string{"-id", "-1"})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}

func TestDeleteTaskWithUnexistingId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	dtc := NewDeleteTaskCommand(output, repo)
	err := dtc.Run([]string{"-id", "10"})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}
