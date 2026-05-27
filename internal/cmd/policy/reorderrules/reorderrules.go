// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package reorderrules

import (
	"github.com/gittuf/gittuf/internal/cmd/policy/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p          *persistent.Options
	policyName string
	ruleNames  []string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
