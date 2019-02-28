package main

import (
	"errors"
	"flag"
	"fmt"

	"./gohat"
)

func parseArgs() ( command string, commandArgs []string, err error) {
	flag.Parse()
	if flag.NArg() < 1 {
		err = errors.New("invalid argument")
		return
	}

	args := flag.Args()
	command = args[0]
	commandArgs = args[1:]
	return
}

// entry point
func main() {
	 command, args, err := parseArgs()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("usage:")
		fmt.Println("  gohat [command]")
		fmt.Println("  gohat [command] [command options]")
		return
	}

	gohat.Exec(command, args)
}
