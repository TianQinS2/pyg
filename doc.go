// Copyright 2011, 2018 Julian Phillips and the gopy contributors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
`pyg` is forked and derived from https://github.com/qur/gopy; it has been adapted to target python 3.7.1.

This package presents an idomatic Go interface to the CPython C API described at
http://docs.python.org/3/c-api/index.html

Embedding Python is fully supported, with the ability to initialise the
interpreter, enable threading support, manipulate the GIL and call Python API
functions to manipulate Python objects.

Installation is a simple matter of re-pointing the `python-3.7m.pc` symlink to point to the pkg-config .pc file in your python installation.

*/
package pyg
