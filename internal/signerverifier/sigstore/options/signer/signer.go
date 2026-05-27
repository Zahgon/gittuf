// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package signer

const (
	defaultIssuerURL = "https://oauth2.sigstore.dev/auth"
	defaultClientID  = "sigstore"
	defaultFulcioURL = "https://fulcio.sigstore.dev"
	defaultRekorURL  = "https://rekor.sigstore.dev"
)

type Options struct {
	IssuerURL   string
	ClientID    string
	RedirectURL string
	FulcioURL   string
	RekorURL    string
}

var DefaultOptions = &Options{
	IssuerURL: defaultIssuerURL,
	ClientID:  defaultClientID,
	FulcioURL: defaultFulcioURL,
	RekorURL:  defaultRekorURL,
}

type Option func(o *Options)

func WithIssuerURL(issuerURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClientID(clientID string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRedirectURL(redirectURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFulcioURL(fulcioURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRekorURL(rekorURL string) Option { _ = "STUB: not implemented"; return *new(Option) }
