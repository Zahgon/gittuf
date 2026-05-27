// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

const DefaultRemoteName = "origin"

type FetchOptions struct {
	Depth int
}

type FetchOption func(*FetchOptions)

func WithFetchDepth(depth int) FetchOption { _ = "STUB: not implemented"; return *new(FetchOption) }

func (r *Repository) PushRefSpec(remoteName string, refSpecs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) Push(remoteName string, refs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) FetchRefSpec(remoteName string, refSpecs []string, opts ...FetchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) Fetch(remoteName string, refs []string, fastForwardOnly bool, opts ...FetchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) FetchObject(remoteName string, objectID Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func CloneAndFetchRepository(remoteURL, dir, initialBranch string, refs []string, bare bool) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) CreateRemote(remoteName, remoteURL string) error {
	_ = "STUB: not implemented"
	return nil
}
