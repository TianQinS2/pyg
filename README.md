[![GoDoc](https://godoc.org/github.com/glycerine/pyg?status.svg)](https://godoc.org/github.com/glycerine/pyg)

`pyg` is forked and derived from https://github.com/qur/gopy; it has been adapted to target python 3.7.1.

This package presents an idomatic Go interface to the CPython C API described at
http://docs.python.org/3/c-api/index.html

Embedding Python is fully supported, with the ability to initialise the
interpreter, enable threading support, manipulate the GIL and call Python API
functions to manipulate Python objects.
