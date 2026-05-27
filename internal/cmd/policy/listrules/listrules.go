// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package listrules

import (
	"github.com/spf13/cobra"
)

type options struct {
	targetRef string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// Iterate through the rules, they are already in order, and the depth tells us how to indent.
// The order is a pre-order traversal of the delegation tree, so that the parent is always before the children.

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
