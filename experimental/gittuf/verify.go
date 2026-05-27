// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	verifyopts "github.com/gittuf/gittuf/experimental/gittuf/options/verify"
	verifymergeableopts "github.com/gittuf/gittuf/experimental/gittuf/options/verifymergeable"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

// ErrRefStateDoesNotMatchRSL is returned when a Git reference being verified
// does not have the same tip as identified in the latest RSL entry for the
// reference. This can happen for a number of reasons such as incorrectly
// modifying reference state away from what's recorded in the RSL to not
// creating an RSL entry for some new changes. Depending on the context, one
// resolution is to update the reference state to match the RSL entry, while
// another is to create a new RSL entry for the current state.
var ErrRefStateDoesNotMatchRSL = errors.New("current state of Git reference does not match latest RSL entry")

func (r *Repository) VerifyRef(ctx context.Context, refName string, opts ...verifyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Track localRefName to check the expected tip as we may override refName

// remote ref name is different
// We must consider RSL entries that have refNameOverride rather than
// refName

// To verify the tip, we _must_ use the localRefName

func (r *Repository) VerifyRefFromEntry(ctx context.Context, refName, entryID string, opts ...verifyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Track localRefName to check the expected tip as we may override refName

// remote ref name is different
// We must consider RSL entries that have refNameOverride rather than
// refName

// To verify the tip, we _must_ use the localRefName

// VerifyMergeable checks if the targetRef can be updated to reflect the changes
// in featureRef. It checks if sufficient authorizations / approvals exist for
// the merge to happen, indicated by the error being nil. Additionally, a
// boolean value is also returned that indicates whether a final authorized
// signature is still necessary via the RSL entry for the merge.
//
// Summary of return combinations:
// (false, err) -> merge is not possible
// (false, nil) -> merge is possible and can be performed by anyone
// (true,  nil) -> merge is possible but it MUST be performed by an authorized
// person for the rule, i.e., an authorized person must sign the merge's RSL
// entry
func (r *Repository) VerifyMergeable(ctx context.Context, targetRef, featureRef string, opts ...verifymergeableopts.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Repository) VerifyNetwork(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// verifyRefTip inspects the specified reference in the local repository to
// check if it points to the expected Git object.
func (r *Repository) verifyRefTip(target string, expectedTip gitinterface.Hash) error {
	_ = "STUB: not implemented"
	return nil
}
