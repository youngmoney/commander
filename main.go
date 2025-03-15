package main

import (
	"flag"
	"fmt"
	"os"
)

func commandCommand(name string, args []string, commands *[]Command) error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	match := Match(name, cwd, commands)
	if match == nil {
		return nil
	}

	return ExecuteCommandInteractive(match.Command, args)
}

func listCommands(commands *[]Command) {
	for _, c := range *commands {
		fmt.Println(c.Name)
	}
}

func main() {
	configFilename := flag.String("config", os.Getenv("COMMANDER_CONFIG"), "config file (yaml), or set COMMANDER_CONFIG")
	flag.Parse()

	config := ReadConfig(*configFilename)

	if flag.Arg(0) != "command" {
		fmt.Fprintln(os.Stderr, "command is the only supported command")
		os.Exit(1)
	}

	fs := flag.NewFlagSet("command", flag.ExitOnError)
	fs.Parse(flag.Args()[1:])

	if fs.NArg() == 0 {
		listCommands(&config.Commander.Commands)
		os.Exit(1)
	}

	commandError := commandCommand(fs.Arg(0), fs.Args()[1:], &config.Commander.Commands)
	ExitIfNonZero(commandError)
}
