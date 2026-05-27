// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

// This file contains modified code from the lua-sandbox project, available at
// https://github.com/kikito/lua-sandbox/blob/master/sandbox.lua, and licensed
// under the MIT License

package luasandbox

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/internal/luasandbox/options/luasandbox"
	"github.com/gittuf/gittuf/pkg/gitinterface"
	lua "github.com/yuin/gopher-lua"
)

const (
	LuaTimeOut = 100
)

var (
	ErrMismatchedAPINames = errors.New("name of API to be registered does not match API implementation")
)

type LuaEnvironment struct {
	lState        *lua.LState
	contextCancel context.CancelFunc
	repository    *gitinterface.Repository
	allAPIs       []API
}

// NewLuaEnvironment creates a new Lua state with the specified timeout.
func NewLuaEnvironment(ctx context.Context, repository *gitinterface.Repository, opts ...luasandbox.EnvironmentOption) (*LuaEnvironment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a new Lua state

// Load default safe libraries

// Must be first

// Load the modules to the Lua state

// Enable only safe functions

// Set the instruction quota and timeout

// Register the Go functions with the Lua state

// RunScript runs the specified script in the given Lua environment, and returns
// the result of running the script. Parameters are provided as strings.
func (l *LuaEnvironment) RunScript(script string, parameters lua.LTable) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If a table is returned, then this likely means that the hook didn't
// return an exit code. Return a 1 for safety.

func (l *LuaEnvironment) GetAPIs() []API { _ = "STUB: not implemented"; return nil }

func (l *LuaEnvironment) Cleanup() {
	_ = "STUB: not implemented"

	// enableOnlySafeFunctions disables all functions that are deemed to be unsafe.
	return
}

func (l *LuaEnvironment) enableOnlySafeFunctions() {
	_ = "STUB: not implemented"
	// -- List of unsafe packages/functions:
	// -- * string.rep: can be used to allocate millions of bytes in 1 operation
	// -- * {set|get}metatable: can be used to modify the metatable of global objects (strings, integers)
	// -- * collectgarbage: can affect performance of other systems
	// -- * dofile: can access the server filesystem
	// -- * _G: It has access to everything. It can be mocked to other things though.
	// -- * load{file|string}: All unsafe because they can grant acces to global env
	// -- * raw{get|set|equal}: Potentially unsafe
	// -- * module|require|module: Can modify the host settings
	// -- * string.dump: Can display confidential server info (implementation of functions)
	// -- * math.randomseed: Can affect the host system
	// -- * io.*, os.*: Most stuff there is unsafe
	// -- * debug.*: Unsafe, see https://www.lua.org/pil/23.html
	// -- * package.*: Allows arbitrary module loading, see https://www.lua.org/manual/5.3/manual.html#pdf-package
	return
}

// Disable all unsafe functions

// Load protected modules with only safe functions

// protectModule protects the specified module from being modified by setting a
// protected metatable with __newindex and __metatable fields.
func (l *LuaEnvironment) protectModule(tbl *lua.LTable, moduleName string) {
	_ = "STUB: not implemented"
	return
}

// setTimeOut sets the timeout for the Lua state.
func (l *LuaEnvironment) setTimeOut(ctx context.Context, timeOut int) {
	_ = "STUB: not implemented"
	return
}

// registerAPIFunctions makes the sandbox's standard APIs available.
func (l *LuaEnvironment) registerAPIFunctions() error {
	_ = "STUB: not implemented"
	// Set global variables for the Lua state
	return nil
}
