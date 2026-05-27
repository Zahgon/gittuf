// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package dismissgithubapproval

import (
	"github.com/spf13/cobra"
)

type options struct {
	signingKey string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
