package process

import (
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Process struct {
	Name       string `json:"name"`
	Bin        string `json:"bin"`
	TimeoutSec int64  `json:"timoutseconds"`
	PsEx ProcessEx
}
type ProcessEx struct {
	User    string
	Pid     int64
	Cpu     float64
	Mem     float64
	Vsz     int64
	Rss     int64
	TTY     string
	Stat    string
	Start   string
	Time    string
	Command string
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
func (ps *Process) Run() {
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
func (ps *Process) Monit() {
	if isWin() {
		println("Monitorización: TODO win some day")
		return
	}
	ps.PsEx = getProcessByName(ps.Name)
	if ps.Name == "nil" {
		println("Monitorización: unknown process")
		return
	}
	if ps.PsEx.Command != "" {
		println("Monitorización: comando:", ps.PsEx.Command)
	} else {
		println("Monitorización: Ningún programa en ejecución")
	}
}
func getProcessByName(name string) ProcessEx{
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		println("error checking monitor: ", err)
	}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(line, name) {
			pex := ProcessEx{}
			i := 0
			for _, it := range strings.Split(line, " ") {
				if it == "" {
					continue
				}
				switch i {
				case 0:
					pex.User = it
				case 1:
					pex.Pid, _ = strconv.ParseInt(it, 0, 64)

				case 2:
					pex.Cpu, _ = strconv.ParseFloat(it, 32)
				case 3:
					pex.Mem, _ = strconv.ParseFloat(it, 32)
				case 4:
					pex.Vsz, _ = strconv.ParseInt(it, 0, 64)
				case 5:
					pex.Rss, _ = strconv.ParseInt(it, 0, 64)
				case 6:
					pex.TTY = it
				case 7:
					pex.Stat = it
				case 8:
					pex.Start = it
				case 9:
					pex.Time = it
				case 10:
					pex.Command = it
				}
				i++
			}
			return pex
		}
	}
	return ProcessEx{}
}

func (ps *Process) Stop() {
	ps.PsEx = getProcessByName(ps.Name)
	if ps.PsEx.Pid > 0 {
		out, err := exec.Command("kill", strconv.Itoa(int(ps.PsEx.Pid))).Output()
		if err != nil {
			println("error:", err.Error())
		}
		println("Deteniendo: programa finalizado:", ps.PsEx.Pid, string(out))
	} else {
		println("Deteniendo: Ningún programa pendiente finalizar")
	}
}
func isWin() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "win")
}