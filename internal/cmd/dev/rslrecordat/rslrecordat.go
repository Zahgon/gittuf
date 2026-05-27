// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package rslrecordat

import (
	"github.com/spf13/cobra"
)

type options struct {
	targetID       string
	signingKeyPath string
	dstRef         string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

func (o *options) Run(_ *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
