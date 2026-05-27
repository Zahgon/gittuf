// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package rsl

type RecordOptions struct {
	RefNameOverride       string
	RemoteName            string
	LocalOnly             bool
	SkipCheckForDuplicate bool
}

type RecordOption func(o *RecordOptions)

func WithOverrideRefName(refNameOverride string) RecordOption {
	_ = "STUB: not implemented"
	return *new(RecordOption)
}

// WithSkipCheckForDuplicateEntry indicates that the RSL entry creation must not
// check if the latest entry for the reference has the same target ID.
func WithSkipCheckForDuplicateEntry() RecordOption {
	_ = "STUB: not implemented"
	return *new(RecordOption)
}

func WithRecordRemote(remoteName string) RecordOption {
	_ = "STUB: not implemented"
	return *new(RecordOption)
}

func WithRecordLocalOnly() RecordOption { _ = "STUB: not implemented"; return *new(RecordOption) }

type AnnotateOptions struct {
	RemoteName string
	LocalOnly  bool
}

type AnnotateOption func(o *AnnotateOptions)

func WithAnnotateRemote(remoteName string) AnnotateOption {
	_ = "STUB: not implemented"
	return *new(AnnotateOption)
}

func WithAnnotateLocalOnly() AnnotateOption { _ = "STUB: not implemented"; return *new(AnnotateOption) }
