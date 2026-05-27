// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"context"
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
)

// Commit creates a new commit in the repo and sets targetRef's to the commit.
// This function is meant only for gittuf references, and therefore it does not
// mutate repository worktrees.
func (r *Repository) Commit(treeID Hash, targetRef, message string, sign bool) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// CommitUsingSpecificKey creates a new commit in the repository for the
// specified parameters. The commit is signed using the PEM encoded SSH or GPG
// private key. This function is expected for use in tests and gittuf's
// developer mode. In standard workflows, Commit() must be used instead which
// infers the signing key from the user's Git config.
func (r *Repository) CommitUsingSpecificKey(treeID Hash, targetRef, message string, signingKeyPEMBytes []byte) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// commitWithParents creates a new commit in the repo but does not update any
// references. It is only meant to be used for tests, and therefore accepts
// specific parent commit IDs.
func (r *Repository) commitWithParents(t *testing.T, treeID Hash, parentIDs []Hash, message string, sign bool) Hash {
	_ = "STUB: not implemented" //nolint:unparam
	return *new(Hash)
}

// verifyCommitSignature verifies a signature for the specified commit using
// the provided public key.
func (r *Repository) verifyCommitSignature(ctx context.Context, commitID Hash, key *signerverifier.SSLibKey) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCommitMessage returns the commit's message.
func (r *Repository) GetCommitMessage(commitID Hash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetCommitTreeID returns the commit's Git tree ID.
func (r *Repository) GetCommitTreeID(commitID Hash) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// GetCommitParentIDs returns the commit's parent commit IDs.
func (r *Repository) GetCommitParentIDs(commitID Hash) ([]Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KnowsCommit returns true if the `testCommit` is a descendent of the
// `ancestorCommit`. That is, the testCommit _knows_ the ancestorCommit as it
// has a path in the commit graph to the ancestorCommit.
func (r *Repository) KnowsCommit(testCommitID, ancestorCommitID Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// GetCommonAncestor finds the common ancestor commit for the two supplied
// commits.
func (r *Repository) GetCommonAncestor(commitAID, commitBID Hash) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// ensureIsCommit is a helper to check that the ID represents a Git commit
// object.
func (r *Repository) ensureIsCommit(commitID Hash) error { _ = "STUB: not implemented"; return nil }

func getCommitBytesWithoutSignature(commit *object.Commit) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
