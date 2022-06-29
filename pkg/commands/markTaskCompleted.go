package commands

import (
	"errors"
	"flag"
	"fmt"
	"io"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type completeTaskCommand struct {
	fs           flag.FlagSet
	id           int
	output       io.Writer
	dataProvider dataproviders.Repository
}

func NewCompleteTaskCommand(output io.Writer, dp dataproviders.Repository) *completeTaskCommand {
	ctc := &completeTaskCommand{
		fs:           *flag.NewFlagSet("complete", flag.ContinueOnError),
		output:       output,
		dataProvider: dp,
	}
	ctc.fs.IntVar(&ctc.id, "id", 0, "id of task")
	return ctc
}

func (ctc *completeTaskCommand) Name() string {
	return ctc.fs.Name()
}

func (ctc *completeTaskCommand) Run(args []string) error {
	ctc.dataProvider.UpdateState()
	err := ctc.fs.Parse(args)
	if err != nil {
		return err
	}
	if ctc.id < 1 {
		return errors.New("-id flag is mandatory and can't be below 0")
	}
	completedAt, err := ctc.dataProvider.MarkCompleted(ctc.id)
	if err != nil {
		return err
	}
	ctc.output.Write([]byte(fmt.Sprintf("Task completed succesfully at %s!\n", completedAt.String()[:10])))
	return nil
}
