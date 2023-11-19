package main

import (
	"github.com/dobycontech/pmsgo/process"
	"os"
	"strings"
)

func main() {
	var ps process.Process
	var err error
	for _, arg := range os.Args {
		if strings.Contains(arg, "--file") {
			value := strings.Split(arg, "=")[1]
			if _, err := os.Stat(value); err != nil {
				println("file not exists: ", value)
				os.Exit(-1)
			}
			ps, err = process.New(value)
			if err != nil {
				println("error creating process:", err)
				os.Exit(-1)
			}
			err = ps.Run()
			if err != nil {
				println("error running process:", err)
				os.Exit(-1)
			}
		}
			if strings.Contains(arg, "--monit") {
				ps.Monit()
			}
			if strings.Contains(arg, "--stop") {
				ps.Stop()
			}
	}
}
