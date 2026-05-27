// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v02

import (
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	ita "github.com/in-toto/attestation/go/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	PredicateType = "https://gittuf.dev/reference-authorization/v0.2"

	digestGitTreeKey   = "gitTree"
	digestGitCommitKey = "gitCommit"
	targetRefKey       = "targetRef"
	fromIDKey          = "fromID"
	targetIDKey        = "targetID"
)

// ReferenceAuthorization is a lightweight record of a detached authorization in
// a gittuf repository. It is meant to be used as a "predicate" in an in-toto
// attestation.
type ReferenceAuthorization struct {
	TargetRef string `json:"targetRef"`
	FromID    string `json:"fromID"`
	TargetID  string `json:"targetID"`
}

func (r *ReferenceAuthorization) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *ReferenceAuthorization) GetFromID() string { _ = "STUB: not implemented"; return "" }

func (r *ReferenceAuthorization) GetTargetID() string {
	_ = "STUB: not implemented"

	// NewReferenceAuthorizationForCommit creates a new reference authorization for
	// the provided information. The authorization is embedded in an in-toto
	// "statement" and returned with the appropriate "predicate type" set. The
	// `fromID` and `targetID` specify the change to `targetRef` that is to be
	// authorized by invoking this function. The targetID is expected to be the Git
	// tree ID of the resultant commit.
	return ""
}

func NewReferenceAuthorizationForCommit(targetRef, fromID, targetID string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewReferenceAuthorizationForTag creates a new reference authorization for the
// provided information. The authorization is embedded in an in-toto "statement"
// and returned with the appropriate "predicate type" set. The `fromID` and
// `targetID` specify the change to `targetRef` that is to be authorized by
// invoking this function. The targetID is expected to be the ID of the commit
// the tag will point to.
func NewReferenceAuthorizationForTag(targetRef, fromID, targetID string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate checks that the returned envelope contains the expected in-toto
// attestation and predicate contents.
func Validate(env *sslibdsse.Envelope, targetRef, fromID, targetID string) error {
	_ = "STUB: not implemented"
	return nil
}

func newReferenceAuthorizationStruct(targetRef, fromID, targetID string) (*structpb.Struct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
