package process

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Process struct {
	Name       string `json:"name"`
	Bin        string `json:"bin"`
	TimeoutSec int64  `json:"timoutseconds"`
}

func New(jsonPath string) (Process, error) {
	jsoncontent, err := os.ReadFile(jsonPath)
	if err != nil {
		return Process{}, err
	}
	var ps Process
	err = json.Unmarshal(jsoncontent, &ps)
	if err != nil {
		return Process{}, err
	}
	return ps, nil
}
func (ps *Process) Run() error {
	cmd := exec.Command(ps.Bin)
	cmd.Stdout = os.Stdout
	return cmd.Start()
}
func (ps *Process) Monit() {
	if ps.Name == "nil" {
		println("unknown process")
	}
	if isWin() {
		println("todo win some day")
	} else {
		out, err := exec.Command("ps", "aux").Output()
		if err != nil {
			println("error checking monitor: ", err)
		}
		for _, line := range strings.Split(string(out), "\n") {
			if strings.Contains(line, ps.Name) {
				println("out:", line)
			}
		}
	}
}
func (ps *Process) Stop() error {
	return nil
}
func isWin() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "win")
}
