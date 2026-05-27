// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package listprincipals

import (
	"github.com/spf13/cobra"
)

const indentString = "    "

type options struct {
	policyRef  string
	policyName string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
