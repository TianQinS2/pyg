package py

import (
	"fmt"
	"os"
	"testing"
)

func TestAddingImportPathAndRun(t *testing.T) {

	goPath := os.Getenv("GOPATH")
	pathAdd := goPath + "/src/github.com/glycerine/gp/install/env3/lib/python3.7/site-packages"

	_, err := os.Stat(pathAdd)
	if err != nil {
		t.Skip()
		return // skip test if custom env3 not available.
	}

	Initialize()
	defer Finalize()
	main, err := NewDict()
	if err != nil {
		t.Fatal(err)
	}
	g, err := GetBuiltins()
	if err != nil {
		t.Fatal(err)
	}

	// g = '&py.Dict{AbstractObject:py.AbstractObject{}, o:(*py._Ctype_struct___154)(0x5281900)}'
	vv("g = '%#v'", g)

	err = main.SetItemString("__builtins__", g)
	if err != nil {
		t.Fatal(err)
	}
	_, err = RunString("a = 'hello world!'", FileInput, main, nil)
	if err != nil {
		t.Fatal(err)
	}
	a, err := main.GetItemString("a")
	if err != nil {
		t.Fatal(err)
	}
	b, ok := a.(*Unicode)
	if !ok || b.String() != "hello world!" {
		t.Error(b, err)
	}

	_, err = RunString(fmt.Sprintf("import sys; sys.path.insert(0, '%s'); sp=sys.path;", pathAdd), FileInput, main, nil)
	panicOn(err)

	sysPath, err := main.GetItemString("sp")
	panicOn(err)

	vv("sysPath = '%s'", sysPath)

	_, err = RunString("import pandas;", FileInput, main, nil)
	panicOn(err)

}

func TestRunString(t *testing.T) {

	Initialize()
	defer Finalize()
	main, err := NewDict()
	if err != nil {
		t.Fatal(err)
	}
	g, err := GetBuiltins()
	if err != nil {
		t.Fatal(err)
	}

	err = main.SetItemString("__builtins__", g)
	if err != nil {
		t.Fatal(err)
	}
	_, err = RunString("a = 'hello world!'", FileInput, main, nil)
	if err != nil {
		t.Fatal(err)
	}
	a, err := main.GetItemString("a")
	if err != nil {
		t.Fatal(err)
	}
	b, ok := a.(*Unicode)
	if !ok || b.String() != "hello world!" {
		t.Error(b, err)
	}
}
