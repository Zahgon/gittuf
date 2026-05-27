// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package attestations

import (
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/pkg/gitinterface"
	ita "github.com/in-toto/attestation/go/v1"
)

// NewReferenceAuthorizationForCommit creates a new reference authorization for
// the provided information. The authorization is embedded in an in-toto
// "statement" and returned with the appropriate "predicate type" set. The
// `fromID` and `toID` specify the change to `targetRef` that is to be
// authorized by invoking this function. Since this is for a commit, the `toID`
// is expected to be a Git tree ID.
func NewReferenceAuthorizationForCommit(targetRef, fromID, toID string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewReferenceAuthorizationForTag creates a new reference authorization for the
// provided information. The authorization is embedded in an in-toto "statement"
// and returned with the appropriate "predicate type" set. The `fromID` and
// `toID` specify the change to `targetRef` that is to be authorized by invoking
// this function. Since this is for a tag, the `toID` is expected to be a Git
// commit ID.
func NewReferenceAuthorizationForTag(targetRef, fromID, toID string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetReferenceAuthorization writes the new reference authorization attestation
// to the object store and tracks it in the current attestations state.
func (a *Attestations) SetReferenceAuthorization(repo *gitinterface.Repository, env *sslibdsse.Envelope, refName, fromID, toID string) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveReferenceAuthorization removes a set reference authorization
// attestation entirely. The object, however, isn't removed from the object
// store as prior states may still need it.
func (a *Attestations) RemoveReferenceAuthorization(refName, fromID, toID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetReferenceAuthorizationFor returns the requested reference authorization
// attestation (with its signatures).
func (a *Attestations) GetReferenceAuthorizationFor(repo *gitinterface.Repository, refName, fromID, toID string) (*sslibdsse.Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReferenceAuthorizationPath constructs the expected path on-disk for the
// reference authorization attestation.
func ReferenceAuthorizationPath(refName, fromID, toID string) string {
	_ = "STUB: not implemented"
	return ""
}
