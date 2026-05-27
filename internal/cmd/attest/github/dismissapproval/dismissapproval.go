// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package dismissapproval

import (
	"github.com/gittuf/gittuf/internal/cmd/attest/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p                 *persistent.Options
	baseURL           string
	reviewID          int64
	dismissedApprover string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
