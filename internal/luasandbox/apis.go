// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package luasandbox

import (
	lua "github.com/yuin/gopher-lua"
)

// API presents the interface for any API made available within the sandbox.
type API interface {
	GetName() string
	GetSignature() string
	GetHelp() string
	GetExamples() []string
}

// LuaAPI implements the API interface. This is used when the API is implemented
// as a Lua function.
type LuaAPI struct {
	Name           string
	Signature      string
	Help           string
	Examples       []string
	Implementation string
}

func (l *LuaAPI) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *LuaAPI) GetSignature() string { _ = "STUB: not implemented"; return "" }

func (l *LuaAPI) GetHelp() string { _ = "STUB: not implemented"; return "" }

func (l *LuaAPI) GetExamples() []string {
	_ = "STUB: not implemented"

	// GoAPI implements the API interface. This is used when the API is implemented
	// in Go.
	return nil
}

type GoAPI struct {
	Name           string
	Signature      string
	Help           string
	Examples       []string
	Implementation lua.LGFunction
}

func (g *GoAPI) GetName() string { _ = "STUB: not implemented"; return "" }

func (g *GoAPI) GetSignature() string { _ = "STUB: not implemented"; return "" }

func (g *GoAPI) GetHelp() string { _ = "STUB: not implemented"; return "" }

func (g *GoAPI) GetExamples() []string { _ = "STUB: not implemented"; return nil }

func (l *LuaEnvironment) apiMatchRegex() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiStrSplit() API { _ = "STUB: not implemented"; return *new(API) }

// TODO: check if examples are right

func (l *LuaEnvironment) apiGitReadBlob() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetObjectSize() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetTagTarget() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetReference() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetAbsoluteReference() API {
	_ = "STUB: not implemented"
	return *new(API)
}

func (l *LuaEnvironment) apiGitGetSymbolicReferenceTarget() API {
	_ = "STUB: not implemented"
	return *new(API)
}

func (l *LuaEnvironment) apiGitGetCommitMessage() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetFilePathsChangedByCommit() API {
	_ = "STUB: not implemented"
	return *new(API)
}

func (l *LuaEnvironment) apiGitGetRemoteURL() API { _ = "STUB: not implemented"; return *new(API) }

func (l *LuaEnvironment) apiGitGetStagedFilePaths() API {
	_ = "STUB: not implemented"
	return *new(API)
}
