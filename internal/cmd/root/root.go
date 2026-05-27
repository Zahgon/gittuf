// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package root

import (
	"github.com/spf13/cobra"
)

type options struct {
	noColor           bool
	verbose           bool
	profile           bool
	cpuProfileFile    string
	memoryProfileFile string
}

func (o *options) AddFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func (o *options) PreRunE(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Check if colored output must be disabled
	return nil
}

// Setup logging

// Start profiling if flag is set

func New() *cobra.Command { _ = "STUB: not implemented"; return nil }
