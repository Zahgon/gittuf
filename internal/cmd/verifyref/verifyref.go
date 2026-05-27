// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package verifyref

import (
	"github.com/spf13/cobra"
)

type options struct {
	latestOnly    bool
	fromEntry     string
	remoteRefName string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
