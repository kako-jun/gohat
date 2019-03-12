// Package gohat is ***
package gohat

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// Gohat is ***
type Gohat struct{}

func (gohat Gohat) exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func (gohat Gohat) isSUIDEnabled(gohatPath string) (enabled bool) {
	enabled = false

	f, err := os.Open(gohatPath)
	if err != nil {
		log.Fatal(err)
	}

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	mode := info.Mode().String()
	if strings.Contains(mode, "u") {
		enabled = true
	}

	return
}

func (gohat Gohat) setSUID(gohatPath string) (errReturn error) {
	if err := os.Chown(gohatPath, 0, 0); err != nil {
		fmt.Println("info:", "at the first launch, gohat needs sudo.")
		fmt.Println("")

		if strings.Contains(gohatPath, " ") {
			fmt.Println("  e.g. sudo \"" + gohatPath + "\"")
		} else {
			fmt.Println("  e.g. sudo " + gohatPath)
		}

		fmt.Println("")
		errReturn = err
		return
	}

	if runtime.GOOS == "darwin" {
		fmt.Println("chown root:wheel " + gohatPath)
	} else {
		fmt.Println("chown root:root " + gohatPath)
	}

	// if err := os.Chmod("gohat", os.FileMode ModeSetuid); err != nil {
	if err := exec.Command("chmod", "u+s", gohatPath).Run(); err != nil {
		errReturn = err
		return
	}

	fmt.Println("chmod u+s " + gohatPath)
	return
}

func (gohat Gohat) execScriptAsSu(command string, scriptPath string) (err error) {
	cmd := exec.Command(command, scriptPath)

	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	outScanner := bufio.NewScanner(stdout)
	for outScanner.Scan() {
		fmt.Fprintln(os.Stdout, outScanner.Text())
	}

	errScanner := bufio.NewScanner(stderr)
	for errScanner.Scan() {
		fmt.Fprintln(os.Stderr, errScanner.Text())
	}

	cmd.Wait()
	return
}

func (gohat Gohat) start(scriptPath string) (err error) {
	gohatPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	enabled := gohat.isSUIDEnabled(gohatPath)
	if enabled {
		if !gohat.exists(scriptPath) {
			err = errors.New(scriptPath + " not found.")
			return
		}

		command := ""
		ext := filepath.Ext(scriptPath)
		switch ext {
		case ".rb":
			command = "ruby"
		case ".py":
			command = "python"
		default:
			command = "sh"
		}

		err = gohat.execScriptAsSu(command, scriptPath)
	} else {
		err = gohat.setSUID(gohatPath)
	}

	return
}

// Exec is ***
func Exec(scriptPath string) (errReturn error) {
	gohat := new(Gohat)
	if err := gohat.start(scriptPath); err != nil {
		fmt.Println("error:", err)
		errReturn = errors.New("error")
		return
	}

	return
}
