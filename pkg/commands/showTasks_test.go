package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowAllTasks(t *testing.T) {
	repo := NewTestRepo()

	output := bytes.NewBuffer([]byte{})
	stc := NewShowTasksCommand(output, repo)
	err := stc.Run([]string{})
	assert.Nil(t, err)
	assert.Equal(t, "id: 1, title: test task 1\ndetails: this is an active task\ndeadline: 1 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 2, title: test task 2\ndetails: this is an expired task\ndeadline: 22 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 3, title: test task 3\ndetails: this is a completed task\ndeadline: no deadline, expired: false\ncompleted: true, completedAt: 28 June 2022\n-------------------------------------------------------\nid: 4, title: test task 4\ndetails: this task has been expired at 2022-06-28, but not marked yet\ndeadline: 28 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 5, title: test task 5\ndetails: one more active task\ndeadline: 1 August 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())

	output = bytes.NewBuffer([]byte{})
	stc = NewShowTasksCommand(output, repo)
	err = stc.Run([]string{"-all"})
	assert.Nil(t, err)
	assert.Equal(t, "id: 1, title: test task 1\ndetails: this is an active task\ndeadline: 1 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 2, title: test task 2\ndetails: this is an expired task\ndeadline: 22 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 3, title: test task 3\ndetails: this is a completed task\ndeadline: no deadline, expired: false\ncompleted: true, completedAt: 28 June 2022\n-------------------------------------------------------\nid: 4, title: test task 4\ndetails: this task has been expired at 2022-06-28, but not marked yet\ndeadline: 28 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 5, title: test task 5\ndetails: one more active task\ndeadline: 1 August 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())
}

func TestShowTasksWrongInput(t *testing.T) {
	repo := NewTestRepo()

	output := bytes.NewBuffer([]byte{})
	stc := NewShowTasksCommand(output, repo)
	err := stc.Run([]string{"-all", "-todo", "-exp"})
	assert.Error(t, err)

	output = bytes.NewBuffer([]byte{})
	stc = NewShowTasksCommand(output, repo)
	err = stc.Run([]string{"-all", "-todo"})
	assert.Error(t, err)

	output = bytes.NewBuffer([]byte{})
	stc = NewShowTasksCommand(output, repo)
	err = stc.Run([]string{"-all", "-exp"})
	assert.Error(t, err)

	output = bytes.NewBuffer([]byte{})
	stc = NewShowTasksCommand(output, repo)
	err = stc.Run([]string{"-todo", "-exp"})
	assert.Error(t, err)
}

func TestShowActiveTasks(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	stc := NewShowTasksCommand(output, repo)
	err := stc.Run([]string{"-todo"})
	assert.Nil(t, err)
	assert.Equal(t, "id: 5, title: test task 5\ndetails: one more active task\ndeadline: 1 August 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 1, title: test task 1\ndetails: this is an active task\ndeadline: 1 September 2022, expired: false\ncompleted: false, completedAt: wasn't completed\n", output.String())
}

func TestShowExpiredTasks(t *testing.T) {
	repo := NewTestRepo()
	output := bytes.NewBuffer([]byte{})
	stc := NewShowTasksCommand(output, repo)
	err := stc.Run([]string{"-exp"})
	assert.Nil(t, err)
	assert.Equal(t, "id: 2, title: test task 2\ndetails: this is an expired task\ndeadline: 22 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n-------------------------------------------------------\nid: 4, title: test task 4\ndetails: this task has been expired at 2022-06-28, but not marked yet\ndeadline: 28 June 2022, expired: true\ncompleted: false, completedAt: wasn't completed\n", output.String())
}
