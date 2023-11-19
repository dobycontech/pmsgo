package process

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/dobycontech/pmsgo/internal/util"
)

type Process struct {
	Name       string `json:"name"`
	Bin        string `json:"bin"`
	TimeoutSec int64  `json:"timoutseconds"`
	PsEx       util.ProcessEx
}

func New(jsonPath string) (Process, error) {
	jsoncontent, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatal("error loading json file, check exists:", jsonPath, err.Error())
	}
	var ps Process
	err = json.Unmarshal(jsoncontent, &ps)
	if err != nil {
		log.Fatal("error loading params from json", err.Error())
	}
	ps.PsEx = util.GetProcessByName(ps.Name)
	return ps, nil
}
func (ps Process) Run() {
	ended := make(chan int)
	go func() {
		cmd := exec.Command(ps.Bin)
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			println("error running process: ", err.Error())
			os.Exit(-1)
		}
		println("process", ps.Name, "begin run")
		ended <- 1
	}()
	<-ended
}
func (ps Process) Monit() {
	if ps.Name == "nil" {
		println("Monitoring: unknown process")
		os.Exit(2)
	}
	if ps.PsEx.Command != "" {
		println("Monitoring: Command running:", ps.PsEx.Command)
		os.Exit(1)
	} else {
		println("Monitoring: No process found with this config")
		os.Exit(1)
	}
}

func (ps Process) Stop() {
	if ps.PsEx.Pid > 0 {
		out, err := exec.Command("kill", strconv.Itoa(int(ps.PsEx.Pid))).Output()
		if err != nil {
			log.Fatal("error killing:", ps.PsEx.Pid, err.Error())

		}
		println("Process stopped succesfully: ", ps.PsEx.Pid, string(out))
	} else {
		println("No process to stop found")
	}
}
