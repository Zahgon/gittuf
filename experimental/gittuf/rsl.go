// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	rslopts "github.com/gittuf/gittuf/experimental/gittuf/options/rsl"
	"github.com/gittuf/gittuf/internal/rsl"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const gittufTransportPrefix = "gittuf::"

var (
	ErrCommitNotInRef              = errors.New("specified commit is not in ref")
	ErrPushingRSL                  = errors.New("unable to push RSL")
	ErrPullingRSL                  = errors.New("unable to pull RSL")
	ErrDivergedRefs                = errors.New("references in local repository have diverged from upstream")
	ErrRemoteNotSpecified          = errors.New("remote not specified")
	ErrCannotUseRemoteAndLocalOnly = errors.New("cannot indicate local-only and push to specified remote")
)

// RecordRSLEntryForReference is the interface for the user to add an RSL entry
// for the specified Git reference.
func (r *Repository) RecordRSLEntryForReference(ctx context.Context, refName string, signCommit bool, opts ...rslopts.RecordOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Track localRefName to check the expected tip as we may override refName

// dst differs from src
// Eg: git push <remote> <src>:<dst>

// The tip of the ref is always from the localRefName

// TODO: once policy verification is in place, the signing key used by
// signCommit must be verified for the refName in the delegation tree.

// RecordRSLEntryForReferenceAtTarget is a special version of
// RecordRSLEntryForReference used for evaluation. It is only invoked when
// gittuf is explicitly set in developer mode.
func (r *Repository) RecordRSLEntryForReferenceAtTarget(refName, targetID string, signingKeyBytes []byte, opts ...rslopts.RecordOption) error {
	_ = "STUB: not implemented"
	// Double check that gittuf is in developer mode
	return nil
}

// dst differs from src
// Eg: git push <remote> <src>:<dst>

// TODO: once policy verification is in place, the signing key used by
// signCommit must be verified for the refName in the delegation tree.

func (r *Repository) SkipAllInvalidReferenceEntriesForRef(targetRef string, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RecordRSLAnnotation is the interface for the user to add an RSL annotation
// for one or more prior RSL entries.
func (r *Repository) RecordRSLAnnotation(ctx context.Context, rslEntryIDs []string, skip bool, message string, signCommit bool, opts ...rslopts.AnnotateOption) error {
	_ = "STUB: not implemented"
	// TODO: local only?
	return nil
}

// TODO: once policy verification is in place, the signing key used by
// signCommit must be verified for the refNames of the rslEntryIDs.

// ReconcileLocalRSLWithRemote checks the local RSL against the specified remote
// and reconciles the local RSL if needed. If the local RSL doesn't exist or is
// strictly behind the remote RSL, then the local RSL is updated to match the
// remote RSL. If the local RSL is ahead of the remote RSL, nothing is updated.
// Finally, if the local and remote RSLs have diverged, then the local only RSL
// entries are reapplied over the latest entries in the remote if the local only
// RSL entries and remote only entries are for different Git references.
func (r *Repository) ReconcileLocalRSLWithRemote(ctx context.Context, remoteName string, sign bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// Fetch status of RSL on the remote

// Load status of the local RSL for comparison

// Check if local is nil and exit appropriately

// Local RSL has not been populated but remote is not zero
// Fetch updates to the local RSL

// Check if equal and exit early if true

// Next, check if remote is ahead of local

// If not ancestor, local may be ahead or they may have diverged
// If remote is ancestor, only local is ahead, no updates
// If remote is not ancestor, the two have diverged, local needs to pull updates

// We don't push to the remote RSL, that's handled alongside
// other pushes (eg. via the transport) or explicitly

// This is the tricky one
// First, we find a common ancestor for the two
// Second, we identify all the entries in the local that is not in the
// remote
// Third, we set local to the remote's tip
// Fourth, we apply all the entries that we identified over the new tip

// Check if remote has entries for refs that are also updated locally
// We don't want to do conflict resolution right now

// Set local RSL to match the remote state

// Apply local only entries on top of the new local RSL
// localOnlyEntries is in reverse order

// We create a new object so as to apply anything the
// entry may contain that is inferred at commit time
// For example, an incrementing number inferred from the
// parent entry

// Sync is responsible for synchronizing references between the local copy of
// the repository and the specified remote.
func (r *Repository) Sync(ctx context.Context, remoteName string, overwriteLocalRefs, signCommit bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) sync(remoteName string, overwriteLocalRefs bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// Fetch status of RSL on the remote

//nolint:errcheck

// Load status of the local RSL for comparison

// Check if equal and exit early if true

// non error exit

// local RSL is ahead of remote RSL

// remote RSL is ahead of local RSL -> check if any local ref changes
// conflict and display message to user

// Track the latest tips in the remote RSL using the entries that are new
// compared to the local RSL

// Find local tip for same ref

// Fetch remote objects for each ref

// Now we actually have remoteTip in the object store

// if remoteTip is ahead of localTip, we're good
// otherwise, mark that ref as candidate for overwriting locally

// If tags (or other ref->obj mappings) are not equal, mark that
// ref as candidate for overwriting locally

// non error exit
// TODO: restore worktree if checked out HEAD is in divergedRefs

// The RSL itself has diverged
// We can't fix this if overwriteLocalRefs is not true

// Find local tip for same ref

// Fetch remote objects for each ref

// Now we actually have remoteTip in the object store

// if remoteTip is ahead of localTip, we're good
// otherwise, mark that ref as candidate for overwriting locally

// If tags (or other ref->obj mappings) are not equal, mark that
// ref as candidate for overwriting locally

// non error exit

func getRSLEntriesUntil(repo *gitinterface.Repository, start, until gitinterface.Hash) ([]rsl.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getLatestRefTipsFromRSLEntries(entries []rsl.Entry) map[string]gitinterface.Hash {
	_ = "STUB: not implemented"
	return nil
}

// PushRSL pushes the local RSL to the specified remote. As this push defaults
// to fast-forward only, divergent RSL states are detected.
func (r *Repository) PushRSL(remoteName string) error { _ = "STUB: not implemented"; return nil }

// PullRSL pulls RSL contents from the specified remote to the local RSL. The
// fetch is marked as fast forward only to detect RSL divergence.
func (r *Repository) PullRSL(remoteName string) error { _ = "STUB: not implemented"; return nil }

// isDuplicateEntry checks if the latest unskipped entry for the ref has the
// same target ID. Note that it's legal for the RSL to have target A, then B,
// then A again, this is not considered a duplicate entry
func (r *Repository) isDuplicateEntry(refName string, targetID gitinterface.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// PropagateChangesFromUpstreamRepositories invokes gittuf's propagation
// workflow. It inspects the latest policy metadata to find the applicable
// propagation directives, and executes the workflow on each one.
func (r *Repository) PropagateChangesFromUpstreamRepositories(ctx context.Context, sign bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Group directives for the same repository together

// Calculate base64 encoding of the URL: this is the subdirectory where
// the contents are sent

// FIXME: this assumes tufv01.PropagationDirective

// directive name

// upstream location

// upstream ref

// upstream path

// downstream ref

// downstream path

// FIXME: we're cloning some repositories twice, once to see the
// manifest to resolve controller graph, another time to actually
// propagate contents

// DFS to resolve transitive propagations

//nolint:errcheck

//nolint:errcheck

// TODO: we see this error when required upstream ref isn't found, handle gracefully?

// TODO: atomic? abort?
