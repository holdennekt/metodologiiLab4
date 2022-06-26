package main

import (
	"os"

	"github.com/holdennekt/metodologiiLab4/pkg/commands"
)

type Command interface {
	Run([]string) error
	Name() string
}

func main() {
	cmds := []Command{
		commands.NewShowTasksCommand(),
	}
	for _, cmd := range cmds {
		if os.Args[1] == cmd.Name() {
			cmd.Run(os.Args[2:])
		}
	}
}
