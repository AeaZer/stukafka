package cmd

import (
	"github.com/gokafka/cmd/consumer"
	"github.com/gokafka/cmd/producer"
	"github.com/gokafka/config"
)

// Command interface
type Command interface {
	Name() string
	Run(conf *config.Config)
}

var (
	commands = []Command{
		&producer.Command{},
		&consumer.Command{},
	}
)

// Register register commands
func Register() map[string]Command {
	return register(commands)
}

func register(cmds []Command) map[string]Command {
	commandMap := make(map[string]Command)
	for _, cmd := range cmds {
		commandMap[cmd.Name()] = cmd
	}
	return commandMap
}
