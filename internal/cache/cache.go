// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"errors"

	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const (
	Ref = "refs/local/gittuf/persistent-cache"

	persistentTreeEntryName = "persistentCache"

	policyRef = "refs/gittuf/policy" // this is copied from internal/policy to avoid an import cycle
)

var (
	ErrNoPersistentCache = errors.New("persistent cache not found")
	ErrEntryNotNumbered  = errors.New("one or more entries are not numbered")
)

type Persistent struct {
	// PolicyEntries is a list of index values for entries pertaining to the
	// policy ref. The list is ordered by each entry's Number.
	PolicyEntries []RSLEntryIndex `json:"policyEntries"`

	// AttestationEntries is a list of index values for entries pertaining to
	// the attestations ref. The list is ordered by each entry's Number.
	AttestationEntries []RSLEntryIndex `json:"attestationEntries"`

	// AddedAttestationsBeforeNumber tracks the number up to which
	// attestations have been searched for and added to
	// attestationsEntryNumbers. We need to track this for attestations in
	// particular because attestations are optional in gittuf repositories,
	// meaning attestationsEntryNumbers may be empty which would trigger a
	// full search.
	AddedAttestationsBeforeNumber uint64 `json:"addedAttestationsBeforeNumber"`

	// LastVerifiedEntryForRef is a map that indicates the last verified RSL
	// entry for a ref.
	LastVerifiedEntryForRef map[string]RSLEntryIndex `json:"lastVerifiedEntryForRef"`
}

func (p *Persistent) Commit(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// nothing to do

//nolint:errcheck

// no change in cache contents, noop

// PopulatePersistentCache scans the repository's RSL and generates a persistent
// local-only cache of policy and attestation entries. This makes subsequent
// verifications faster.
func PopulatePersistentCache(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadPersistentCache loads the persistent cache from the tip of the local ref.
// If an instance has already been loaded and a pointer has been stored in
// memory, that instance is returned.
func LoadPersistentCache(repo *gitinterface.Repository) (*Persistent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Persistent cache doesn't exist

// Persistent cache doesn't seem to exist? This maybe warrants
// an error but we may have more than one file here in future?

// DeletePersistentCache deletes the local persistent cache ref.
func DeletePersistentCache(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// RSLEntryIndex is essentially a tuple that maps RSL entry IDs to numbers. This
// may be expanded in future to include more information as needed.
type RSLEntryIndex struct {
	EntryID     string `json:"entryID"`
	EntryNumber uint64 `json:"entryNumber"`
}

func (r *RSLEntryIndex) GetEntryID() gitinterface.Hash {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash)
}

// TODO: error?

func (r *RSLEntryIndex) GetEntryNumber() uint64 { _ = "STUB: not implemented"; return 0 }

func binarySearch(a, b RSLEntryIndex) int { _ = "STUB: not implemented"; return 0 }

// Exact match

// Precedes

// Succeeds
