package commands

import (
	"errors"
	"flag"
	"io"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type deleteTaskCommand struct {
	fs           flag.FlagSet
	id           int
	output       io.Writer
	dataProvider dataproviders.Repository
}

func NewDeleteTaskCommand(output io.Writer, dp dataproviders.Repository) *deleteTaskCommand {
	dtc := &deleteTaskCommand{
		fs:           *flag.NewFlagSet("delete", flag.ContinueOnError),
		output:       output,
		dataProvider: dp,
	}
	dtc.fs.IntVar(&dtc.id, "id", 0, "id of task")
	return dtc
}

func (dtc *deleteTaskCommand) Name() string {
	return dtc.fs.Name()
}

func (dtc *deleteTaskCommand) Run(args []string) error {
	dtc.dataProvider.UpdateState()
	err := dtc.fs.Parse(args)
	if err != nil {
		return err
	}
	if dtc.id < 1 {
		return errors.New("-id flag is mandatory and can't be below 0")
	}
	task, err := dtc.dataProvider.DeleteTask(dtc.id)
	if err != nil {
		return err
	}
	taskStr := "Task deleted:\n" + getTaskStr(task)
	dtc.output.Write([]byte(taskStr))
	return nil
}
