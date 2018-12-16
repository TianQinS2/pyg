[![GoDoc](https://godoc.org/github.com/glycerine/gopy?status.svg)](https://godoc.org/github.com/glycerine/gopy)

This package presents an idomatic Go interface to the CPython C API described at
http://docs.python.org/3/c-api/index.html

Embedding Python is fully supported, with the ability to initialise the
interpreter, enable threading support, manipulate the GIL and call Python API
functions to manipulate Python objects.

The package targets Python 3.7.1 at present.
