package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kako-jun/gohat/gohat-core"
)

func parseArgs() (shPath string, commandArgs []string, err error) {
	flag.Parse()
	args := flag.Args()

	if os.Getuid() != 0 {
		if flag.NArg() < 1 {
			err = errors.New("invalid argument")
			return
		}
	}

	if flag.NArg() >= 1 {
		shPath = args[0]
	}

	if flag.NArg() >= 2 {
		commandArgs = args[1:]
	}

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
