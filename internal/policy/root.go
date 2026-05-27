// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"github.com/gittuf/gittuf/internal/tuf"
)

// InitializeRootMetadata initializes a new instance of tuf.RootMetadata with
// default values and a given key. The default values are version set to 1,
// expiry date set to one year from now, and the provided key is added.
func InitializeRootMetadata(key tuf.Principal) (tuf.RootMetadata, error) {
	_ = "STUB: not implemented"
	return *new(tuf.RootMetadata), nil
}
