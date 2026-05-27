// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/experimental/gittuf/options/root"
	trustpolicyopts "github.com/gittuf/gittuf/experimental/gittuf/options/trustpolicy"
	"github.com/gittuf/gittuf/internal/policy"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
)

var (
	ErrNoHookName         = errors.New("hook name not provided")
	ErrInvalidHookTimeout = errors.New("hook timeout must be greater than 1 second")
)

// InitializeRoot is the interface for the user to create the repository's root
// of trust.
func (r *Repository) InitializeRoot(ctx context.Context, signer sslibdsse.SignerVerifier, signCommit bool, opts ...root.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) SetRepositoryLocation(ctx context.Context, signer sslibdsse.SignerVerifier, location string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddRootKey is the interface for the user to add an authorized key
// for the Root role.
func (r *Repository) AddRootKey(ctx context.Context, signer sslibdsse.SignerVerifier, newRootKey tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveRootKey is the interface for the user to de-authorize a key
// trusted to sign the Root role.
func (r *Repository) RemoveRootKey(ctx context.Context, signer sslibdsse.SignerVerifier, keyID string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddTopLevelTargetsKey is the interface for the user to add an authorized key
// for the top level Targets role / policy file.
func (r *Repository) AddTopLevelTargetsKey(ctx context.Context, signer sslibdsse.SignerVerifier, targetsKey tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveTopLevelTargetsKey is the interface for the user to de-authorize a key
// trusted to sign the top level Targets role / policy file.
func (r *Repository) RemoveTopLevelTargetsKey(ctx context.Context, signer sslibdsse.SignerVerifier, targetsKeyID string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGitHubApp is the interface for the user to add the authorized key for the
// trusted GitHub app. This key is used to verify GitHub pull request approval
// attestation signatures recorded by the app.
func (r *Repository) AddGitHubApp(ctx context.Context, signer sslibdsse.SignerVerifier, appName string, appKey tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveGitHubApp is the interface for the user to de-authorize the key for the
// special GitHub app role.
func (r *Repository) RemoveGitHubApp(ctx context.Context, signer sslibdsse.SignerVerifier, appName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TrustGitHubApp updates the root metadata to mark GitHub app pull request
// approvals as trusted.
func (r *Repository) TrustGitHubApp(ctx context.Context, signer sslibdsse.SignerVerifier, appName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UntrustGitHubApp updates the root metadata to mark GitHub app pull request
// approvals as untrusted.
func (r *Repository) UntrustGitHubApp(ctx context.Context, signer sslibdsse.SignerVerifier, appName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateRootThreshold sets the threshold of valid signatures required for the
// Root role.
func (r *Repository) UpdateRootThreshold(ctx context.Context, signer sslibdsse.SignerVerifier, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTopLevelTargetsThreshold sets the threshold of valid signatures
// required for the top level Targets role.
func (r *Repository) UpdateTopLevelTargetsThreshold(ctx context.Context, signer sslibdsse.SignerVerifier, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGlobalRuleThreshold adds a threshold global rule to the root metadata.
func (r *Repository) AddGlobalRuleThreshold(ctx context.Context, signer sslibdsse.SignerVerifier, name string, patterns []string, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGlobalRuleBlockForcePushes adds a global rule that blocks force pushes to the root metadata.
func (r *Repository) AddGlobalRuleBlockForcePushes(ctx context.Context, signer sslibdsse.SignerVerifier, name string, patterns []string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateGlobalRuleThreshold updates an existing threshold global rule in the root metadata.
func (r *Repository) UpdateGlobalRuleThreshold(ctx context.Context, signer sslibdsse.SignerVerifier, name string, patterns []string, threshold int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateGlobalRuleBlockForcePushes updates an existing block-force-pushes global rule in the root metadata.
func (r *Repository) UpdateGlobalRuleBlockForcePushes(ctx context.Context, signer sslibdsse.SignerVerifier, name string, patterns []string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveGlobalRule removes a global rule from the root metadata.
func (r *Repository) RemoveGlobalRule(ctx context.Context, signer sslibdsse.SignerVerifier, name string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) AddPropagationDirective(ctx context.Context, signer sslibdsse.SignerVerifier, directiveName, upstreamRepository, upstreamReference, upstreamPath, downstreamReference, downstreamPath string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) UpdatePropagationDirective(ctx context.Context, signer sslibdsse.SignerVerifier, directiveName, upstreamRepository, upstreamReference, upstreamPath, downstreamReference, downstreamPath string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) IncrementRootVersion(ctx context.Context, signer sslibdsse.SignerVerifier, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Just pass it to updateRootMetadata as it will increment the version

func (r *Repository) RemovePropagationDirective(ctx context.Context, signer sslibdsse.SignerVerifier, name string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// ListPropagationDirectives lists all propagation directives defined in the
// root metadata.
func (r *Repository) ListPropagationDirectives(ctx context.Context, policyRef string) ([]tuf.PropagationDirective, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddHook defines the workflow for adding a file to be executed as a hook. It
// writes the hook file, populates all fields in the hooks metadata associated
// with this file and commits it to the root of trust metadata.
func (r *Repository) AddHook(ctx context.Context, signer sslibdsse.SignerVerifier, stages []tuf.HookStage, hookName string, hookBytes []byte, environment tuf.HookEnvironment, principalIDs []string, timeout int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: hash agility

// RemoveHook defines the workflow for removing a hook defined in gittuf policy.
func (r *Repository) RemoveHook(ctx context.Context, signer sslibdsse.SignerVerifier, stages []tuf.HookStage, hookName string, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateHook updates the hook specified by stage and hookName with the new
// principalIDs, hashes, environment, and timeout.
func (r *Repository) UpdateHook(ctx context.Context, signer sslibdsse.SignerVerifier, stages []tuf.HookStage, hookName string, hookBytes []byte, environment tuf.HookEnvironment, principalIDs []string, timeout int, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// EnableController makes the current repository a "controller" repository used
// to specify gittuf policies for other repositories.
func (r *Repository) EnableController(ctx context.Context, signer sslibdsse.SignerVerifier, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// DisableController makes the repository not a controller for a network of
// gittuf repositories. Any policies declared in this repository will not be
// enforced for other repositories part of the network.
func (r *Repository) DisableController(ctx context.Context, signer sslibdsse.SignerVerifier, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddControllerRepository adds a repository as a controller to the current
// repository.
func (r *Repository) AddControllerRepository(ctx context.Context, signer sslibdsse.SignerVerifier, repositoryName, repositoryLocation string, initialRootPrincipals []tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddNetworkRepository adds a repository as part of the network overseen by the
// current repository.
func (r *Repository) AddNetworkRepository(ctx context.Context, signer sslibdsse.SignerVerifier, repositoryName, repositoryLocation string, initialRootPrincipals []tuf.Principal, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// SignRoot adds a signature to the Root envelope. Note that the metadata itself
// is not modified, so its version remains the same.
func (r *Repository) SignRoot(ctx context.Context, signer sslibdsse.SignerVerifier, signCommit bool, opts ...trustpolicyopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) loadRootMetadata(state *policy.State, keyID string) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}

func (r *Repository) updateRootMetadata(ctx context.Context, state *policy.State, signer sslibdsse.SignerVerifier, rootMetadata tuf.RootMetadata, commitMessage string, createRSLEntry, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}
