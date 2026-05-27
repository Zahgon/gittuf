// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

// GetCommitsBetweenRange returns the IDs of the commits that exist between the
// specified new and old commit identifiers.
func (r *Repository) GetCommitsBetweenRange(commitNewID, commitOldID Hash) ([]Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: we should ideally be sorting this in the order of occurrence
// rather than by commit ID. The only reason this is happening is because
// the ordering of commitRange by default is not deterministic. Rather than
// walking through them and identifying the right order, we're sorting by
// commit ID. The intended use case of this function is to get a list of
// commits that are then checked for the changes they introduce. At that
// point, they must be diffed with their parent directly.
