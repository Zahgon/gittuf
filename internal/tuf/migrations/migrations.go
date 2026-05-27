// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package migrations

import (
	tufv01 "github.com/gittuf/gittuf/internal/tuf/v01"
	tufv02 "github.com/gittuf/gittuf/internal/tuf/v02"
)

/*
	adityasaky: We should probably have an automatic migration to the _next_
	version every time so that it's easy to migrate from v01 to v0k easily using
	consecutive migration functions. However, I think building that out now may
	be overkill; I don't know if we expect a bunch of schema changes.
*/

// MigrateRootMetadataV01ToV02 converts tufv01.RootMetadata into
// tufv02.RootMetadata.
func MigrateRootMetadataV01ToV02(rootMetadata *tufv01.RootMetadata) *tufv02.RootMetadata {
	_ = "STUB: not implemented"
	return nil
}

// Set same expires

// Set same version number

// Set repository location

// Set keys

// Set roles

// Set app attestations support

// Set global rules

// Set propagations

// Set hooks

// MigrateTargetsMetadataV01ToV02 converts tufv01.TargetsMetadata into
// tufv02.TargetsMetadata.
func MigrateTargetsMetadataV01ToV02(targetsMetadata *tufv01.TargetsMetadata) *tufv02.TargetsMetadata {
	_ = "STUB: not implemented"
	return nil
}

// Set same expires

// Set same version number

// Set delegations
