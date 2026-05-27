// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

func (p *Persistent) GetAttestationsEntries() []RSLEntryIndex {
	_ = "STUB: not implemented"
	return nil
}

// FindAttestationsEntryNumberForEntry returns the index of the attestations
// entry to use. If the returned index has EntryNumber set to 0, it indicates
// that an applicable entry was not found in the cache.
func (p *Persistent) FindAttestationsEntryNumberForEntry(entryNumber uint64) (RSLEntryIndex, bool) {
	_ = "STUB: not implemented"
	// Set entryNumber as max scanned if it's higher than what's already there
	return *new(RSLEntryIndex), false
}

// this happens when an attestations entry doesn't exist before the
// specified entryNumber. No need to use the fallthrough.

func (p *Persistent) InsertAttestationEntryNumber(entryNumber uint64, entryID gitinterface.Hash) {
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

func (p *Persistent) SetAddedAttestationsBeforeNumber(entryNumber uint64) {
	_ = "STUB: not implemented"
	return
}
