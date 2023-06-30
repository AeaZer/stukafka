package main

import (
	"fmt"
	"github.com/gokafka/cmd"
	"log"
	"os"

	"github.com/gokafka/config"
)

const defaultMicroName = "producer"

func init() {
	c, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	config.Conf = c
}

func main() {
	commandMap := cmd.Register()
	if len(os.Args) == 1 {
		os.Args = append(os.Args, defaultMicroName)
	}
	command, ok := commandMap[os.Args[1]]
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "Unknown run mode: %s\n", os.Args[1])
		os.Exit(1)
	}
	command.Run(config.Conf)
}
