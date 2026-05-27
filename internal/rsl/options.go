// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package rsl

import "github.com/gittuf/gittuf/pkg/gitinterface"

type GetLatestReferenceUpdaterEntryOptions struct {
	Reference string

	BeforeEntryID     gitinterface.Hash
	BeforeEntryNumber uint64

	UntilEntryID     gitinterface.Hash
	UntilEntryNumber uint64

	Unskipped bool

	NonGittuf bool

	IsReferenceEntry                bool
	IsPropagationEntryForRepository string
}

type GetLatestReferenceUpdaterEntryOption func(*GetLatestReferenceUpdaterEntryOptions)

// ForReference indicates that the reference entry returned must be for a
// specific Git reference.
func ForReference(reference string) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// BeforeEntryID searches for the matching reference entry before the specified
// entry ID. It cannot be used in combination with BeforeEntryNumber.
// BeforeEntryID is exclusive: the returned entry cannot be the reference entry
// that matches the specified ID.
func BeforeEntryID(entryID gitinterface.Hash) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// BeforeEntryNumber searches for the matching reference entry before the
// specified entry number. It cannot be used in combination with BeforeEntryID.
// BeforeEntryNumber is exclusive: the returned entry cannot be the reference
// entry that matches the specified number.
func BeforeEntryNumber(number uint64) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// UntilEntryID terminates the search for the desired reference entry when an
// entry with the specified ID is encountered. It cannot be used in combination
// with UntilEntryNumber. UntilEntryID is inclusive: the returned entry can be
// the entry that matches the specified ID.
func UntilEntryID(entryID gitinterface.Hash) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// UntilEntryNumber terminates the search for the desired reference entry when
// an entry with the specified number is encountered. It cannot be used in
// combination with UntilEntryID. UntilEntryNumber is inclusive: the returned
// entry can be the entry that matches the specified number.
func UntilEntryNumber(number uint64) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// IsUnskipped ensures that the returned reference entry has not been skipped by
// a subsequent annotation entry.
func IsUnskipped() GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// ForNonGittufReference ensures that the returned reference entry is not for a
// gittuf-specific reference.
func ForNonGittufReference() GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// IsReferenceEntry ensures that the returned entry is a reference entry
// specifically, rather than any entry type that matches the ReferenceUpdater
// interface.
func IsReferenceEntry() GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}

// IsPropagationEntryForRepository ensures that the returned entry is a
// propagation entry for the specified upstream repository.
func IsPropagationEntryForRepository(repositoryLocation string) GetLatestReferenceUpdaterEntryOption {
	_ = "STUB: not implemented"
	return *new(GetLatestReferenceUpdaterEntryOption)
}
