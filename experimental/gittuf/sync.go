// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/internal/tuf"
)

var (
	ErrCloningRepository          = errors.New("unable to clone repository")
	ErrDirExists                  = errors.New("directory exists")
	ErrExpectedRootKeysDoNotMatch = errors.Join(ErrCloningRepository, errors.New("cloned root keys do not match the expected keys"))
)

// Clone wraps a typical git clone invocation, fetching gittuf refs in addition
// to the standard refs. It performs a verification of the RSL against the
// specified HEAD after cloning the repository.
// TODO: resolve how root keys are trusted / bootstrapped.
func Clone(ctx context.Context, remoteURL, dir, initialBranch string, expectedRootKeys []tuf.Principal, bare bool) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FIXME: my understanding is backslashes are not used in URLs but I haven't dived into the RFCs to check yet

// Trim spaces and trailing slashes if any

// We sort the root keys so that we can check if the root keys array match's the expected root key array
