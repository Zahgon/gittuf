// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	trustpolicyopts "github.com/gittuf/gittuf/experimental/gittuf/options/trustpolicy"
	"github.com/gittuf/gittuf/internal/policy"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
)

var ErrInvalidPolicyName = errors.New("invalid rule or policy file name, cannot be 'root'")

// InitializeTargets is the interface for the user to create the specified
// policy file.
func (r *Repository) InitializeTargets(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: verify is role can be signed using the presented key. This requires
// the user to pass in the delegating role as well as we do not want to
// assume which role is the delegating role (diamond delegations are legal).
// See: https://github.com/gittuf/gittuf/issues/246.

// AddDelegation is the interface for the user to add a new rule to gittuf
// policy.
func (r *Repository) AddDelegation(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, ruleName string, authorizedPrincipalIDs, rulePatterns []string, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: verify if role can be signed using the presented key. This requires
// the user to pass in the delegating role as well as we do not want to
// assume which role is the delegating role (diamond delegations are legal).
// See: https://github.com/gittuf/gittuf/issues/246.

// UpdateDelegation is the interface for the user to update a rule to gittuf
// policy.
func (r *Repository) UpdateDelegation(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, ruleName string, authorizedPrincipalIDs, rulePatterns []string, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: verify if role can be signed using the presented key. This requires
// the user to pass in the delegating role as well as we do not want to
// assume which role is the delegating role (diamond delegations are legal).
// See: https://github.com/gittuf/gittuf/issues/246.

// ReorderDelegations is the interface for the user to reorder rules in gittuf
// policy.
func (r *Repository) ReorderDelegations(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, ruleNames []string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveDelegation is the interface for a user to remove a rule from gittuf
// policy.
func (r *Repository) RemoveDelegation(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, ruleName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: verify if role can be signed using the presented key. This requires
// the user to pass in the delegating role as well as we do not want to
// assume which role is the delegating role (diamond delegations are legal).
// See: https://github.com/gittuf/gittuf/issues/246.

// AddPrincipalToTargets is the interface for a user to add a trusted principal
// to gittuf rule file metadata.
func (r *Repository) AddPrincipalToTargets(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, authorizedPrincipals []tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: verify is role can be signed using the presented key. This requires
// the user to pass in the delegating role as well as we do not want to
// assume which role is the delegating role (diamond delegations are legal).
// See: https://github.com/gittuf/gittuf/issues/246.

// UpdatePrincipalInTargets is the interface for a user to update a principal's
// information in gittuf rule file metadata.
func (r *Repository) UpdatePrincipalInTargets(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, principal tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemovePrincipalFromTargets is the interface for a user to remove a principal
// from gittuf rule file metadata.
func (r *Repository) RemovePrincipalFromTargets(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, principalID string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// SignTargets adds a signature to specified Targets role's envelope. Note that
// the metadata itself is not modified, so its version remains the same.
func (r *Repository) SignTargets(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) IncrementTargetsVersion(ctx context.Context, signer sslibdsse.SignerVerifier, targetsRoleName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Just pass it to updateTargetsMetadata as it will increment the version

func (r *Repository) updateTargetsMetadata(ctx context.Context, state *policy.State, signer sslibdsse.SignerVerifier, targetsMetadataName string, targetsMetadata tuf.TargetsMetadata, commitMessage string, createRSLEntry, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}
