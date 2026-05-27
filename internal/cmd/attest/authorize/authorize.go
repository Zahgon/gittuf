// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package authorize

import (
	"github.com/gittuf/gittuf/internal/cmd/attest/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p       *persistent.Options
	fromRef string
	revoke  bool
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
