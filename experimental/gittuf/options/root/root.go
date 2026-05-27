// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package root

type Options struct {
	RepositoryLocation string
	CreateRSLEntry     bool
}

type Option func(o *Options)

func WithRepositoryLocation(location string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRSLEntry() Option { _ = "STUB: not implemented"; return *new(Option) }
