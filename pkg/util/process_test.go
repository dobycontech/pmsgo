package util

import (
	"os"
	"os/exec"

	//"strconv"
	"testing"
)

func Test_GetByName(t *testing.T) {
	pid := make(chan int)
	go func() {
		secs := "5 second"
		if IsMac() {
			secs = "5"
		}
		cmd := exec.Command("/bin/sleep", secs)
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			t.Error(err)
		}
		pid <- cmd.Process.Pid
	}()
	processId := <-pid
	close(pid)
	proc := GetProcessByName("sleep")
	if int(proc.Pid) != processId {
		t.Log("process not exists")
	}
	if proc.Command == "" {
		t.Error("Process not found")
	}
}
