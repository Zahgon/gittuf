// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package verifymergeable

import (
	"github.com/spf13/cobra"
)

type options struct {
	baseBranch    string
	featureBranch string
	bypassRSL     bool
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
