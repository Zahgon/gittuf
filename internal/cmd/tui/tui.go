// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	"context"

	"github.com/gittuf/gittuf/internal/cmd/policy/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p          *persistent.Options
	policyName string
	targetRef  string
	readOnly   bool
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }

// startTUI intitializes a new model for the TUI
func startTUI(ctx context.Context, o *options) error { _ = "STUB: not implemented"; return nil }
