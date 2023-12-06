// The above code defines a Go package named "process" that contains a struct type "Process" with
// methods for running, monitoring, and stopping a process.
package process

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/dobycontech/pmsgo/pkg/util"
)

// The above code defines a struct type named "Process" with fields for name, binary path, timeout in
// seconds, and a process executor.
// @property {string} Name - The Name property is a string that represents the name of the process.
// @property {string} Bin - The `Bin` property represents the binary executable file associated with
// the process. It specifies the path to the executable file that will be executed when the process is
// started.
// @property {int64} TimeoutSec - TimeoutSec is an integer property that represents the timeout
// duration in seconds for the process.
// @property PsEx - The `PsEx` property is of type `util.ProcessEx`. It is likely a custom struct or
// type defined in the `util` package.
type Process struct {
	Name       string `json:"name"`
	Bin        string `json:"bin"`
	TimeoutSec int64  `json:"timoutseconds"`
	PsEx       util.ProcessEx
}

// The function `New` loads a JSON file, unmarshals its content into a `Process` struct, and assigns
// the process execution based on the process name.

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

// The `Run()` function is a method defined on the `Process` struct. It is responsible for executing
// the process associated with the `Process` instance.

func (ps Process) Run() bool {
	ended := make(chan int)
	go func() {
		cmd := exec.Command(ps.Bin)
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			println("error running process: ", err.Error())
			ended <- -1
		}
		println("process", ps.Name, "begin run")
		ended <- 1
	}()
	return <-ended > 0
}

// The `Monit()` function is a method defined on the `Process` struct. It is responsible for monitoring
// the status of the process associated with the `Process` instance.

func (ps Process) Monit() {
	if ps.Name == "nil" {
		println("Monitoring: unknown process")
		os.Exit(2)
	}
	if ps.PsEx.Command != "" {
		println("Monitoring: Command running:", ps.PsEx.Command)
	} else {
		println("Monitoring: No process found with this config")
	}
}

// The `Stop()` function is a method defined on the `Process` struct. It is responsible for stopping
// the process associated with the `Process` instance.

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
