// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package reconcile

import (
	"github.com/spf13/cobra"
)

type options struct{}

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
