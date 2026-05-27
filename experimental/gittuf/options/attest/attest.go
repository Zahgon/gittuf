// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package attest

type Options struct {
	CreateRSLEntry bool
}

type Option func(o *Options)

func WithRSLEntry() Option { _ = "STUB: not implemented"; return *new(Option) }
