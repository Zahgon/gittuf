// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package hooks

import (
	"errors"
)

var ErrRequiredOptionNotSet = errors.New("required option not set")

type Options struct {
	PrePush *PrePushOptions
}

type Option func(o *Options)

// WithPrePush can be used to specify arguments normally passed to Git pre-push
// hooks.
func WithPrePush(remoteName, remoteURL string, refSpecs []string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

type PrePushOptions struct {
	RemoteName string
	RemoteURL  string
	RefSpecs   []string
}

func (o *PrePushOptions) Validate() error { _ = "STUB: not implemented"; return nil }
