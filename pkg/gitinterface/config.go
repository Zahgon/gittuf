// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

// GetGitConfig reads the applicable Git config for a repository and returns
// it. The "keys" for each config are normalized to lowercase.
func (r *Repository) GetGitConfig() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetGitConfig sets the specified key to the value locally for a repository.
func (r *Repository) SetGitConfig(key, value string) error { _ = "STUB: not implemented"; return nil }
