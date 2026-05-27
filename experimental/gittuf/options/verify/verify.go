// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package verify

type Options struct {
	RefNameOverride string
	LatestOnly      bool
}

type Option func(o *Options)

func WithOverrideRefName(refNameOverride string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLatestOnly() Option { _ = "STUB: not implemented"; return *new(Option) }
