package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/holdennekt/metodologiiLab4/pkg/commands"
)

type Command interface {
	Run([]string) error
	Name() string
}

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	config, err := ParseConfig("config.json")
	if err != nil {
		log.Fatal("failed to parse config:", err)
	}
	fmt.Println(config)
	cmds := []Command{
		commands.NewShowTasksCommand(),
		// rest of commands
	}
	for _, cmd := range cmds {
		if os.Args[1] == cmd.Name() {
			err := cmd.Run(os.Args[2:])
			if err != nil {
				log.Fatalf("failed to run command %v: %v", cmd.Name(), err)
			}
		}
	}
}
