package py

import (
	"testing"
)

func TestProgramName(t *testing.T) {
	const want = "myProgramName"
	SetProgramName(want)
	name := GetProgramName()
	if name != want {
		t.Fatalf("got=%q. want=%q", name, want)
	}
}
func TestPythonHome(t *testing.T) {
	const want = "/gopath/src/github.com/gp/install/env3/lib/python3.7/site-packages"
	SetPythonHome(want)
	got := GetPythonHome()
	if got != want {
		t.Fatalf("got=%q. want=%q", got, want)
	}
}
