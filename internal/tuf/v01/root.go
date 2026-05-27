// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v01

import (
	"github.com/gittuf/gittuf/internal/common/set"
	"github.com/gittuf/gittuf/internal/tuf"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const (
	rootVersion = "https://gittuf.dev/policy/root/v0.1"
)

// RootMetadata defines the schema of TUF's Root role.
type RootMetadata struct {
	Type               string                     `json:"type"`
	Expires            string                     `json:"expires"`
	Version            uint64                     `json:"version"`
	RepositoryLocation string                     `json:"repositoryLocation,omitempty"`
	Keys               map[string]*Key            `json:"keys"`
	Roles              map[string]Role            `json:"roles"`
	GitHubApps         map[string]*GitHubApp      `json:"githubApps,omitempty"`
	GlobalRules        []tuf.GlobalRule           `json:"globalRules,omitempty"`
	Propagations       []tuf.PropagationDirective `json:"propagations,omitempty"`
	MultiRepository    *MultiRepository           `json:"multiRepository,omitempty"`
	Hooks              map[tuf.HookStage][]*Hook  `json:"hooks,omitempty"`
}

// NewRootMetadata returns a new instance of RootMetadata.
func NewRootMetadata() *RootMetadata { _ = "STUB: not implemented"; return nil }

// SetExpires sets the expiry date of the RootMetadata to the value passed in.
func (r *RootMetadata) SetExpires(expires string) { _ = "STUB: not implemented"; return }

// GetSchemaVersion returns the metadata schema version.
func (r *RootMetadata) GetSchemaVersion() string {
	_ = "STUB: not implemented"

	// GetVersion returns the version number of the metadata.
	return ""
}

func (r *RootMetadata) GetVersion() uint64 {
	_ = "STUB: not implemented"

	// IncrementVersion increments the metadata version number by 1.
	return 0
}

func (r *RootMetadata) IncrementVersion() {
	_ = "STUB: not implemented"

	// GetRepositoryLocation returns the canonical location of the Git repository.
	return
}

func (r *RootMetadata) GetRepositoryLocation() string { _ = "STUB: not implemented"; return "" }

// SetRepositoryLocation sets the specified repository location in the root
// metadata.
func (r *RootMetadata) SetRepositoryLocation(location string) { _ = "STUB: not implemented"; return }

// AddRootPrincipal adds the specified key to the root metadata and authorizes the key
// for the root role.
func (r *RootMetadata) AddRootPrincipal(key tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// Add key to metadata

// Create a new root role entry with this key

// Add key ID to the root role if it's not already in it

// DeleteRootPrincipal removes keyID from the list of trusted Root public keys
// in rootMetadata. It does not remove the key entry itself as it does not check
// if other roles can be verified using the same key.
func (r *RootMetadata) DeleteRootPrincipal(keyID string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPrimaryRuleFilePrincipal adds the 'targetsKey' as a trusted public key in
// 'rootMetadata' for the top level Targets role.
func (r *RootMetadata) AddPrimaryRuleFilePrincipal(key tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// Add key to the metadata file

// Create a new targets role entry with this key

// DeletePrimaryRuleFilePrincipal removes the key matching 'keyID' from trusted
// public keys for top level Targets role in 'rootMetadata'. Note: It doesn't
// remove the key entry itself as it doesn't check if other roles can use the
// same key.
func (r *RootMetadata) DeletePrimaryRuleFilePrincipal(keyID string) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGitHubAppPrincipal adds the 'appKey' as a trusted public key in
// 'rootMetadata' for the special GitHub app role. This key is used to verify
// GitHub pull request approval attestation signatures.
func (r *RootMetadata) AddGitHubAppPrincipal(name string, key tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: support multiple keys / threshold for app

// DeleteGitHubAppPrincipal removes the special GitHub app role from the root
// metadata.
func (r *RootMetadata) DeleteGitHubAppPrincipal(name string) { _ = "STUB: not implemented"; return }

// EnableGitHubAppApprovals sets GitHubApprovalsTrusted to true in the
// root metadata.
func (r *RootMetadata) EnableGitHubAppApprovals(appName string) { _ = "STUB: not implemented"; return }

// DisableGitHubAppApprovals sets GitHubApprovalsTrusted to false in the root
// metadata.
func (r *RootMetadata) DisableGitHubAppApprovals(appName string) { _ = "STUB: not implemented"; return }

func (r *RootMetadata) GetGitHubAppEntries() (map[string]tuf.GitHubApp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateRootThreshold sets the threshold for the Root role.
func (r *RootMetadata) UpdateRootThreshold(threshold int) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdatePrimaryRuleFileThreshold sets the threshold for the top level Targets
// role.
func (r *RootMetadata) UpdatePrimaryRuleFileThreshold(threshold int) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPrincipals returns all the principals in the root metadata.
func (r *RootMetadata) GetPrincipals() map[string]tuf.Principal {
	_ = "STUB: not implemented"
	return nil
}

// GetRootThreshold returns the threshold of principals that must sign the root
// of trust metadata.
func (r *RootMetadata) GetRootThreshold() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GetRootPrincipals returns the principals trusted for the root of trust
// metadata.
func (r *RootMetadata) GetRootPrincipals() ([]tuf.Principal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPrimaryRuleFileThreshold returns the threshold of principals that must
// sign the primary rule file.
func (r *RootMetadata) GetPrimaryRuleFileThreshold() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetPrimaryRuleFilePrincipals returns the principals trusted for the primary
// rule file.
func (r *RootMetadata) GetPrimaryRuleFilePrincipals() ([]tuf.Principal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsGitHubAppApprovalTrusted indicates if the GitHub app is trusted.
//
// TODO: this needs to be generalized across tools
func (r *RootMetadata) IsGitHubAppApprovalTrusted(appName string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetGitHubAppPrincipals returns the principals trusted for the GitHub app
// attestations.
//
// TODO: this needs to be generalized across tools
func (r *RootMetadata) GetGitHubAppPrincipals(appName string) ([]tuf.Principal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddGlobalRule adds a new global rule to RootMetadata.
func (r *RootMetadata) AddGlobalRule(globalRule tuf.GlobalRule) error {
	_ = "STUB: not implemented"
	return nil
}

// check for duplicates

// DeleteGlobalRule removes the specified global rule from the RootMetadata.
func (r *RootMetadata) DeleteGlobalRule(ruleName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RootMetadata) GetGlobalRules() []tuf.GlobalRule { _ = "STUB: not implemented"; return nil }

// UpdateGlobalRule updates the specified global rule from the RootMetadata.
func (r *RootMetadata) UpdateGlobalRule(globalRule tuf.GlobalRule) error {
	_ = "STUB: not implemented"
	return nil
}

// AddPropagationDirective adds a propagation directive to the root metadata.
func (r *RootMetadata) AddPropagationDirective(directive tuf.PropagationDirective) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdatePropagationDirective updates a propagation directive in the root
// metadata.
func (r *RootMetadata) UpdatePropagationDirective(directive tuf.PropagationDirective) error {
	_ = "STUB: not implemented"
	return nil
}

// GetPropagationDirectives returns the propagation directives found in the root
// metadata.
func (r *RootMetadata) GetPropagationDirectives() []tuf.PropagationDirective {
	_ = "STUB: not implemented"
	return nil

	// DeletePropagationDirective removes a propagation directive from the root
	// metadata.
}

func (r *RootMetadata) DeletePropagationDirective(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsController indicates if the repository serves as the controller for a
// multi-repository gittuf network.
func (r *RootMetadata) IsController() bool { _ = "STUB: not implemented"; return false }

// EnableController marks the current repository as a controller repository.
func (r *RootMetadata) EnableController() error { _ = "STUB: not implemented"; return nil }

// TODO: what if it's already a controller? noop?

// DisableController marks the current repository as not-a-controller.
func (r *RootMetadata) DisableController() error { _ = "STUB: not implemented"; return nil }

// nothing to do

// TODO: should we remove the network repository entries?

// AddControllerRepository adds the specified repository as a controller for the
// current repository.
func (r *RootMetadata) AddControllerRepository(name, location string, initialRootPrincipals []tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// AddNetworkRepository adds the specified repository as part of the network for
// which the current repository is a controller. The current repository must be
// marked as a controller before this can be used.
func (r *RootMetadata) AddNetworkRepository(name, location string, initialRootPrincipals []tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// EnableController must be called first

// GetControllerRepositories returns the repositories that serve as the
// controllers for the networks the current repository is a part of.
func (r *RootMetadata) GetControllerRepositories() []tuf.OtherRepository {
	_ = "STUB: not implemented"
	return nil
}

// GetNetworkRepositories returns the repositories that are part of the network
// for which the current repository is a controller. IsController must return
// true for this to be set.
func (r *RootMetadata) GetNetworkRepositories() []tuf.OtherRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *RootMetadata) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// this type _has_ to be a copy of RootMetadata, minus the use of
	// json.RawMessage for tuf interfaces
	return nil
}

// addKey adds a key to the RootMetadata instance.
func (r *RootMetadata) addKey(key tuf.Principal) error { _ = "STUB: not implemented"; return nil }

// addRole adds a role object and associates it with roleName in the
// RootMetadata instance.
func (r *RootMetadata) addRole(roleName string, role Role) { _ = "STUB: not implemented"; return }

type GlobalRuleThreshold struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Paths     []string `json:"paths"`
	Threshold int      `json:"threshold"`
}

func NewGlobalRuleThreshold(name string, paths []string, threshold int) *GlobalRuleThreshold {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalRuleThreshold) GetName() string { _ = "STUB: not implemented"; return "" }

func (g *GlobalRuleThreshold) Matches(path string) bool { _ = "STUB: not implemented"; return false }

// We validate pattern when it's added to / updated in the metadata

func (g *GlobalRuleThreshold) GetProtectedNamespaces() []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalRuleThreshold) GetThreshold() int { _ = "STUB: not implemented"; return 0 }

type GlobalRuleBlockForcePushes struct {
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Paths []string `json:"paths"`
}

func NewGlobalRuleBlockForcePushes(name string, paths []string) (*GlobalRuleBlockForcePushes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: set prefix correctly

func (g *GlobalRuleBlockForcePushes) GetName() string { _ = "STUB: not implemented"; return "" }

func (g *GlobalRuleBlockForcePushes) Matches(path string) bool {
	_ = "STUB: not implemented"
	return false
}

// We validate pattern when it's added to / updated in the metadata

func (g *GlobalRuleBlockForcePushes) GetProtectedNamespaces() []string {
	_ = "STUB: not implemented"
	return nil
}

type PropagationDirective struct {
	Name                string `json:"name"`
	UpstreamRepository  string `json:"upstreamRepository"`
	UpstreamReference   string `json:"upstreamReference"`
	UpstreamPath        string `json:"upstreamPath"`
	DownstreamReference string `json:"downstreamReference"`
	DownstreamPath      string `json:"downstreamPath"`
}

func (p *PropagationDirective) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PropagationDirective) GetUpstreamRepository() string { _ = "STUB: not implemented"; return "" }

func (p *PropagationDirective) GetUpstreamReference() string { _ = "STUB: not implemented"; return "" }

func (p *PropagationDirective) GetUpstreamPath() string { _ = "STUB: not implemented"; return "" }

func (p *PropagationDirective) GetDownstreamReference() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PropagationDirective) GetDownstreamPath() string { _ = "STUB: not implemented"; return "" }

func NewPropagationDirective(name, upstreamRepository, upstreamReference, upstreamPath, downstreamReference, downstreamPath string) tuf.PropagationDirective {
	_ = "STUB: not implemented"
	return *new(tuf.PropagationDirective)
}

type MultiRepository struct {
	Controller             bool               `json:"controller"`
	ControllerRepositories []*OtherRepository `json:"controllerRepositories,omitempty"`
	NetworkRepositories    []*OtherRepository `json:"networkRepositories,omitempty"`
}

func (m *MultiRepository) IsController() bool { _ = "STUB: not implemented"; return false }

func (m *MultiRepository) GetControllerRepositories() []tuf.OtherRepository {
	_ = "STUB: not implemented"
	return nil
}

func (m *MultiRepository) GetNetworkRepositories() []tuf.OtherRepository {
	_ = "STUB: not implemented"
	return nil
}

type OtherRepository struct {
	Name                  string `json:"name"`
	Location              string `json:"location"`
	InitialRootPrincipals []*Key `json:"initialRootPrincipals"`
}

func (o *OtherRepository) GetName() string { _ = "STUB: not implemented"; return "" }

func (o *OtherRepository) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (o *OtherRepository) GetInitialRootPrincipals() []tuf.Principal {
	_ = "STUB: not implemented"
	return nil
}

// AddHook adds the specified hook to the metadata.
func (r *RootMetadata) AddHook(stages []tuf.HookStage, hookName string, principalIDs []string, hashes map[string]string, environment tuf.HookEnvironment, timeout int) (tuf.Hook, error) {
	_ = "STUB: not implemented"
	// TODO: Check if principal exists in RootMetadata/TargetsMetadata
	return *new(tuf.Hook), nil
}

// UpdateHook updates the hook specified by stage and hookName with the new
// principalIDs, hashes, environment, and timeout.
func (r *RootMetadata) UpdateHook(stages []tuf.HookStage, hookName string, principalIDs []string, hashes map[string]string, environment tuf.HookEnvironment, timeout int) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveHook removes the hook specified by stage and hookName.
func (r *RootMetadata) RemoveHook(stages []tuf.HookStage, hookName string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHooks returns the hooks for the specified stage.
func (r *RootMetadata) GetHooks(stage tuf.HookStage) ([]tuf.Hook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Hook defines the schema for a hook.
type Hook struct {
	Name         string              `json:"name"`
	PrincipalIDs *set.Set[string]    `json:"principals"`
	Hashes       map[string]string   `json:"hashes"`
	Environment  tuf.HookEnvironment `json:"environment"`
	Timeout      int                 `json:"timeout"`
}

// ID returns the identifier of the hook, its name.
func (h *Hook) ID() string {
	_ = "STUB: not implemented"

	// GetPrincipalIDs returns the principals that must run this hook.
	return ""
}

func (h *Hook) GetPrincipalIDs() *set.Set[string] { _ = "STUB: not implemented"; return nil }

// GetHashes returns the hashes of the hook file.
func (h *Hook) GetHashes() map[string]string { _ = "STUB: not implemented"; return nil }

func (h *Hook) GetBlobID() gitinterface.Hash {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash)
}

// GetEnvironment returns the environment that the hook is to run in.
func (h *Hook) GetEnvironment() tuf.HookEnvironment {
	_ = "STUB: not implemented"
	return *

	// GetTimeout returns the maximum duration the hook can run for, in seconds.
	new(tuf.HookEnvironment)
}

func (h *Hook) GetTimeout() int { _ = "STUB: not implemented"; return 0 }

type GitHubApp struct {
	Trusted      bool             `json:"trusted"`
	PrincipalIDs *set.Set[string] `json:"principalIDs"`
	Threshold    int              `json:"threshold"`
}

func (g *GitHubApp) GetPrincipalIDs() []string { _ = "STUB: not implemented"; return nil }

func (g *GitHubApp) GetThreshold() int { _ = "STUB: not implemented"; return 0 }

func (g *GitHubApp) IsTrusted() bool { _ = "STUB: not implemented"; return false }
