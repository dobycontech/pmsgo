package util

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)
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

func GetProcessByName(name string) ProcessEx{
	if isWin(){
		log.Fatal("Windows: Next Iteration")
	}
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		println("error checking monitor: ", err)
	}
	pex := ProcessEx{}
	for _, line := range strings.Split(string(out), "\n") {
		if strings.Contains(strings.ToLower(line), strings.ToLower(name)) {
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
			// TODO: Con motivo de pruebas y dado que sleep existe con el sistema quito los sleep de sistema
			if pex.Pid > 100 && strings.Contains(strings.ToLower(pex.Command), strings.ToLower(name)){
				return pex
			}
			pex = ProcessEx{}
		}
	}
	return pex
}
func isWin() bool {
	return strings.Contains(strings.ToLower(runtime.GOOS), "win")
}