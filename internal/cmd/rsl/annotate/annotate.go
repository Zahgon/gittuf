// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package annotate

import (
	"github.com/spf13/cobra"
)

type options struct {
	skip       bool
	message    string
	remoteName string
	localOnly  bool
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
