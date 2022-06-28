package commands

import (
	"flag"
	"fmt"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

type showTasksCommand struct {
	fs             flag.FlagSet
	all, todo, exp bool
	dataProvider   *dataproviders.DataProvider
}

func showTask(t *dataproviders.Task) {
	var details string
	if t.Details.Valid {
		details = t.Details.String
	} else {
		details = "no details"
	}
	var deadline string
	if t.Deadline.Valid {
		year := t.Deadline.Time.Year()
		month := t.Deadline.Time.Month().String()
		day := t.Deadline.Time.Day()
		deadline = fmt.Sprintf("%v %v %v", day, month, year)
	} else {
		deadline = "no deadline"
	}
	var completedAt string
	if t.CompletedAt.Valid {
		year := t.CompletedAt.Time.Year()
		month := t.CompletedAt.Time.Month().String()
		day := t.CompletedAt.Time.Day()
		completedAt = fmt.Sprintf("%v %v %v", day, month, year)
	} else {
		completedAt = "wasn't completed"
	}
	fmt.Printf("id: %v, title: %v\ndetails: %v\ndeadline: %v, expired: %v\ncompleted: %v, completedAt: %v\n", t.Id, t.Title, details, deadline, t.Expired, t.Completed, completedAt)
	fmt.Println("----------------------------------------------------------------")
}

func showTasks(tasks []*dataproviders.Task) {
	for _, task := range tasks {
		showTask(task)
	}
}

func (sc *showTasksCommand) Run(args []string) error {
	err := sc.fs.Parse(args)
	if err != nil {
		return err
	}
	ifAny := sc.all || sc.todo || sc.exp
	if !ifAny || sc.all {
		tasks, err := sc.dataProvider.ListAllTasks()
		if err != nil {
			return err
		}
		showTasks(tasks)
	}
	if sc.todo {
		tasks, err := sc.dataProvider.ListActiveTasks()
		if err != nil {
			return err
		}
		showTasks(tasks)
	}
	if sc.exp {
		tasks, err := sc.dataProvider.ListExpiredTasks()
		if err != nil {
			return err
		}
		showTasks(tasks)
	}
	return nil
}

func (sc *showTasksCommand) Name() string {
	return sc.fs.Name()
}

func NewShowTasksCommand(dp *dataproviders.DataProvider) *showTasksCommand {
	sc := &showTasksCommand{
		fs:           *flag.NewFlagSet("show", flag.ContinueOnError),
		dataProvider: dp,
	}
	sc.fs.BoolVar(&sc.all, "all", false, "all tasks")
	sc.fs.BoolVar(&sc.todo, "todo", false, "uncompleted tasks")
	sc.fs.BoolVar(&sc.exp, "exp", false, "expired tasks")
	return sc
}
