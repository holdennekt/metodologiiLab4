package commands

import (
	"errors"
	"flag"
	"io"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type newTaskCommand struct {
	fs                       flag.FlagSet
	title, details, deadline string
	output                   io.Writer
	dataProvider             dataproviders.Repository
}

func NewNewTaskCommand(output io.Writer, dp dataproviders.Repository) *newTaskCommand {
	ntc := &newTaskCommand{
		fs:           *flag.NewFlagSet("new", flag.ContinueOnError),
		output:       output,
		dataProvider: dp,
	}
	ntc.fs.StringVar(&ntc.title, "title", "", "title of task")
	ntc.fs.StringVar(&ntc.details, "details", "", "details of task")
	ntc.fs.StringVar(&ntc.deadline, "deadline", "", "deadline of task")
	return ntc
}

func (ntc *newTaskCommand) Name() string {
	return ntc.fs.Name()
}

func (ntc *newTaskCommand) Run(args []string) error {
	ntc.dataProvider.UpdateState()
	err := ntc.fs.Parse(args)
	if err != nil {
		return err
	}
	if ntc.title == "" {
		return errors.New("-title flag is mandatory and can't be \"\"")
	}
	task, err := ntc.dataProvider.NewTask(ntc.title, ntc.details, ntc.deadline)
	if err != nil {
		return err
	}
	taskStr := "New task created:\n" + getTaskStr(task)
	ntc.output.Write([]byte(taskStr))
	return nil
}
