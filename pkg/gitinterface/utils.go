// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"testing"
)

// ResetDueToError reverses a change applied to a ref to the specified target
// ID. It is used to ensure a gittuf operation is atomic: if a gittuf operation
// fails, any changes made to the repository in refs/gittuf can be rolled back.
// Worktrees are not updated.
func (r *Repository) ResetDueToError(cause error, refName string, commitID Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoteRef(refName, remoteName string) string { _ = "STUB: not implemented"; return "" }

// refs/heads/<path> -> refs/remotes/<remote>/<path>

// refs/tags/<path> -> refs/tags/<path>

// refs/<path> -> refs/remotes/<remote>/<path>

// RestoreWorktree is a test helper to fix the worktree in tests where we need
// to operate in a checked out copy of the repository. This is primarily needed
// for support with older Git versions.
func (r *Repository) RestoreWorktree(t *testing.T) { _ = "STUB: not implemented"; return }

// TODO: this doesn't support detached git dir

//nolint:errcheck

// IsNiceGitVersion determines whether the version of git is "nice". Certain Git
// subcommands that gittuf uses were added in newer versions than some common
// client versions. Instead of using a workaround for all clients, we determine
// if we can use the newer features or instead need to use workarounds.
func isNiceGitVersion() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func testNameToRefName(testName string) string { _ = "STUB: not implemented"; return "" }
