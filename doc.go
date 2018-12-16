// Copyright 2011, 2018 Julian Phillips and the gopy contributors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package py (github.com/glycerine/gopy) provides access to the CPython C API.  This
package presents an idomatic Go interface to the CPython C API described at
http://docs.python.org/c-api/index.html

Instead of simply exposing the C API as-is, this package uses interfaces,
embedding, type assertions and methods to try and present the functionality of
the Python API in a manner that feels more natural in Go.

Embedding Python

Embedding Python is fully supported, with the ability to initialise the
interpreter, enable threading support, manipulate the GIL and call Python API
functions to manipulate Python objects.

In addition to providing the ability to use the API to call into Python calling
from Python back into Go is also supported.  New types can be implemented in Go
and exposed into Python.

The package targets Python 3.7.1 at present.

*/
package pyg
