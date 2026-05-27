// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package verifier

const (
	defaultRekorURL = "https://rekor.sigstore.dev"
)

type Options struct {
	RekorURL string
}

var DefaultOptions = &Options{
	RekorURL: defaultRekorURL,
}

type Option func(o *Options)

func WithRekorURL(rekorURL string) Option { _ = "STUB: not implemented"; return *new(Option) }
