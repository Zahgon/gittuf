// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

// AddRemote adds a remote with the specified name and URL.
func (r *Repository) AddRemote(remoteName, url string) error { _ = "STUB: not implemented"; return nil }

// RemoveRemote removes the remote with the specified name.
func (r *Repository) RemoveRemote(remoteName string) error { _ = "STUB: not implemented"; return nil }

// GetRemoteURL gets the URL of the remote with the specified name.
func (r *Repository) GetRemoteURL(remoteName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
