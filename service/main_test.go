package process

import (
	"testing"
)

func Test_New(t *testing.T) {
	// need json path, take from /test
	jsonpath := "../test/ps.json"
	ps, err := New(jsonpath)
	if err != nil {
		t.Error("error creating process", err.Error())
	}
	if ps.Name == "" {
		t.Error("error creating process, bad json config")
	}
}
func Test_Run(t *testing.T) {
	jsonpath := "../test/ps.json"
	ps, err := New(jsonpath)
	if err != nil {
		t.Error("error creating process", err.Error())
	}
	if !ps.Run() {
		t.Error("fail to run the process")
	}
}
func Test_Monit(t *testing.T) {
	jsonpath := "../test/ps.json"
	ps, err := New(jsonpath)
	if err != nil {
		t.Error("error creating process", err.Error())
	}
	ps.Monit()
	if ps.PsEx.Command == "" || ps.PsEx.Pid == 0 {
		t.Error("Fail to run process")
	}
}
func Test_Stop(t *testing.T) {
	jsonpath := "../test/ps.json"
	ps, err := New(jsonpath)
	if err != nil {
		t.Error("error creating process", err.Error())
	}
	ps.Monit()
	if ps.PsEx.Command == "" || ps.PsEx.Pid == 0 {
		t.Error("Fail to run process")
	}
	ps.Stop()
}
