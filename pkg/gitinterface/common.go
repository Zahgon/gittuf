// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"testing"
	"time"

	"github.com/jonboulle/clockwork"
)

const (
	testName  = "Jane Doe"
	testEmail = "jane.doe@example.com"
)

var (
	testClock = clockwork.NewFakeClockAt(time.Date(1995, time.October, 26, 9, 0, 0, 0, time.UTC))
)

// CreateTestGitRepository creates a Git repository in the specified directory.
// This is meant to be used by tests across gittuf packages. This helper also
// sets up an ED25519 signing key that can be used to create reproducible
// commits.
func CreateTestGitRepository(t *testing.T, dir string, bare bool) *Repository {
	_ = "STUB: not implemented"
	return nil
}

// Set up author / committer identity

// Set up signing via SSH key

func setupRepository(t *testing.T, dir string, bare bool) *Repository {
	_ = "STUB: not implemented"
	return nil
}

func setupSigningKeys(t *testing.T, dir string) { _ = "STUB: not implemented"; return }
