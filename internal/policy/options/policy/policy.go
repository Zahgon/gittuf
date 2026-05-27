// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import "github.com/gittuf/gittuf/internal/tuf"

type LoadStateOptions struct {
	InitialRootPrincipals []tuf.Principal
	BypassRSL             bool
}

type LoadStateOption func(*LoadStateOptions)

func WithInitialRootPrincipals(initialRootPrincipals []tuf.Principal) LoadStateOption {
	_ = "STUB: not implemented"
	return *new(LoadStateOption)
}

func BypassRSL() LoadStateOption { _ = "STUB: not implemented"; return *new(LoadStateOption) }
