// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package clone

import (
	"github.com/gittuf/gittuf/internal/cmd/common"
	"github.com/spf13/cobra"
)

type options struct {
	branch           string
	expectedRootKeys common.PublicKeys
	bare             bool
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
