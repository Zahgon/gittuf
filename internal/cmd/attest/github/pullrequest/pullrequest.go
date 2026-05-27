// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package pullrequest

import (
	"github.com/gittuf/gittuf/internal/cmd/attest/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p                 *persistent.Options
	baseURL           string
	repository        string
	pullRequestNumber int
	commitID          string
	baseBranch        string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

// When we're using commit, we need the base branch to filter through nested
// pull requests

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
