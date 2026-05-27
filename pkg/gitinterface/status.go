// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"errors"
)

// See https://git-scm.com/docs/git-status#_porcelain_format_version_1.

var (
	ErrInvalidStatusCodeLength = errors.New("status code string must be of length 1")
	ErrInvalidStatusCode       = errors.New("status code string is unrecognized")
)

type StatusCode uint

const (
	StatusCodeUnmodified StatusCode = iota + 1 // we use 0 as error code
	StatusCodeModified
	StatusCodeTypeChanged
	StatusCodeAdded
	StatusCodeDeleted
	StatusCodeRenamed
	StatusCodeCopied
	StatusCodeUpdatedUnmerged
	StatusCodeUntracked
	StatusCodeIgnored
)

func (s StatusCode) String() string { _ = "STUB: not implemented"; return "" }

// is this actually a space or empty string?

func NewStatusCodeFromByte(s byte) (StatusCode, error) {
	_ = "STUB: not implemented"
	return *new(StatusCode), nil
}

type FileStatus struct {
	X StatusCode
	Y StatusCode
}

func (f *FileStatus) Untracked() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) Status() (map[string]FileStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: this doesn't support detached git dir

//nolint:errcheck

// `git status --porcelain=1 -z` emits NUL-separated tokens.
// For rename/copy records, the source path is emitted as an additional
// token after the main status token.

// first two characters are status codes, find the corresponding
// statuses

// Note: we identify the status after inspecting the path so we can
// provide better error messages

// then, we have a single space followed by the path, ignore space and
// read in the rest as the filepath

// After splitting on NUL, rename/copy records have an additional token
// for the source path immediately after the main token.
