package commands

import (
	"errors"
	"flag"
	"io"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type editTaskCommand struct {
	fs                       flag.FlagSet
	id                       int
	title, details, deadline string
	output                   io.Writer
	dataProvider             dataproviders.Repository
}

func NewEditTaskCommand(output io.Writer, dp dataproviders.Repository) *editTaskCommand {
	etc := &editTaskCommand{
		fs:           *flag.NewFlagSet("edit", flag.ContinueOnError),
		output:       output,
		dataProvider: dp,
	}
	etc.fs.IntVar(&etc.id, "id", 0, "id of task")
	etc.fs.StringVar(&etc.title, "title", "", "title of task")
	etc.fs.StringVar(&etc.details, "details", "", "details of task")
	etc.fs.StringVar(&etc.deadline, "deadline", "", "deadline of task")
	return etc
}

func (etc *editTaskCommand) Name() string {
	return etc.fs.Name()
}

func (etc *editTaskCommand) Run(args []string) error {
	etc.dataProvider.UpdateState()
	err := etc.fs.Parse(args)
	if err != nil {
		return err
	}
	if etc.id < 1 {
		return errors.New("-id flag is mandatory and can't be below 0")
	}
	task, err := etc.dataProvider.EditTask(etc.id, etc.title, etc.details, etc.deadline)
	if err != nil {
		return err
	}
	taskStr := "Edited task:\n" + getTaskStr(task)
	etc.output.Write([]byte(taskStr))
	return nil
}
