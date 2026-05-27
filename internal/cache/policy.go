// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

func (p *Persistent) GetPolicyEntries() []RSLEntryIndex { _ = "STUB: not implemented"; return nil }

func (p *Persistent) HasPolicyEntryNumber(entryNumber uint64) (gitinterface.Hash, bool) {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash), false
}

// Unlike Find... we're actively checking if a policy number has been
// inserted into the cache before, so we return the ID from that index
// exactly

func (p *Persistent) FindPolicyEntryNumberForEntry(entryNumber uint64) RSLEntryIndex {
	_ = "STUB: not implemented"
	return *new(RSLEntryIndex)
}

// this is a special case

// The entry number given to us is the first entry which happens
// to be the start of verification as well
// This can happen for full verification

// this happens when a policy entry doesn't exist before the specified
// entryNumber

// When !has, index is point of insertion, but we want the applicable
// entry which is index-1

func (p *Persistent) FindPolicyEntriesInRange(firstNumber, lastNumber uint64) ([]RSLEntryIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check if custom error makes sense

// When !has, index is point of insertion, but we want the applicable
// entry which is index-1

// When has, lastIndex is an entry we want to return, so we increment
// lastIndex to ensure the corresponding entry is included in the return

func (p *Persistent) InsertPolicyEntryNumber(entryNumber uint64, entryID gitinterface.Hash) {
	_ = "STUB: not implemented"
	return

	// For now, we don't have a way to track non-numbered entries
	// We likely never want to track non-numbered entries in this
	// cache as this is very dependent on numbering
}

// TODO: check this is for the right ref?

// No entries yet, just add the current entry

// Current entry clearly belongs at the very end

// We don't check the converse where the current entry is less than the
// first entry because we're inserting as entries are encountered
// chronologically. Worst case, binary search fallthrough below will still
// handle it

// We could assume that if we've seen an entry with a number greater
// than this, we should have seen this one too, but for now...
