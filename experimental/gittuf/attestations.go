// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"context"
	"errors"

	attestopts "github.com/gittuf/gittuf/experimental/gittuf/options/attest"
	githubopts "github.com/gittuf/gittuf/experimental/gittuf/options/github"
	"github.com/gittuf/gittuf/internal/attestations"
	"github.com/gittuf/gittuf/internal/attestations/github"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	gogithub "github.com/google/go-github/v61/github"
)

var (
	ErrNotSigningKey = errors.New("expected signing key")
	ErrNoGitHubToken = errors.New("authentication token for GitHub API not provided")
)

// ApplyAttestations records the state of the attestations reference and syncs
// it with the specified remote, making the attestation available at the
// synchronization point.
func (r *Repository) ApplyAttestations(ctx context.Context, remoteName string, localOnly, signRSLEntry bool) error {
	_ = "STUB: not implemented"
	return nil
}

// AddReferenceAuthorization adds a reference authorization attestation to the
// repository for the specified target ref. The from ID is identified using the
// last RSL entry for the target ref. The to ID is that of the expected Git tree
// created by merging the feature ref into the target ref. The commit used to
// calculate the merge tree ID is identified using the RSL for the feature ref.
func (r *Repository) AddReferenceAuthorization(ctx context.Context, signer sslibdsse.SignerVerifier, targetRef, featureRef string, signCommit bool, opts ...attestopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't have an RSL entry for the feature ref to use to approve the
// merge

// for tags, the toID is the commitID the tag will point to

// Does a reference authorization already exist for the parameters?

// Create a new reference authorization and embed in env

// RemoveReferenceAuthorization removes a previously issued authorization for
// the specified parameters. The issuer of the authorization is identified using
// their key.
func (r *Repository) RemoveReferenceAuthorization(ctx context.Context, signer sslibdsse.SignerVerifier, targetRef, fromID, toID string, signCommit bool, opts ...attestopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure only the key that created a reference authorization can remove it

// No reference authorization at all

// This handles cases where the envelope may unintentionally have
// multiple signatures from the same key

// No signatures, we can remove the ReferenceAuthorization altogether

// We still have other signatures, so set the ReferenceAuthorization
// envelope

// AddGitHubPullRequestAttestationForCommit identifies the pull request for a
// specified commit ID and triggers AddGitHubPullRequestAttestationForNumber for
// that pull request. The source of the authentication token for the GitHub API
// can be passed in as an option. If a source is not provided, the token is read
// from the GITHUB_TOKEN environment variable. A custom GitHub instance can be
// specified via opts.
func (r *Repository) AddGitHubPullRequestAttestationForCommit(ctx context.Context, signer sslibdsse.SignerVerifier, owner, repository, commitID, baseBranch string, signCommit bool, opts ...githubopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// pullRequest.Merged is not set on this endpoint for some reason

// AddGitHubPullRequestAttestationForNumber wraps the API response for the
// specified pull request in an in-toto attestation. `pullRequestID` must be the
// number of the pull request. The source of the authentication token for the
// GitHub API can be passed in as an option. If it is not passed in, the token
// is read from the GITHUB_TOKEN environment variable. A custom GitHub instance
// can be specified via opts.
func (r *Repository) AddGitHubPullRequestAttestationForNumber(ctx context.Context, signer sslibdsse.SignerVerifier, owner, repository string, pullRequestNumber int, signCommit bool, opts ...githubopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AddGitHubPullRequestApprover adds a GitHub pull request approval attestation
// for the specified parameters. If an attestation already exists, the specified
// approver is added to the existing attestation's predicate and it is re-signed
// and stored in the repository. To find the review information, the GitHub API
// is used and the source for the authentication token for the API is passed in
// as an option.  If the source is not passed in, the token is read from the
// GITHUB_TOKEN environment variable. A custom GitHub instance can be specified
// via opts.
func (r *Repository) AddGitHubPullRequestApprover(ctx context.Context, signer sslibdsse.SignerVerifier, owner, repository string, pullRequestNumber int, reviewID int64, approver string, signCommit bool, opts ...githubopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: make this configurable, check appName's key matches signer

// TODO: if the helper above has an indexPath, we can directly load that blob, simplifying the logic here

// Create a new GitHub pull request approval attestation

// Update existing statement's predicate and create new env

// DismissGitHubPullRequestApprover removes an approver from the GitHub pull
// request approval attestation for the specified parameters. A custom GitHub
// instance can be specified via opts.
func (r *Repository) DismissGitHubPullRequestApprover(ctx context.Context, signer sslibdsse.SignerVerifier, reviewID int64, dismissedApprover string, signCommit bool, opts ...githubopts.Option) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't hit the GitHub API for this flow, so no need to check the token
// option

// Update existing statement's predicate and create new env

func (r *Repository) addGitHubPullRequestAttestation(ctx context.Context, signer sslibdsse.SignerVerifier, githubBaseURL, owner, repository string, pullRequest *gogithub.PullRequest, createRSLEntry, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// not yet merged

// merged

func getGitHubPullRequestApprovalPredicateFromEnvelope(env *sslibdsse.Envelope) (github.PullRequestApprovalAttestation, error) {
	_ = "STUB: not implemented"
	return *new(github.PullRequestApprovalAttestation), nil
}

// TODO: support multiple versions here

// tmpGitHubPullRequestApprovalStatement is essentially a definition of
// in-toto's v1 Statement. The difference is that we fix the predicate to be
// the GitHub pull request approval type, making unmarshalling easier.

func indexPathToComponents(indexPath string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// remove last two items which are from-to and system

// reconstruct ref

func (r *Repository) getGitHubPullRequestReviewDetails(ctx context.Context, currentAttestations *attestations.Attestations, githubBaseURL, githubToken, owner, repository string, pullRequestNumber int, reviewID int64, useGitHubAPI bool) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// Compute details for review, this is when the review is first created as
// other times we use the existing indexPath for the reviewID
// Note: there's the potential for a TOCTOU issue here, we may query the
// repo after things have moved in either branch.

// testing validity of reviewID for the pull request in question

// We have to fetch the base ref name via the API response, we don't have a choice

// For the rest we have a choice. If useGitHubAPI is true, pull all
// information from the API. Otherwise, compute it.

// current tip of base ref

// Check the RSL instead for from ID.

// The API seems not to be outdated for the head commit, but we should
// monitor this. For now, use the API to get the head commit ID to avoid
// having to fetch the actual ref...

// Now compute the merge tree ID

// getGitHubClient creates a client to interact with a GitHub instance. If a
// base URL other than https://github.com is supplied, the client is configured
// to interact with the specified enterprise instance.
func getGitHubClient(baseURL, githubToken string) (*gogithub.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
