// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	"context"
)

type rule struct {
	name      string
	pattern   string
	key       string
	threshold int
}

// getCurrRules returns the current rules from the policy file.
func getCurrRules(ctx context.Context, o *options) []rule { _ = "STUB: not implemented"; return nil }

// repoAddRule adds a rule to the policy file.
func repoAddRule(ctx context.Context, o *options, rule rule, authorizedPrincipalIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// repoUpdateRule updates an existing rule in the policy file.
func repoUpdateRule(ctx context.Context, o *options, r rule, authorizedPrincipalIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// repoRemoveRule removes a rule from the policy file.
func repoRemoveRule(ctx context.Context, o *options, rule rule) error {
	_ = "STUB: not implemented"
	return nil
}

// repoReorderRules reorders the rules in the policy file.
func repoReorderRules(ctx context.Context, o *options, rules []rule) error {
	_ = "STUB: not implemented"
	return nil
}
