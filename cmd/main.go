package main

import (
	"log"
	"os"

	"github.com/holdennekt/metodologiiLab4/pkg/commands"
)

type Command interface {
	Run([]string) error
	Name() string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	cmds := []Command{
		commands.NewShowTasksCommand(),
		// rest of commands
	}
	for _, cmd := range cmds {
		if os.Args[1] == cmd.Name() {
			cmd.Run(os.Args[2:])
		}
	}
}
