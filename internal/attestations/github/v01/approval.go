// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v01

import (
	authorizationsv01 "github.com/gittuf/gittuf/internal/attestations/authorizations/v01"
	"github.com/gittuf/gittuf/internal/common/set"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	ita "github.com/in-toto/attestation/go/v1"
)

const (
	PullRequestApprovalPredicateType = "https://gittuf.dev/github-pull-request-approval/v0.1"

	digestGitTreeKey = "gitTree"
)

type PullRequestApprovalAttestation struct {
	// Approvers contains the list of currently applicable approvers.
	Approvers *set.Set[string] `json:"approvers"`

	// DismissedApprovers contains the list of approvers who then dismissed
	// their approval.
	DismissedApprovers *set.Set[string] `json:"dismissedApprovers"`

	*authorizationsv01.ReferenceAuthorization
}

func (pra *PullRequestApprovalAttestation) GetApprovers() []string {
	_ = "STUB: not implemented"
	return nil
}

func (pra *PullRequestApprovalAttestation) GetDismissedApprovers() []string {
	_ = "STUB: not implemented"
	return nil
}

// NewPullRequestApprovalAttestation creates a new GitHub pull request approval
// attestation for the provided information. The attestation is embedded in an
// in-toto "statement" and returned with the appropriate "predicate type" set.
// The `fromTargetID` and `toTargetID` specify the change to `targetRef` that is
// approved on the corresponding GitHub pull request.
func NewPullRequestApprovalAttestation(targetRef, fromRevisionID, targetTreeID string, approvers, dismissedApprovers []string) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidatePullRequestApproval(env *sslibdsse.Envelope, targetRef, fromRevisionID, targetTreeID string) error {
	_ = "STUB: not implemented"
	return nil
}
