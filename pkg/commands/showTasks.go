package commands

import (
	"flag"
	"fmt"
)

type showTasksCommand struct {
	fs             flag.FlagSet
	all, todo, exp bool
}

func (sc *showTasksCommand) Run(args []string) error {
	err := sc.fs.Parse(args)
	if err != nil {
		return err
	}
	ifAny := sc.all || sc.todo || sc.exp
	if !ifAny || sc.all {
		// TODO: implement "show all tasks" scenario
		fmt.Println("all")
	}
	if sc.todo {
		// TODO: implement "show active tasks" scenario
		fmt.Println("uncompleted")
	}
	if sc.exp {
		// TODO: implement "show expired tasks" scenario
		fmt.Println("expire")
	}
	return nil
}

func (sc *showTasksCommand) Name() string {
	return sc.fs.Name()
}

func NewShowTasksCommand() *showTasksCommand {
	sc := &showTasksCommand{
		fs: *flag.NewFlagSet("show", flag.ContinueOnError),
	}
	sc.fs.BoolVar(&sc.all, "all", false, "all tasks")
	sc.fs.BoolVar(&sc.todo, "todo", false, "uncompleted tasks")
	sc.fs.BoolVar(&sc.exp, "exp", false, "expired tasks")
	return sc
}
