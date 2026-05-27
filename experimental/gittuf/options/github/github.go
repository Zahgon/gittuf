// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package github

import (
	"context"
)

const (
	DefaultGitHubBaseURL = "https://github.com"

	githubTokenEnvKey = "GITHUB_TOKEN" //nolint:gosec
)

// TokenSource is a lightweight interface that can be used to fetch a GitHub
// token.
type TokenSource interface {
	Token(context.Context) (string, error)
}

type Options struct {
	GitHubTokenSource TokenSource
	GitHubBaseURL     string
	CreateRSLEntry    bool
	UseGitHubAPI      bool
}

var DefaultOptions = &Options{
	GitHubBaseURL:     DefaultGitHubBaseURL,
	GitHubTokenSource: &TokenSourceEnvironment{},
}

type Option func(o *Options)

// WithGitHubTokenSource can be used to specify an authentication token source
// to fetch a token to use the GitHub API.
func WithGitHubTokenSource(tokenSource TokenSource) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGitHubBaseURL can be used to specify a custom GitHub instance, such as an
// on-premises GitHub Enterprise Server.
func WithGitHubBaseURL(baseURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRSLEntry() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithUseGitHubAPI() Option { _ = "STUB: not implemented"; return *new(Option) }

// TokenSourceEnvironment reads the GitHub API token from the GITHUB_TOKEN
// environment variable. It implements the TokenSource interface.
type TokenSourceEnvironment struct{}

func (t *TokenSourceEnvironment) Token(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
