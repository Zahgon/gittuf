// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v01

import (
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	ita "github.com/in-toto/attestation/go/v1"
)

const (
	PredicateType = "https://gittuf.dev/reference-authorization/v0.1"

	digestGitTreeKey  = "gitTree"
	targetRefKey      = "targetRef"
	fromRevisionIDKey = "fromRevisionID"
	targetTreeIDKey   = "targetTreeID"
)

// ReferenceAuthorization is a lightweight record of a detached authorization in
// a gittuf repository. It is meant to be used as a "predicate" in an in-toto
// attestation.
type ReferenceAuthorization struct {
	TargetRef      string `json:"targetRef"`
	FromRevisionID string `json:"fromRevisionID"`
	TargetTreeID   string `json:"targetTreeID"`
}

func (r *ReferenceAuthorization) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *ReferenceAuthorization) GetFromID() string { _ = "STUB: not implemented"; return "" }

func (r *ReferenceAuthorization) GetTargetID() string { _ = "STUB: not implemented"; return "" }

// NewReferenceAuthorization creates a new reference authorization for the
// provided information. The authorization is embedded in an in-toto "statement"
// and returned with the appropriate "predicate type" set. The `fromRevisionID`
// and `targetTreeID` specify the change to `targetRef` that is to be authorized
// by invoking this function.
func NewReferenceAuthorization(targetRef, fromRevisionID, targetTreeID string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate checks that the returned envelope contains the expected in-toto
// attestation and predicate contents.
func Validate(env *sslibdsse.Envelope, targetRef, fromRevisionID, targetTreeID string) error {
	_ = "STUB: not implemented"
	return nil
}
