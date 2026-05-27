// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"github.com/gittuf/gittuf/internal/cache"
	"github.com/gittuf/gittuf/internal/rsl"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

// searcher defines the interface for finding policy and attestation entries in
// the RSL.
type searcher interface {
	FindFirstPolicyEntry() (rsl.ReferenceUpdaterEntry, error)
	FindLatestPolicyEntry() (rsl.ReferenceUpdaterEntry, error)
	FindPolicyEntryFor(rsl.Entry) (rsl.ReferenceUpdaterEntry, error)
	FindPolicyEntriesInRange(rsl.Entry, rsl.Entry) ([]rsl.ReferenceUpdaterEntry, error)
	FindAttestationsEntryFor(rsl.Entry) (rsl.ReferenceUpdaterEntry, error)
	FindLatestAttestationsEntry() (rsl.ReferenceUpdaterEntry, error)
}

func newSearcher(repo *gitinterface.Repository) searcher {
	_ = "STUB: not implemented"
	return *new(searcher)
}

// regularSearcher implements the searcher interface. It walks back the RSL from
// to identify the requested policy or attestation entries.
type regularSearcher struct {
	repo *gitinterface.Repository
}

// FindFirstPolicyEntry identifies the very first policy entry in the RSL.
func (r *regularSearcher) FindFirstPolicyEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// we don't have a policy entry yet

// FindLatestPolicyEntry returns the latest policy entry in the RSL.
func (r *regularSearcher) FindLatestPolicyEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// we don't have a policy entry

// FindPolicyEntryFor identifies the latest policy entry for the specified
// entry.
func (r *regularSearcher) FindPolicyEntryFor(entry rsl.Entry) (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	// If the requested entry itself is for the policy ref, return as is
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// Any other err must be returned

// FindPolicyEntriesInRange returns all policy RSL entries in the specified
// range. firstEntry and lastEntry are included if they are for the policy ref.
func (r *regularSearcher) FindPolicyEntriesInRange(firstEntry, lastEntry rsl.Entry) ([]rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAttestationsEntryFor identifies the latest attestations entry for the
// specified entry.
func (r *regularSearcher) FindAttestationsEntryFor(entry rsl.Entry) (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	// If the requested entry itself is for the attestations ref, return as is
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// Attestations may not be used yet, they're not
// compulsory

// FindLatestAttestationsEntry returns the latest RSL entry for the attestations
// reference.
func (r *regularSearcher) FindLatestAttestationsEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// we don't have an attestations entry

func newRegularSearcher(repo *gitinterface.Repository) *regularSearcher {
	_ = "STUB: not implemented"
	return nil
}

// cacheSearcher implements the searcher interface. It checks the persistent
// cache for results before falling back to the regular searcher if the
// persistent cache yields no results.
type cacheSearcher struct {
	repo            *gitinterface.Repository
	persistentCache *cache.Persistent
	searcher        *regularSearcher
}

// FindFirstPolicyEntry identifies the very first policy entry in the RSL.
func (c *cacheSearcher) FindFirstPolicyEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

func (c *cacheSearcher) FindLatestPolicyEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// FindPolicyEntryFor identifies the latest policy entry for the specified
// entry.
func (c *cacheSearcher) FindPolicyEntryFor(entry rsl.Entry) (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// no number is set

// FindPolicyEntriesInRange returns all policy RSL entries in the specified
// range. firstEntry and lastEntry are included if they are for the policy ref.
func (c *cacheSearcher) FindPolicyEntriesInRange(firstEntry, lastEntry rsl.Entry) ([]rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first or last entry doesn't have a number

// FindAttestationsEntryFor identifies the latest attestations entry for the
// specified entry.
func (c *cacheSearcher) FindAttestationsEntryFor(entry rsl.Entry) (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

// no number is set

func (c *cacheSearcher) FindLatestAttestationsEntry() (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}

func newCacheSearcher(repo *gitinterface.Repository, persistentCache *cache.Persistent) *cacheSearcher {
	_ = "STUB: not implemented"
	return nil
}

func loadRSLReferenceUpdaterEntry(repo *gitinterface.Repository, entryID gitinterface.Hash) (rsl.ReferenceUpdaterEntry, error) {
	_ = "STUB: not implemented"
	return *new(rsl.ReferenceUpdaterEntry), nil
}
