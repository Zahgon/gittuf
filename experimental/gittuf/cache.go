// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

// PopulateCache scans the repository's RSL and generates a persistent
// local-only cache of policy and attestation entries. This makes subsequent
// verifications faster.
func (r *Repository) PopulateCache() error { _ = "STUB: not implemented"; return nil }

// DeleteCache deletes the local persistent cache.
func (r *Repository) DeleteCache() error { _ = "STUB: not implemented"; return nil }
