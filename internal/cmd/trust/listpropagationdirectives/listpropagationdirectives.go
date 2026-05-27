// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package listpropagationdirectives

import (
	"github.com/spf13/cobra"
)

type options struct {
	targetRef string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// TODO: switch to the display package

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
