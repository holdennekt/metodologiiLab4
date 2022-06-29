package commands

import (
	"bytes"
	"errors"
	"flag"
	"io"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type showTasksCommand struct {
	fs             flag.FlagSet
	all, todo, exp bool
	output         io.Writer
	dataProvider   dataproviders.Repository
}

func NewShowTasksCommand(output io.Writer, dp dataproviders.Repository) *showTasksCommand {
	stc := &showTasksCommand{
		fs:           *flag.NewFlagSet("show", flag.ContinueOnError),
		output:       output,
		dataProvider: dp,
	}
	stc.fs.BoolVar(&stc.all, "all", false, "all tasks")
	stc.fs.BoolVar(&stc.todo, "todo", false, "uncompleted tasks")
	stc.fs.BoolVar(&stc.exp, "exp", false, "expired tasks")
	return stc
}

func (stc *showTasksCommand) Name() string {
	return stc.fs.Name()
}

func (stc *showTasksCommand) showTasks(tasks []*dataproviders.Task) {
	out := make([][]byte, 0)
	for _, task := range tasks {
		out = append(out, []byte(getTaskStr(task)))
	}
	sep := []byte("-------------------------------------------------------\n")
	stc.output.Write(bytes.Join(out, sep))
}

func (stc *showTasksCommand) Run(args []string) error {
	stc.dataProvider.UpdateState()
	err := stc.fs.Parse(args)
	if err != nil {
		return err
	}
	flags := []bool{stc.all, stc.todo, stc.exp}
	timesTrue := 0
	for _, v := range flags {
		if v {
			timesTrue++
		}
	}
	if timesTrue > 1 {
		return errors.New("only 1 of the folowing flags allowed at a time: -all, -todo, -exp")
	}
	ifAny := stc.all || stc.todo || stc.exp
	if !ifAny || stc.all {
		tasks, err := stc.dataProvider.ListAllTasks()
		if err != nil {
			return err
		}
		stc.showTasks(tasks)
	}
	if stc.todo {
		tasks, err := stc.dataProvider.ListActiveTasks()
		if err != nil {
			return err
		}
		stc.showTasks(tasks)
	}
	if stc.exp {
		tasks, err := stc.dataProvider.ListExpiredTasks()
		if err != nil {
			return err
		}
		stc.showTasks(tasks)
	}
	return nil
}
