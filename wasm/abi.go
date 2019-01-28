// Copyright (C) 2018-2019 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package wasm

// +build wasm/js

import (
	"syscall/js"

	"honnef.co/go/js/dom"
)

var domWindow = dom.GetWindow()

var jsGlobs = js.Global()
