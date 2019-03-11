package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kako-jun/gohat/gohat-core"
)

var Version string = "1.0.0"

func parseArgs() (shPath string, err error) {
	var (
		versionFlag bool
	)

	flag.BoolVar(&versionFlag, "version", false, "print version number")

	flag.Parse()
	args := flag.Args()

	if versionFlag {
		fmt.Println(Version)
		os.Exit(0)
	}

	if os.Getuid() != 0 {
		if flag.NArg() < 1 {
			err = errors.New("invalid argument")
			return
		}
	}

	if flag.NArg() >= 1 {
		shPath = args[0]
	}

	return
}

// entry point
func main() {
	shPath, err := parseArgs()
	if err != nil {
		fmt.Println("error:", err)
		fmt.Println("usage:")
		fmt.Println("  gohat [.sh path]")
		return
	}

	gohat.Exec(shPath)
}
