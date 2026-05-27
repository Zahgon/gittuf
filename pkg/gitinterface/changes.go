// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

// GetFilePathsChangedByCommit returns the paths changed by the commit relative
// to its parent commit. If the commit is a merge commit, i.e., it has more than
// one parent, check if the commit is the same as at least one of its parents.
// If there is a matching parent, we return no changes. If there is no matching
// parent commit, we return the changes between the commit and each of its parents.
func (r *Repository) GetFilePathsChangedByCommit(commitID Hash) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if tree matches last commit
