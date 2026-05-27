// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/internal/common/set"
	policyopts "github.com/gittuf/gittuf/internal/policy/options/policy"
	"github.com/gittuf/gittuf/internal/rsl"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const (
	// PolicyRef defines the Git namespace used for gittuf policies.
	PolicyRef = "refs/gittuf/policy"

	// PolicyStagingRef defines the Git namespace used as a staging area when creating or updating gittuf policies.
	PolicyStagingRef = "refs/gittuf/policy-staging"

	// RootRoleName defines the expected name for the gittuf root of trust.
	RootRoleName = "root"

	// TargetsRoleName defines the expected name for the top level gittuf policy file.
	TargetsRoleName = "targets"

	// DefaultCommitMessage defines the fallback message to use when updating the policy ref if an action specific message is unavailable.
	DefaultCommitMessage = "Update policy state"

	metadataTreeEntryName = "metadata"

	gitReferenceRuleScheme = "git"
	fileRuleScheme         = "file"
)

var (
	ErrMetadataNotFound              = errors.New("unable to find requested metadata file; has it been initialized?")
	ErrDanglingDelegationMetadata    = errors.New("unreachable targets metadata found")
	ErrPolicyNotFound                = errors.New("cannot find policy")
	ErrInvalidPolicy                 = errors.New("invalid policy state (is policy reference out of sync with corresponding RSL entry?)")
	ErrNotAncestor                   = errors.New("cannot apply changes since policy is not an ancestor of the policy staging")
	ErrControllerMetadataNotFound    = errors.New("requested controller repository metadata not found")
	ErrControllerMetadataNotVerified = errors.New("unable to verify controller repository metadata")
)

// State contains the full set of metadata and root keys present in a policy
// state.
type State struct {
	Metadata           *StateMetadata
	ControllerMetadata map[string]*StateMetadata

	Hooks map[tuf.HookStage][]tuf.Hook

	GitHubApps map[string]tuf.GitHubApp

	repository     *gitinterface.Repository
	loadedEntry    rsl.ReferenceUpdaterEntry
	verifiersCache map[string][]*SignatureVerifier
	ruleNames      *set.Set[string]
	allPrincipals  map[string]tuf.Principal
	hasFileRule    bool
	globalRules    map[string][]tuf.GlobalRule
}

type StateMetadata struct {
	RootEnvelope        *sslibdsse.Envelope
	TargetsEnvelope     *sslibdsse.Envelope
	DelegationEnvelopes map[string]*sslibdsse.Envelope
}

func (s *StateMetadata) GetRootMetadata(migrate bool) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}

func (s *StateMetadata) getRootMetadataFromBytes(metadataBytes []byte, migrate bool) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}

// this is tufv01
// Something that's not tufv01 may also lack the schemaVersion field and
// enter this code path. At that point, we're relying on the unmarshal
// to return something that's close to tufv01. We may see strange bugs
// if this happens, but it's also likely someone trying to submit
// incorrect metadata / trigger a version rollback, which we do want to
// be aware of.

func (s *StateMetadata) GetTargetsMetadata(roleName string, migrate bool) (tuf.TargetsMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.TargetsMetadata), nil
}

// this is tufv01
// Something that's not tufv01 may also lack the schemaVersion field and
// enter this code path. At that point, we're relying on the unmarshal
// to return something that's close to tufv01. We may see strange bugs
// if this happens, but it's also likely someone trying to submit
// incorrect metadata / trigger a version rollback, which we do want to
// be aware of.

func (s *StateMetadata) WriteTree(repo *gitinterface.Repository) (gitinterface.Hash, error) {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash), nil
}

// LoadState returns the State of the repository's policy corresponding to the
// entry. It verifies the root of trust for the state from the initial policy
// entry in the RSL. If no policy states are found and the entry is for the
// policy-staging ref, that entry is returned with no verification.
func LoadState(ctx context.Context, repo *gitinterface.Repository, requestedEntry rsl.ReferenceUpdaterEntry, opts ...policyopts.LoadStateOption) (*State, error) {
	_ = "STUB: not implemented"
	// Regardless of whether we've been asked for policy ref or staging ref,
	// we want to examine and verify consecutive policy states that appear
	// before the entry. This is why we don't just load the state and return
	// if entry is for the staging ref.
	return nil, nil
}

// TODO: should this searcher be inherited when invoked via Verifier?

// we don't have a policy entry yet
// we just return the state for the requested entry

// check if firstPolicyEntry is **after** requested entry
// this can happen when the requested entry is for policy-staging before
// Apply() was ever called

// the first policy entry knows the requested entry, meaning the
// requested entry is an ancestor of the first policy entry
// we just return the state for the requested entry

// If requestedEntry.RefName == policy, then allPolicyEntries includes requestedEntry
// If requestedEntry.RefName == policy-staging, then allPolicyEntries does not include requestedEntry

// We load the very first policy entry with no additional verification,
// the root keys are implicitly trusted

// The searcher _may_ include refs/gittuf/attestations
// etc. which should be skipped

// We've already loaded it and done successive verification as
// it was included in allPolicyEntries
// This state is stored in verifiedState, we can do an internal
// verification check and return

// This is reached when requestedEntry is for staging ref
// We've checked that all the policy states prior to this staging entry
// are good (with their root of trust)

// LoadCurrentState returns the State corresponding to the repository's current
// active policy. It verifies the root of trust for the state starting from the
// initial policy entry in the RSL.
func LoadCurrentState(ctx context.Context, repo *gitinterface.Repository, ref string, opts ...policyopts.LoadStateOption) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Note: this will not set the loadedEntry field in the policy state

// LoadFirstState returns the State corresponding to the repository's first
// active policy. It does not verify the root of trust since it is the initial policy.
func LoadFirstState(ctx context.Context, repo *gitinterface.Repository, opts ...policyopts.LoadStateOption) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindVerifiersForPath identifies the trusted set of verifiers for the
// specified path. While walking the delegation graph for the path, signatures
// for delegated metadata files are verified using the verifier context.
func (s *State) FindVerifiersForPath(path string) ([]*SignatureVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cache hit for this path in this policy

// This has to go first so it's prioritized during verification
// At least one global rule exists, return an exhaustive verifier

// we'll add all principals below

// threshold doesn't matter since we set verifyExhaustively to true

// very important!

// Note: we could loop through all global constraints and create a
// verifier with all principals but targeting a specific constraint (or
// an aggregate constraint that has the highest threshold requirement of
// all the constraints that match path). However, this probably paints
// us into a corner (only threshold requirements between two constraints
// can be compared, we may have uncomparable constraints later), and we
// would also want to verify every applicable global constraint for
// safety, so we would be doing extra work for no reason.

// add to cache

// return verifiers

func (s *State) findVerifiersForPathIfProtected(path string) ([]*SignatureVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No policies exist

// This envelope is verified when state is loaded, as this is
// the start for all delegation graph searches
// migrating is fine since this is purely a query, let's start using tufv02 metadata

// each entry is a list of delegations from a particular metadata file

// Exit condition: Only allow rule found in the current group
// => len(currentDelegationGroup) <= 1

// migrating is fine since this is purely a query, let's start using tufv02 metadata

// Add the current metadata's further delegations upfront to
// be depth-first

// Stop processing current delegation group, but proceed
// with other groups

func (s *State) GetAllPrincipals() map[string]tuf.Principal { _ = "STUB: not implemented"; return nil }

// Verify verifies the contents of the State for internal consistency.
// Specifically, it checks that the root keys in the root role match the ones
// stored on disk in the state. Further, it also verifies the signatures of the
// top level Targets role and all reachable delegated Targets roles. Any
// unreachable role returns an error.
func (s *State) Verify(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Check GitHub app approvals
// don't migrate: this may be for a write and we don't want to write tufv02 metadata yet

// TODO: retire IsGitHubAppApprovalTrusted

// Check that the GitHub app role is declared

// Check top-level targets and delegations

// don't migrate: this may be for a write and we don't want to write tufv02 metadata yet

// Check reachable delegations

// Exit condition: The last entry in the queue is always the allow
// rule, which we don't process during DFS

// don't migrate: this may be for a write and we don't want to write tufv02 metadata yet

// Check controller root metadata

//nolint:errcheck

// We need to LoadState() the state from which the root is derived
// For that, we need to know when it was propagated into this repository

// Check this entry

// not found yet
// find propagation entry in local repo

// We know propagationEntry is of this type because of the rsl.IsPropagationEntryForReference opt

// LoadState does full verification up until the requested entry

// TODO: verify git tree ID in upstream matches propagated

// Commit verifies and writes the State to the policy-staging namespace.
func (s *State) Commit(repo *gitinterface.Repository, commitMessage string, createRSLEntry, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get treeIDs for state.Metadata and each of the state.ControllerMetadata entries

// We must reset to original policy commit if err != nil from here onwards.

// Apply takes valid changes from the policy staging ref, and fast-forward
// merges it into the policy ref. Apply only takes place if the latest state on
// the policy staging ref is valid. This prevents invalid changes to the policy
// taking affect, and allowing new changes, that until signed by multiple users
// would be invalid to be made, by utilizing the policy staging ref.
func Apply(ctx context.Context, repo *gitinterface.Repository, signRSLEntry bool) error {
	_ = "STUB: not implemented"
	// First, reconcile staging with policy
	return nil
}

// Get the reference for the PolicyRef

// case 1: both found -> verify tip matches entry
// case 2: only one found -> return error
// case 3: neither found -> nothing to verify

// Nothing to check or return here

// Get the reference for the PolicyStagingRef

// Check if the PolicyStagingRef is ahead of PolicyRef (fast-forward)

// This check ensures that the policy staging branch is a direct forward progression of the policy branch,
// preventing any overwrites of policy history and maintaining a linear policy evolution, since a
// fast-forward merge does not work with a non-linear history.

// This is only being checked if there are no problems finding the tip of the policy ref, since if there
// is no tip, then it cannot be an ancestor of the tip of the policy staging ref

// using LoadCurrentState to load and verify if the PolicyStagingRef's
// latest state is valid

// Update the reference for the base to point to the new commit

// Discard resets the policy staging ref, discarding any changes made to the policy staging ref.
func Discard(repo *gitinterface.Repository) error { _ = "STUB: not implemented"; return nil }

// Reset PolicyStagingRef to match the actual policy ref

func ReconcileStaging(repo *gitinterface.Repository, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the reference for the PolicyRef

// case 1: both found -> verify tip matches entry
// case 2: only one found -> return error
// case 3: neither found -> nothing to verify

// Nothing to check or return here

// Get the reference for the PolicyStagingRef

// case 1: both found -> verify tip matches entry
// case 2: only one found -> return error
// case 3: neither found -> nothing to verify

// Nothing to check or return here

// If neither policy nor policy staging is found, then return nil, as
// there's nothing to reconcile.

// If policy isn't found but policy staging is, then return nil as well.

// There are a few possible scenarios here.
// A) policy = policy-staging -> nothing to do; this is the case when the
// last update to the policy related refs were applied
// B) policy is behind policy-staging -> nothing to do
// C) policy is strictly ahead of policy-staging -> reconciliation is
// necessary, and HAS to be because a change landed directly in policy
// without going through policy-staging, ff update policy-staging as well
// D) policy and policy-staging have diverged -> reconciliation is
// necessary, and HAS to be because policy-staging was updated AND policy
// was propagated into, meaning they have divergent (but non conflicting
// changes), ff update does not suffice
// Reconciliation goal: policy-staging must be ff-ahead of policy so Apply()
// is not affected by controller propagations

// Reconciliation overview:
// In case C, ff-update policy-staging to include the propagated changes and
// record RSL reference entry (not propagation entry).
// In case D, "stash" changes in staging, apply policy ref changes over
// common ancestor, re-apply stashed changes into policy-staging.
// This requires rewriting policy-staging's history on clients, but luckily,
// this cannot result in a conflict in the current workflows.
// This is because the unapplied changes to policy-staging are necessarily
// in the state's metadata (and not in the controller metadata). The changes
// in policy that don't exist in policy-staging MUST be due to controller
// propagation, i.e., completely different files are updated. We know this
// to be true because an update to policy's local repository metadata MUST
// have gone through policy-staging and been applied, so propagation is the
// only legitimate reason for policy to have a change not seen in
// policy-staging.

// nothing to do

// nothing to do

// update staging to match policy

// Diverged
// Create new policy-staging that is "rebased"

// TODO: fix RSL entries for staging that are now orphaned

// This includes the changes made in staging + the controller changes
// propagated into policy

func (s *State) GetRootKeys() ([]tuf.Principal, error) { _ = "STUB: not implemented"; return nil, nil }

// don't migrate: this may be for a write and we don't want to write tufv02 metadata yet

// GetRootMetadata returns the deserialized payload of the State's RootEnvelope.
// The `migrate` parameter determines if the schema must be converted to a newer
// version.
func (s *State) GetRootMetadata(migrate bool) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}

func (s *State) GetControllerRootMetadata(controllerName string) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}

// never migrate controller metadata

// GetTargetsMetadata returns the deserialized payload of the State's
// TargetsEnvelope for the specified `roleName`.  The `migrate` parameter
// determines if the schema must be converted to a newer version.
func (s *State) GetTargetsMetadata(roleName string, migrate bool) (tuf.TargetsMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.TargetsMetadata), nil
}

func (s *State) HasTargetsRole(roleName string) bool { _ = "STUB: not implemented"; return false }

func (s *State) HasRuleName(name string) bool { _ = "STUB: not implemented"; return false }

// preprocess handles several "one time" tasks when the state is first loaded.
// This includes things like loading the set of rule names present in the state,
// checking if it has file rules, etc.
func (s *State) preprocess() error { _ = "STUB: not implemented"; return nil }

func (s *State) getRootVerifier() (*SignatureVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *State) getTargetsVerifier() (*SignatureVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadStateForEntry returns the State for a specified RSL reference entry for
// the policy namespace. This helper is focused on reading the Git object store
// and loading the policy contents. Typically, LoadCurrentState of LoadState
// must be used. The exception is VerifyRelative... which performs root
// verification between consecutive policy states.
func loadStateForEntry(repo *gitinterface.Repository, entry rsl.ReferenceUpdaterEntry) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadStateFromCommit(repo *gitinterface.Repository, commitID gitinterface.Hash) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// metadataQueue is populated with metadata/ subtrees we want to load for
// either the current repository or its controllers.

type policyTreeItem struct {
	name   string
	treeID gitinterface.Hash
}
