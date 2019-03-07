// gohat is ***
package gohat

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	// if err := os.Chmod("gohat", os.FileMode ModeSetuid); err != nil {
	if err := exec.Command("chmod", "u+s", gohatPath).Run(); err != nil {
		errReturn = err
		return
	}

	return
}

func (gohat Gohat) execScriptAsSu(shPath string, args ...string) (err error) {
	cmd := exec.Command(shPath, args...)

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

func (gohat Gohat) Start(shPath string, args []string) (err error) {
	gohatPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	enabled := gohat.isSUIDEnabled(gohatPath)
	if enabled {
		if !gohat.exists(shPath) {
			err = errors.New(shPath + " not found.")
			return
		}

		err = gohat.execScriptAsSu("sh", shPath)
	} else {
		err = gohat.setSUID(gohatPath)
	}

	return
}

// Exec is ***
func Exec(shPath string, args []string) (errReturn error) {
	gohat := new(Gohat)
	if err := gohat.Start(shPath, args); err != nil {
		fmt.Println("error:", err)
		errReturn = errors.New("error")
		return
	}

	return
}
