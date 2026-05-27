// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package init

import (
	"github.com/spf13/cobra"
)

type options struct{}

func (o *options) AddFlags(_ *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) Run(_ *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
