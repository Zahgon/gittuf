// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package v01

import (
	gogithub "github.com/google/go-github/v61/github"
	ita "github.com/in-toto/attestation/go/v1"
)

const (
	PullRequestPredicateType = "https://gittuf.dev/github-pull-request/v0.1"

	digestGitCommitKey = "gitCommit"
)

func NewPullRequestAttestation(owner, repository string, pullRequestNumber int, commitID string, pullRequest *gogithub.PullRequest) (*ita.Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
