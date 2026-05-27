// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/internal/tuf"
)

var (
	ErrPushingPolicy     = errors.New("unable to push policy")
	ErrPullingPolicy     = errors.New("unable to pull policy")
	ErrNoRemoteSpecified = errors.New("no remote specified to push policy")
)

// PushPolicy pushes the local gittuf policy to the specified remote. As this
// push defaults to fast-forward only, divergent policy states are detected.
// Note that this also pushes the RSL as the policy cannot change without an
// update to the RSL.
func (r *Repository) PushPolicy(remoteName string) error { _ = "STUB: not implemented"; return nil }

// PullPolicy fetches gittuf policy from the specified remote. The fetches is
// marked as fast forward only to detect divergence. Note that this also fetches
// the RSL as the policy must be updated in sync with the RSL.
func (r *Repository) PullPolicy(remoteName string) error { _ = "STUB: not implemented"; return nil }

// HasPolicy indicates if the repository has a gittuf policy applied.
func (r *Repository) HasPolicy() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *Repository) ApplyPolicy(ctx context.Context, remoteName string, localOnly, signRSLEntry bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) DiscardPolicy() error { _ = "STUB: not implemented"; return nil }

type DelegationWithDepth struct {
	Delegation tuf.Rule
	Depth      int
}

func (r *Repository) ListRules(ctx context.Context, targetRef string) ([]*DelegationWithDepth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allDelegations will be the returned list of all the delegations in pre-order traversal, no delegations will be popped off

// We construct localDelegations first so that we preserve the order
// of delegations in currentMetadata in delegationsToSearch

func (r *Repository) ListPrincipals(ctx context.Context, targetRef, policyName string) (map[string]tuf.Principal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListGlobalRules returns a list of all global rules as an array of tuf.GlobalRules.
func (r *Repository) ListGlobalRules(ctx context.Context, targetRef string) ([]tuf.GlobalRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) ListHooks(ctx context.Context, targetRef string) (map[tuf.HookStage][]tuf.Hook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) StagePolicy(ctx context.Context, remoteName string, localOnly, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}
