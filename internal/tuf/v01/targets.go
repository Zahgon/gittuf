// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v01

import (
	"encoding/json"

	"github.com/gittuf/gittuf/internal/common/set"
	"github.com/gittuf/gittuf/internal/tuf"
)

const (
	targetsVersion = "http://gittuf.dev/policy/rule-file/v0.1"
)

// TargetsMetadata defines the schema of TUF's Targets role.
type TargetsMetadata struct {
	Type        string         `json:"type"`
	Expires     string         `json:"expires"`
	Version     uint64         `json:"version"`
	Targets     map[string]any `json:"targets"`
	Delegations *Delegations   `json:"delegations"`
}

// NewTargetsMetadata returns a new instance of TargetsMetadata.
func NewTargetsMetadata() *TargetsMetadata { _ = "STUB: not implemented"; return nil }

// SetExpires sets the expiry date of the TargetsMetadata to the value passed
// in.
func (t *TargetsMetadata) SetExpires(expires string) { _ = "STUB: not implemented"; return }

// GetSchemaVersion returns the metadata schema version.
func (t *TargetsMetadata) GetSchemaVersion() string { _ = "STUB: not implemented"; return "" }

// GetVersion returns the version number of the metadata.
func (t *TargetsMetadata) GetVersion() uint64 {
	_ = "STUB: not implemented"

	// IncrementVersion increments the metadata version number by 1.
	return 0
}

func (t *TargetsMetadata) IncrementVersion() {
	_ = "STUB: not implemented"

	// Validate ensures the instance of TargetsMetadata matches gittuf expectations.
	return
}

func (t *TargetsMetadata) Validate() error { _ = "STUB: not implemented"; return nil }

// AddRule adds a new delegation to TargetsMetadata.
func (t *TargetsMetadata) AddRule(ruleName string, authorizedPrincipalIDs, rulePatterns []string, threshold int) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateRule is used to amend a delegation in TargetsMetadata.
func (t *TargetsMetadata) UpdateRule(ruleName string, authorizedPrincipalIDs, rulePatterns []string, threshold int) error {
	_ = "STUB: not implemented"
	return nil
}

// ReorderRules changes the order of delegations, and the new order is specified
// in `ruleNames []string`.
func (t *TargetsMetadata) ReorderRules(ruleNames []string) error {
	_ = "STUB: not implemented"
	// Create a map of all existing delegations for quick look up
	return nil
}

// Create a set of current rules in metadata, skipping the allow rule

// Create newDelegations and set it in the targetsMetadata after adding allow rule

// RemoveRule deletes a delegation entry from TargetsMetadata.
func (t *TargetsMetadata) RemoveRule(ruleName string) error { _ = "STUB: not implemented"; return nil }

// GetPrincipals returns all the principals in the rule file.
func (t *TargetsMetadata) GetPrincipals() map[string]tuf.Principal {
	_ = "STUB: not implemented"
	return nil
}

// GetRules returns all the rules in the metadata.
func (t *TargetsMetadata) GetRules() []tuf.Rule { _ = "STUB: not implemented"; return nil }

// AddPrincipal adds a principal to the metadata.
//
// TODO: this isn't associated with a specific rule; with the removal of
// verify-commit and verify-tag, it may not make sense anymore
func (t *TargetsMetadata) AddPrincipal(principal tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdatePrincipal is not a valid operation for tufv01 metadata, as principals
// were introduced in tufv02. For keys in tufv01, we only support adding new
// ones and removing existing ones.
func (t *TargetsMetadata) UpdatePrincipal(_ tuf.Principal) error {
	_ = "STUB: not implemented"
	return nil
}

// RemovePrincipal removes a principal from the metadata.
func (t *TargetsMetadata) RemovePrincipal(principalID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delegations defines the schema for specifying delegations in TUF's Targets
// metadata.
type Delegations struct {
	Keys  map[string]*Key `json:"keys"`
	Roles []*Delegation   `json:"roles"`
}

// addKey adds a delegations key.
func (d *Delegations) addKey(key tuf.Principal) error { _ = "STUB: not implemented"; return nil }

func (d *Delegations) removeKey(keyID string) error { _ = "STUB: not implemented"; return nil }

// AllowRule returns the default, last rule for all policy files.
func AllowRule() *Delegation { _ = "STUB: not implemented"; return nil }

// Delegation defines the schema for a single delegation entry. It differs from
// the standard TUF schema by allowing a `custom` field to record details
// pertaining to the delegation. It implements the tuf.Rule interface.
type Delegation struct {
	Name        string           `json:"name"`
	Paths       []string         `json:"paths"`
	Terminating bool             `json:"terminating"`
	Custom      *json.RawMessage `json:"custom,omitempty"`
	Role
}

// ID returns the identifier of the delegation, its name.
func (d *Delegation) ID() string {
	_ = "STUB: not implemented"

	// Matches checks if any of the delegation's patterns match the target.
	return ""
}

func (d *Delegation) Matches(target string) bool { _ = "STUB: not implemented"; return false }

// We validate pattern when it's added to / updated in the metadata

// GetPrincipalIDs returns the identifiers of the principals that are listed as
// trusted by the rule.
func (d *Delegation) GetPrincipalIDs() *set.Set[string] {
	_ = "STUB: not implemented"

	// GetThreshold returns the threshold of principals that must approve to meet
	// the rule.
	return nil
}

func (d *Delegation) GetThreshold() int {
	_ = "STUB: not implemented"

	// IsLastTrustedInRuleFile indicates that subsequent rules in the rule file are
	// not to be trusted if the current rule matches the namespace under
	// verification (similar to TUF's terminating behavior). However, the current
	// rule's delegated rules as well as other rules already in the queue are
	// trusted.
	return 0
}

func (d *Delegation) IsLastTrustedInRuleFile() bool { _ = "STUB: not implemented"; return false }

// GetProtectedNamespaces returns the set of namespaces protected by the
// delegation.
func (d *Delegation) GetProtectedNamespaces() []string { _ = "STUB: not implemented"; return nil }
