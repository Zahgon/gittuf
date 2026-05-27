// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	"context"
)

type globalRule struct {
	ruleName     string
	ruleType     string
	rulePatterns []string
	threshold    int
}

// getGlobalRules returns a slice of globalRule for the TUI
func getGlobalRules(ctx context.Context, o *options) []globalRule {
	_ = "STUB: not implemented"
	return nil
}

// repoAddGlobalRule adds a global rule
func repoAddGlobalRule(ctx context.Context, o *options, gr globalRule) error {
	_ = "STUB: not implemented"
	return nil
}

// repoRemoveGlobalRule removes a global rule
func repoRemoveGlobalRule(ctx context.Context, o *options, gr globalRule) error {
	_ = "STUB: not implemented"
	return nil
}

// repoUpdateGlobalRule updates a global rule
func repoUpdateGlobalRule(ctx context.Context, o *options, gr globalRule) error {
	_ = "STUB: not implemented"
	return nil
}
