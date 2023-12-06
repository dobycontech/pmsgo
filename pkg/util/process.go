// The `GetProcessByName` function retrieves information about a process by its name.
package util

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// The type ProcessEx represents information about a process in a system.
// @property {string} User - The user who owns the process.
// @property {int64} Pid - The Pid property represents the process ID, which is a unique identifier
// assigned to each running process in an operating system.
// @property {float64} Cpu - The "Cpu" property represents the CPU usage of the process as a float64
// value. It indicates the percentage of CPU time that the process has used.
// @property {float64} Mem - The "Mem" property in the "ProcessEx" struct represents the amount of
// memory used by the process. It is of type float64, which means it can store decimal values. The unit
// of measurement for memory usage is typically in bytes.
// @property {int64} Vsz - Vsz stands for "Virtual Size" and represents the total amount of virtual
// memory used by the process. It includes both the memory that is currently in use and the memory that
// has been allocated but not yet used.
// @property {int64} Rss - Rss stands for Resident Set Size. It represents the amount of physical
// memory (in kilobytes) that a process is currently using.
// @property {string} TTY - The TTY property in the ProcessEx struct represents the terminal device
// associated with the process. It indicates the terminal or console where the process is running.
// @property {string} Stat - The "Stat" property in the ProcessEx struct represents the current state
// of the process. It provides information about the process's status, such as whether it is running,
// sleeping, or stopped.
// @property {string} Start - The Start property represents the start time of the process. It indicates
// the time when the process was started.
// @property {string} Time - The "Time" property in the ProcessEx struct represents the time at which
// the process was started or last updated. It is a string that typically follows a specific format,
// such as "HH:MM:SS" or "YYYY-MM-DD HH:MM:SS".
// @property {string} Command - The Command property represents the name of the process or command that
// is being executed.
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

// The function `GetProcessByName` retrieves information about a process by its name.
func GetProcessByName(name string) ProcessEx {
	if IsWin() {
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
			if pex.Pid > 100 && strings.Contains(strings.ToLower(pex.Command), strings.ToLower(name)) {
				return pex
			}
			pex = ProcessEx{}
		}
	}
	return pex
}

// The IsWin function checks if the operating system is Windows.
func IsWin() bool {
	os := strings.ToLower(runtime.GOOS)
	return strings.Contains(os, "windows")
}

// The IsLinux function checks if the operating system is Linux.
func IsLinux() bool {
	os := strings.ToLower(runtime.GOOS)
	return strings.Contains(os, "linux")
}

// The IsMac function checks if the operating system is macOS.
func IsMac() bool {
	os := strings.ToLower(runtime.GOOS)
	return strings.Contains(os, "darwin")
}
