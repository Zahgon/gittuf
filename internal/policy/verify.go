// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/internal/attestations"
	"github.com/gittuf/gittuf/internal/cache"
	"github.com/gittuf/gittuf/internal/common/set"
	"github.com/gittuf/gittuf/internal/rsl"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

var (
	ErrVerificationFailed                                = errors.New("gittuf policy verification failed")
	ErrInvalidEntryNotSkipped                            = errors.New("invalid entry found not marked as skipped")
	ErrLastGoodEntryIsSkipped                            = errors.New("entry expected to be unskipped is marked as skipped")
	ErrNoVerifiers                                       = errors.New("no verifiers present for verification")
	ErrInvalidVerifier                                   = errors.New("verifier has invalid parameters (is threshold 0?)")
	ErrVerifierConditionsUnmet                           = errors.New("verifier's key and threshold constraints not met")
	ErrCannotVerifyMergeableForTagRef                    = errors.New("cannot verify mergeable into tag reference")
	ErrNetworkRepositoryDoesNotDeclareRequiredController = errors.New("network repository does not declare required controller repository")
	ErrNetworkRepositoryHasStaleControllerMetadata       = errors.New("network repository has not fetched latest controller metadata")
	ErrMetadataRollbackDetected                          = errors.New("gittuf policy metadata rollback detected")
)

// PolicyVerifier implements various gittuf verification workflows.
type PolicyVerifier struct { //nolint:revive
	// We want to call this PolicyVerifier to avoid any confusion with
	// SignatureVerifier.

	repo     *gitinterface.Repository
	searcher searcher

	persistentCacheEnabled bool
	persistentCache        *cache.Persistent
}

func NewPolicyVerifier(repo *gitinterface.Repository) *PolicyVerifier {
	_ = "STUB: not implemented"
	return nil
}

// VerifyRef verifies the signature on the latest RSL entry for the target ref
// using the latest policy. The expected Git ID for the ref in the latest RSL
// entry is returned if the policy verification is successful.
func (v *PolicyVerifier) VerifyRef(ctx context.Context, target string) (gitinterface.Hash, error) {
	_ = "STUB: not implemented"
	// Find latest entry for target
	return *new(gitinterface.Hash), nil
}

// VerifyRefFull verifies the entire RSL for the target ref from the first
// entry. The expected Git ID for the ref in the latest RSL entry is returned if
// the policy verification is successful.
func (v *PolicyVerifier) VerifyRefFull(ctx context.Context, target string) (gitinterface.Hash, error) {
	_ = "STUB: not implemented"
	// Trace RSL back to the start
	return *new(gitinterface.Hash), nil
}

// break because we've loaded the entry and don't need to fallthrough

// Find latest entry for target

// VerifyRefFromEntry performs verification for the reference from a specific
// RSL entry. The expected Git ID for the ref in the latest RSL entry is
// returned if the policy verification is successful.
func (v *PolicyVerifier) VerifyRefFromEntry(ctx context.Context, target string, entryID gitinterface.Hash) (gitinterface.Hash, error) {
	_ = "STUB: not implemented"
	// Load starting point entry
	return *new(gitinterface.Hash), nil
}

// TODO: we should instead find the latest reference entry
// before the entryID and use that

// Find latest entry for target

// Do a relative verify from start entry to the latest entry

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
func (v *PolicyVerifier) VerifyMergeable(ctx context.Context, targetRef, featureRef string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// VerifyMergeableForCommit checks if the targetRef can be updated to reflect
// the changes in featureID. It checks if sufficient authorizations / approvals
// exist for the merge to happen, indicated by the error being nil.
// Additionally, a boolean value is also returned that indicates whether a final
// authorized signature is still necessary via the RSL entry for the merge.
// Note: this function DOES NOT use the RSL to identify the tip of the feature
// ref.
//
// Summary of return combinations:
// (false, err) -> merge is not possible
// (false, nil) -> merge is possible and can be performed by anyone
// (true,  nil) -> merge is possible but it MUST be performed by an authorized
// person for the rule, i.e., an authorized person must sign the merge's RSL
// entry
func (v *PolicyVerifier) VerifyMergeableForCommit(ctx context.Context, targetRef string, featureID gitinterface.Hash) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (v *PolicyVerifier) verifyMergeable(ctx context.Context, targetRef string, fromID, featureID gitinterface.Hash) (bool, error) {
	_ = "STUB: not implemented"
	// We're specifically focused on commit merges here, this doesn't apply to
	// tags
	return false, nil
}

// Load latest policy

// Load latest attestations

// Attestations are not compulsory, so return err only
// if it's some other error

// Verify modified files

// this will be set after one successful verification of the commit to avoid repeated signature verification

// If we've already verified and identified commit signature, we can
// just check if that verifier is trusted for the new path. If not
// found, we don't make any assumptions about it being a failure in
// case of name mismatches. So, the signature check proceeds as
// usual. Also, we don't use verifyMergeable=true here. File
// verification rules are not met using the signature on the RSL
// entry, so we don't count threshold-1 here.

func (v *PolicyVerifier) VerifyNetwork(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Use the policy searcher to find the latest applicable policy entry
	return nil
}

// Load the policy state, read root.json to enumerate network repositories
// we want to inspect.

// If the metadata says it's not a controller, there are no network
// repositories to inspect.

// We want the tree of the target commit in the policy ref entry. Actually,
// we want the tree ID of the `metadata` subtree. This subdirectory is what
// we expect to have been propagated into each network repo.

// Clone the network repository into tmp

//nolint:errcheck

// Identify the most recent entry in the network repo's RSL that is for
// the policy ref

// Load the policy state of the network repository, verify the initial
// roots

// Check that at least one of the declared controller repositories in
// the network repository's root is for the controller.
// TODO: this is a miss if a repo uses a different transport protocol
// for the metadata.

// We hit this when none of the controller declarations match

// Check if it has correctly propagated changes

// gittuf stores the controller metadata in a subdirectory in the policy
// ref that includes b64 encoded controller repo location

// TODO: should we check the propagation entry's upstream target matches
// the latest controller policy RSL entry?

// Check that the propagated tree ID matches the controller's `metadata`
// subdirectory tree

// VerifyRelativeForRef verifies the RSL between specified start and end entries
// using the provided policy entry for the first entry.
func (v *PolicyVerifier) VerifyRelativeForRef(ctx context.Context, firstEntry, lastEntry rsl.ReferenceUpdaterEntry, target string) error {
	_ = "STUB: not implemented"
	/*
		require firstEntry != nil
		require lastEntry != nil
		require target != ""
	*/return nil
}

//nolint:errcheck

// Load policy applicable at firstEntry

// Searcher gives us nil when firstEntry is the very first entry
// or close to it (i.e., before a policy was applied)

// require currentPolicy != nil || parent(firstEntry) == nil

// Attestations are not compulsory, so return err only
// if it's some other error

// require currentAttestations != nil || (entry.Ref != attestations.Ref for entry in 0..firstEntry)

// Enumerate RSL entries between firstEntry and lastEntry, ignoring irrelevant ones

// require len(entries) != 0

// Verify each entry, looking for a fix when an invalid entry is encountered

// invariant invalidEntry == nil || inRecoveryMode() == true

// Pop entry from queue

// We've already loaded this policy

// require newPolicy != nil

// currentPolicy can be nil when
// verifying from the beginning of the
// RSL entry and we only have staging
// refs

// If the invalid entry is never marked as skipped, we return err

// The invalid entry's been marked as skipped but we still need
// to see if another entry fixed state for non-gittuf users

// Fix entry does not exist after revoking annotation

// Verification has passed, add to cache

// This is only reached when we have an invalid state.
// First, the verification workflow determines the last good state for
// the ref. This is needed to evaluate whether a fix for the invalid
// state is available. After this is found, the workflow looks through
// the remaining entries in the queue to find the fix. Until the fix is
// found, entries encountered that are for other refs are added to a new
// queue. Entries that are for the same ref but not the fix are
// considered invalid. The workflow enters a valid state again when a)
// the fix entry (which hasn't also been revoked) is found, and b) all
// entries for the ref in the invalid range are marked as skipped by an
// annotation. If these conditions don't both hold, the workflow returns
// an error. After the fix is found, all remaining entries in the
// original queue are also added to the new queue. The new queue then
// takes the place of the original queue. This ensures that all entries
// are processed even when an invalid state is reached.

// 1. What's the last good state?

// this type assertion is fine because we use the rsl.IsReferenceEntry opt

// require lastGoodEntry != nil

// TODO: what if the very first entry for a ref is a violation?

// gittuf requires the fix to point to a commit that is tree-same as the
// last good state

// 2. What entries do we have in the current verification set for the
// ref? The first one that is tree-same as lastGoodEntry's commit is the
// fix. Entries prior to that one in the queue are considered invalid
// and must be skipped

// Unrelated entry that must be processed in the outer loop
// Currently this is just policy entries

// propagation entry cannot be a fix entry

// Fix found, we append the rest of the current verification set
// to the new entry queue
// But first, we must check that this fix hasn't been skipped
// If it has been skipped, it's not actually a fix and we need
// to keep looking

// newEntry is not tree-same / commit-same, so it is automatically
// invalid, check that it's been marked as revoked

// If we haven't found a fix, return the original error

// We may have found a fix but if an invalid intermediate entry
// wasn't skipped, return error

// Reset these trackers to continue verification with rest of the queue
// We may encounter other issues

func (s *StateMetadata) VerifyNewStateMetadata(_ context.Context, newStateMetadata *StateMetadata) error {
	_ = "STUB: not implemented"
	// Check new state's root version number is >= current state's root version number
	return nil
}

// Check new state's rule files have version numbers >= current state's rule file version numbers

// First off, compare the primary rule file

// If the current state doesn't have a primary rule file, we can
// return early as there is no risk of rollback with the rule files
// in the new state.

// At this point, we know the current state has a primary rule file, so
// if the new state doesn't have one, it's a rollback. This MAY change
// when we support deleting rule files, though maybe we don't allow that
// for primary rule files.

// Then compare all delegated rule files.
// We need to check that the rule files in the new state with the same name
// as a rule file in the current state have greater or equal version numbers.
// The new state may have a rule file added that doesn't exist in the
// current state, but it shouldn't have any removed rule files compared to
// the current state as we don't yet support deleting rule files.

// At this point, we know the current state has this delegated rule
// file, so if the new state doesn't have it, it's a rollback.
// This will change once we support deleting rule files.

// VerifyNewState ensures that when a new policy is encountered, its root role
// is signed by keys trusted in the current policy.
func (s *State) VerifyNewState(ctx context.Context, newPolicy *State) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify state metadata to protect against rollback attacks

// Verify controller state metadata to protect against rollback attacks upstream

// verifyEntry is a helper to verify an entry's signature using the specified
// policy. The specified policy is used for the RSL entry itself. However, for
// commit signatures, verifyEntry checks when the commit was first introduced
// via the RSL across all refs. Then, it uses the policy applicable at the
// commit's first entry into the repository. If the commit is brand new to the
// repository, the specified policy is used.
func verifyEntry(ctx context.Context, repo *gitinterface.Repository, policy *State, attestationsState *attestations.Attestations, entry *rsl.ReferenceEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// Load the applicable reference authorization and approvals from trusted
// code review systems

// Verify Git namespace policies using the RSL entry and attestations

// Check if policy has file rules at all for efficiency

// No file rules to verify

// Verify modified files

// First, get all commits between the current and last entry for the ref.
// note: this is ordered by commit ID

// this will be set after one successful verification of the commit to avoid repeated signature verification

// If we've already verified and identified commit signature, we
// can just check if that verifier is trusted for the new path.
// If not found, we don't make any assumptions about it being a
// failure in case of name mismatches. So, the signature check
// proceeds as usual.

func verifyTagEntry(ctx context.Context, repo *gitinterface.Repository, policy *State, attestationsState *attestations.Attestations, entry *rsl.ReferenceEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func getApproverAttestationAndKeyIDs(ctx context.Context, repo *gitinterface.Repository, policy *State, attestationsState *attestations.Attestations, entry *rsl.ReferenceEntry) (*sslibdsse.Envelope, *set.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// We need to handle the case where we're approving a tag
// For a tag, the expected toID in the approval is the commit the tag points to
// Otherwise, the expected toID is the tree the commit points to

func getApproverAttestationAndKeyIDsForIndex(ctx context.Context, repo *gitinterface.Repository, policy *State, attestationsState *attestations.Attestations, targetRef string, fromID, toID gitinterface.Hash, isTag bool) (*sslibdsse.Envelope, *set.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// When we add other code review systems, we can move this into a
// generalized helper that inspects the attestations for each system trusted
// in policy.
// We only use this flow right now for non-tags as tags cannot be approved
// on currently supported systems
// TODO: support multiple apps / threshold per system

// if it exists

// TODO: support multiple versions

// getCommits identifies the commits introduced to the entry's ref since the
// last RSL entry for the same ref. These commits are then verified for file
// policies.
func getCommits(repo *gitinterface.Repository, entry *rsl.ReferenceEntry) ([]gitinterface.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyGitObjectAndAttestationsOptions contains the configurable options for
// verifyGitObjectAndAttestations.
type verifyGitObjectAndAttestationsOptions struct {
	approverPrincipalIDs *set.Set[string]
	verifyMergeable      bool
	trustedVerifier      string
	tagObjectID          gitinterface.Hash
}

type verifyGitObjectAndAttestationsOption func(o *verifyGitObjectAndAttestationsOptions)

// withApproverPrincipalIDs allows for optionally passing in approver IDs to
// verifyGitObjectAndAttestations. These IDs may be obtained via a code review
// tool such as GitHub pull request approvals.
func withApproverPrincipalIDs(approverPrincipalIDs *set.Set[string]) verifyGitObjectAndAttestationsOption {
	_ = "STUB: not implemented"
	return *new(verifyGitObjectAndAttestationsOption)
}

// withVerifyMergeable indicates that the verification must check if a change
// can be merged.
func withVerifyMergeable() verifyGitObjectAndAttestationsOption {
	_ = "STUB: not implemented"
	return *new(verifyGitObjectAndAttestationsOption)
}

// withTrustedVerifier is used to specify the name of a verifier that has
// already been used to verify in the past. If the newly discovered set of
// verifiers includes the trusted verifier, then we can return early.
func withTrustedVerifier(name string) verifyGitObjectAndAttestationsOption {
	_ = "STUB: not implemented"
	return *new(verifyGitObjectAndAttestationsOption)
}

// withTagObjectID is used to set the Git ID of a tag object. When this is set,
// the tag object's signature is also verified in addition to the RSL entry for
// the tag.
func withTagObjectID(objID gitinterface.Hash) verifyGitObjectAndAttestationsOption {
	_ = "STUB: not implemented"
	return *new(verifyGitObjectAndAttestationsOption)
}

func verifyGitObjectAndAttestations(ctx context.Context, policy *State, target string, gitID gitinterface.Hash, authorizationAttestation *sslibdsse.Envelope, opts ...verifyGitObjectAndAttestationsOption) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// This target is not protected by gittuf policy

// Verify tag object's signature as well

// explicitly not looking at the attestation
// that applies to the _push_
// thus, we also set threshold to 1

// Signature verification succeeded

// TODO: should we check if a different verifier / signer was
// matched for the tag object compared with the RSL entry?

// Unexpected error

// Haven't found a valid verifier, continue with next verifier

// this is the special case

// We check every global rule

// The global rule applies to the namespace under verification

// Since we're verifying if it's mergeable and we already know
// that the RSL signature is needed to meet threshold, we can
// reduce the global constraint threshold as well

// Check if the verifiedPrincipalIDs meets the required global
// threshold

// TODO: we use policy.repository, not ideal...

// The global rule applies to the namespace under verification

// Cannot check for force pushes for a proposed change

// TODO: should we not look up the entry's afresh in the RSL here?
// the in-memory cache _should_ make this okay, but something to
// consider...

// gitID _must_ be for an RSL reference entry, and we must find
// its predecessor entry.
// Why? Because the rule type only accepts git:<> as patterns.
// If we have another object here, we've gone wrong somewhere.

func verifyGitObjectAndAttestationsUsingVerifiers(ctx context.Context, verifiers []*SignatureVerifier, gitID gitinterface.Hash, authorizationAttestation *sslibdsse.Envelope, appNames []string, approverIDs *set.Set[string], verifyMergeable bool) (string, *set.Set[string], bool, error) {
	_ = "STUB: not implemented"
	return "", nil, false, nil
}

// We meet requirements just from the authorization attestation's sigs

// Unify the principalIDs we've already used with that listed in
// approval attestation
// We ensure that someone who has signed an attestation and is listed in
// the approval attestation is only counted once

// For each approver ID from the app attestation, we try to see
// if it matches a principal in the current verifiers.

// This principal has already been counted towards the
// threshold

// We can only match against a principal if it has a notion
// of associated identities
// Right now, this is just tufv02.Person

// The approver ID from the issuer (appName) matches
// the principal's associated identity for the same
// issuer!

// Get a list of used principals that are also trusted by the verifier

// With approvals, we now meet threshold!

// If verifyMergeable is true, we only need to meet threshold - 1
