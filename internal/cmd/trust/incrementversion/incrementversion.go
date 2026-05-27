// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package incrementversion

import (
	"github.com/gittuf/gittuf/internal/cmd/trust/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p *persistent.Options
}

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
