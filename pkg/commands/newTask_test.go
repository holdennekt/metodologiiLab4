package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	repo := NewTestRepo()

	output := bytes.NewBuffer([]byte{})
	ntc := NewNewTaskCommand(output, repo)
	err := ntc.Run([]string{"-title", "test task 6", "-details", "just one more uncompleted task", "-deadline", "2022-10-30"})
	assert.Nil(t, err)
	assert.Equal(t, 6, len(repo.tasks))
	assert.Equal(t, "New task created:\nid: 6, title: test task 6\ndetails: just one more uncompleted task\ndeadline: 30 October 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())

	output = bytes.NewBuffer([]byte{})
	ntc = NewNewTaskCommand(output, repo)
	err = ntc.Run([]string{"-title", "test task 7", "-details", "and one more"})
	assert.Nil(t, err)
	assert.Equal(t, 7, len(repo.tasks))
	assert.Equal(t, "New task created:\nid: 7, title: test task 7\ndetails: and one more\ndeadline: no deadline, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())

	output = bytes.NewBuffer([]byte{})
	ntc = NewNewTaskCommand(output, repo)
	err = ntc.Run([]string{"-title", "test task 8"})
	assert.Nil(t, err)
	assert.Equal(t, 8, len(repo.tasks))
	assert.Equal(t, "New task created:\nid: 8, title: test task 8\ndetails: no details\ndeadline: no deadline, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())
}

func TestNewTaskWithEmptyTitle(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	ntc := NewNewTaskCommand(output, repo)
	err := ntc.Run([]string{"-title", ""})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}

func TestNewTaskWithOnlyFlagsProvided(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	ntc := NewNewTaskCommand(output, repo)
	err := ntc.Run([]string{"-title", "-details", "-deadline"})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}

func TestNewTaskWithOnlyValuesProvided(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	ntc := NewNewTaskCommand(output, repo)
	err := ntc.Run([]string{"new task", "details", "2022-09-30"})
	assert.Error(t, err)
	assert.Equal(t, 5, len(repo.tasks))
}
