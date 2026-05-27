// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	hookopts "github.com/gittuf/gittuf/experimental/gittuf/options/hooks"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
	lua "github.com/yuin/gopher-lua"
)

type ErrHookExists struct {
	HookType HookType
}

func (e *ErrHookExists) Error() string { _ = "STUB: not implemented"; return "" }

type HookType string

const (
	pushedRefLocalRef   = lua.LString("localRef")
	pushedRefRemoteRef  = lua.LString("remoteRef")
	pushedRefLocalHash  = lua.LString("localHash")
	pushedRefRemoteHash = lua.LString("remoteHash")
)

var (
	ErrNoHooksFoundForPrincipal = errors.New("no hooks found for the specified principal")
)

var HookPrePush = HookType("pre-push")

// InvokeHooksForStage runs the hooks defined in the specified stage for the
// user defined by principalID. Upon successful completion of all hooks for the
// stage for the user, the map of hook names to exit codes is returned.  TODO:
// Add attestations workflow
func (r *Repository) InvokeHooksForStage(ctx context.Context, signer sslibdsse.Signer, stage tuf.HookStage, opts ...hookopts.Option) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the principals from targetsMetadata and attempt to find a match for
// the specified principal to determine which hooks to run.

// Couldn't match the key up to a principal, abort

// Now, read all hooks for the specified stage and find which ones we need
// to run

// Determine what parameters must be supplied based on the hook stage

// At the moment, the only stage that we support that requires parameters is
// the pre-push stage.

// https://git-scm.com/docs/githooks#_pre_push
// For pre-push hooks, we supply:
// * remoteName
// * remoteURL
// * localRef
// * localHash
// * remoteRef
// * remoteHash

// TODO (adityasaky): I wonder if we want to move these into separate
// constructors / helpers as we add more stages...

// Validate that all the required values have been passed in

// This likely means the remote doesn't have the specified ref.
// In this case, provide a zero hash as per original Git
// behavior.

// The best type for supplying information about pushed refs that
// gopher-lua supports without too much work is the LTable. We
// create a table for each ref that is being pushed, with an
// equivalent format in Go: []PushRef{{localRef: <>, remoteRef: <>,
// localHash: <>, remoteHash: <>}}

// UpdateHook updates a git hook in the repository's .git/hooks folder.
// Existing hook files are not overwritten, unless force flag is set.
func (r *Repository) UpdateGitHook(hookType HookType, content []byte, force bool) error {
	_ = "STUB: not implemented"
	// TODO: rely on go-git to find .git folder, once
	// https://github.com/go-git/go-git/issues/977 is available.
	// Note, until then gittuf does not support separate git dir.
	return nil
}

// nolint:gosec

func doesFileExist(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *Repository) executeHook(ctx context.Context, hook tuf.Hook, parameters lua.LTable) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Load the hook contents from the repository
