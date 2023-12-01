package util

import (
	"testing"
	"os/exec"
	"os"
	"time"
)

func Test_GetByName(t *testing.T){
	pid := -1
	go func(){
		cmd := exec.Command("/bin/sleep", "5 second")
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			t.Error(err)
		}
		pid = cmd.Process.Pid
	}()
	for range time.Tick(1*time.Second){
		if pid < -10 {
			t.Log("Cant create process")
		}
		if pid > 0 {
			break
		}
		pid--
	}
	proc := GetProcessByName("sleep 1")
	if int(proc.Pid) != pid{
		t.Log("process not exists")
	}
	if proc.Command == "" {
		t.Error("Process not found")
	}
}