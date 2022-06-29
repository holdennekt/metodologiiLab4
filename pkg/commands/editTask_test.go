package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEditTask(t *testing.T) {
	repo := NewTestRepo()

	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "1", "-title", "test task 1 edited", "-details", "this is an active edited task", "-deadline", "2022-09-21"})
	assert.Nil(t, err)
	assert.Equal(t, 5, len(repo.tasks))
	assert.Equal(t, "Edited task:\nid: 1, title: test task 1 edited\ndetails: this is an active edited task\ndeadline: 21 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())
	assert.Equal(t, true, repo.tasks[3].Expired)

	output = bytes.NewBuffer([]byte{})
	etc = NewEditTaskCommand(output, repo)
	err = etc.Run([]string{"-id", "1", "-title", "test task 1 edited^2", "-details", "this is an active edited^2 task"})
	assert.Nil(t, err)
	assert.Equal(t, "Edited task:\nid: 1, title: test task 1 edited^2\ndetails: this is an active edited^2 task\ndeadline: 21 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())

	output = bytes.NewBuffer([]byte{})
	etc = NewEditTaskCommand(output, repo)
	err = etc.Run([]string{"-id", "1", "-title", "test task 1 edited^3"})
	assert.Nil(t, err)
	assert.Equal(t, "Edited task:\nid: 1, title: test task 1 edited^3\ndetails: this is an active edited^2 task\ndeadline: 21 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())

	output = bytes.NewBuffer([]byte{})
	etc = NewEditTaskCommand(output, repo)
	err = etc.Run([]string{"-id", "1"})
	assert.Nil(t, err)
	assert.Equal(t, "Edited task:\nid: 1, title: test task 1 edited^3\ndetails: this is an active edited^2 task\ndeadline: 21 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())
}

func TestEditTaskWithUnexistingDeadline(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "1", "-deadline", "2022-14-33"})
	assert.Error(t, err)
}

func TestEditTaskWithUnexistingId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "10"})
	assert.Error(t, err)
}

func TestEditTaskWithIdBelowOne(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{"-id", "-7"})
	assert.Error(t, err)
}

func TestEditTaskWithoutId(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	etc := NewEditTaskCommand(output, repo)
	err := etc.Run([]string{})
	assert.Error(t, err)
}
