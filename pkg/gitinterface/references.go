// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"errors"
)

const (
	RefPrefix       = "refs/"
	BranchRefPrefix = "refs/heads/"
	TagRefPrefix    = "refs/tags/"
	RemoteRefPrefix = "refs/remotes/"
)

var (
	ErrReferenceNotFound = errors.New("requested Git reference not found")
)

// GetReference returns the tip of the specified Git reference.
func (r *Repository) GetReference(refName string) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// SetReference sets the specified reference to the provided Git ID.
func (r *Repository) SetReference(refName string, gitID Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteReference deletes the specified Git reference.
func (r *Repository) DeleteReference(refName string) error { _ = "STUB: not implemented"; return nil }

// CheckAndSetReference sets the specified reference to the provided Git ID if
// the reference is currently set to `oldGitID`.
func (r *Repository) CheckAndSetReference(refName string, newGitID, oldGitID Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSymbolicReferenceTarget returns the name of the Git reference the provided
// symbolic Git reference is pointing to.
func (r *Repository) GetSymbolicReferenceTarget(refName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetSymbolicReference sets the specified symbolic reference to the specified
// target reference.
func (r *Repository) SetSymbolicReference(symRefName, targetRefName string) error {
	_ = "STUB: not implemented"
	return nil
}

// AbsoluteReference returns the fully qualified reference path for the provided
// Git ref.
// Source: https://git-scm.com/docs/gitrevisions#Documentation/gitrevisions.txt-emltrefnamegtemegemmasterememheadsmasterememrefsheadsmasterem
func (r *Repository) AbsoluteReference(target string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// not symbolic ref

// symbolic ref such as .git/HEAD

// We may have a ref that isn't available locally but is still ref-prefixed.

// If target is a full ref already and it's stored in the GIT_DIR/refs
// directory, we don't reach this point. Below, we handle cases where the
// ref may be packed.

// Check if custom reference

// Check if tag

// Check if branch

// Check if remote tracker ref

// RefSpec creates a Git refspec for the specified ref.  For more information on
// the Git refspec, please consult:
// https://git-scm.com/book/en/v2/Git-Internals-The-Refspec.
func (r *Repository) RefSpec(refName, remoteName string, fastForwardOnly bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: check if this is correct, AFAICT tags aren't tracked in the
// remotes namespace.

// local is always refPath, destination depends on remoteName

// CustomReferenceName returns the full reference name in the form
// `refs/<customName>`.
func CustomReferenceName(customName string) string { _ = "STUB: not implemented"; return "" }

// TagReferenceName returns the full reference name for the specified tag in the
// form `refs/tags/<tagName>`.
func TagReferenceName(tagName string) string { _ = "STUB: not implemented"; return "" }

// BranchReferenceName returns the full reference name for the specified branch
// in the form `refs/heads/<branchName>`.
func BranchReferenceName(branchName string) string { _ = "STUB: not implemented"; return "" }

// RemoteReferenceName returns the full reference name in the form
// `refs/remotes/<name>`.
func RemoteReferenceName(name string) string { _ = "STUB: not implemented"; return "" }
