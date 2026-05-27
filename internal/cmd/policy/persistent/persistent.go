// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package persistent

import (
	"github.com/spf13/cobra"
)

type Options struct {
	SigningKey   string
	WithRSLEntry bool
}

func (o *Options) AddPersistentFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }
