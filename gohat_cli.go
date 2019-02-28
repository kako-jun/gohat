package main

import (
	"errors"
	"flag"
	"fmt"

	"./gohat"
)

func parseArgs() (shPath string, commandArgs []string, err error) {
	flag.Parse()
	if flag.NArg() < 1 {
		err = errors.New("invalid argument")
		return
	}

	args := flag.Args()
	shPath = args[0]
	commandArgs = args[1:]
	return
}

// entry point
func main() {
	shPath, args, err := parseArgs()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("usage:")
		fmt.Println("  gohat [.sh path]")
		return
	}

	gohat.Exec(shPath, args)
}
