// The main function checks command line arguments and performs actions based on the arguments, such as
// creating a new process, monitoring a process, or stopping a process.
package main

import (
	"os"
	"strings"

	process "github.com/dobycontech/pmsgo/service"
)

// The main function checks command line arguments and performs actions based on the arguments, such as
// creating a new process, monitoring a process, or stopping a process.
func main() {
	var ps process.Process
	var err error
	canRun := false
	for _, arg := range os.Args {
		switch {
		case strings.Contains(arg, "--file"):
			canRun = true
			value := strings.Split(arg, "=")[1]
			if _, err := os.Stat(value); err != nil {
				println("file not exists: ", value)
				os.Exit(-1)
			}
			ps, err = process.New(value)
			if err != nil {
				println("error creating process:", err.Error())
				os.Exit(-1)
			}
		case strings.Contains(arg, "--monit"):
			canRun = false
			ps.Monit()
		case strings.Contains(arg, "--stop"):
			canRun = false
			ps.Stop()
			ps.Monit()
		}
	}
	if canRun {
		ps.Run()
	}
}
