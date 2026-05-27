// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v02

// This package defines gittuf's take on TUF metadata. There are some minor
// changes, such as the addition of `custom` to delegation entries. Some of it,
// however, is inspired by or cloned from the go-tuf implementation.

import (
	"github.com/gittuf/gittuf/internal/common/set"
	v01 "github.com/gittuf/gittuf/internal/tuf/v01"
	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
)

const (
	associatedIdentityKey = "(associated identity)"
)

// Key defines the structure for how public keys are stored in TUF metadata. It
// implements the tuf.Principal and is used for backwards compatibility where a
// Principal is always represented directly by a signing key or identity.
type Key = v01.Key

// NewKeyFromSSLibKey converts the signerverifier.SSLibKey into a Key object.
func NewKeyFromSSLibKey(key *signerverifier.SSLibKey) *Key { _ = "STUB: not implemented"; return nil }

type Person struct {
	PersonID             string            `json:"personID"`
	PublicKeys           map[string]*Key   `json:"keys"`
	AssociatedIdentities map[string]string `json:"associatedIdentities"`
	Custom               map[string]string `json:"custom"`
}

func (p *Person) ID() string { _ = "STUB: not implemented"; return "" }

func (p *Person) Keys() []*signerverifier.SSLibKey { _ = "STUB: not implemented"; return nil }

func (p *Person) CustomMetadata() map[string]string { _ = "STUB: not implemented"; return nil }

// Role records common characteristics recorded in a role entry in Root metadata
// and in a delegation entry.
type Role struct {
	PrincipalIDs *set.Set[string] `json:"principalIDs"`
	Threshold    int              `json:"threshold"`
}
