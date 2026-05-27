// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package addpolicykey

import (
	"github.com/gittuf/gittuf/internal/cmd/trust/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p          *persistent.Options
	targetsKey string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
