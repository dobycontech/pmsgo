package util

import (
	"os"
	"os/exec"
	//"strconv"
	"testing"
)

func Test_GetByName(t *testing.T){
	pid := make(chan int)
	go func(){
		cmd := exec.Command("/bin/sleep", "5 second")
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
	if int(proc.Pid) != processId{
		t.Log("process not exists")
	}
	if proc.Command == "" {
		t.Error("Process not found")
	}
	
	/*
	go func(){
		err := exec.Command("kill", strconv.Itoa(pid))
		if err != nil {
			t.Log("fail to remove pid: ", pid, " with error: ", err)
		}
	}()
	<-time.After(2*time.Second)
	*/
}