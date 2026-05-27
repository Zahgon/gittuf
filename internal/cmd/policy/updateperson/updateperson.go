// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package updateperson

import (
	"github.com/gittuf/gittuf/internal/cmd/policy/persistent"
	"github.com/spf13/cobra"
)

type options struct {
	p                    *persistent.Options
	policyName           string
	personID             string
	publicKeys           []string
	associatedIdentities []string
	customMetadata       []string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (o *options) Run(cmd *cobra.Command, _ []string) error { _ = "STUB: not implemented"; return nil }

// Process public keys

// Process associated identities

// Process custom metadata

// Create a new person with updated fields

// Update the person in the policy

func New(persistent *persistent.Options) *cobra.Command { _ = "STUB: not implemented"; return nil }
