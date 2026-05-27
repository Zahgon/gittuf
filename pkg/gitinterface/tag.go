// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"context"
	"errors"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
)

var (
	ErrTagAlreadyExists = errors.New("tag already exists")
)

// TagUsingSpecificKey creates a Git tag signed using the specified, PEM encoded
// SSH or GPG key. It is primarily intended for use with testing. As of now,
// gittuf is not expected to be used to create tags in developer workflows,
// though this may change with command compatibility.
func (r *Repository) TagUsingSpecificKey(target Hash, name, message string, signingKeyPEMBytes []byte) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// GetTagTarget returns the ID of the Git object a tag points to.
func (r *Repository) GetTagTarget(tagID Hash) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// verifyTagSignature verifies a signature for the specified tag using the
// provided public key.
func (r *Repository) verifyTagSignature(ctx context.Context, tagID Hash, key *signerverifier.SSLibKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) ensureIsTag(tagID Hash) error { _ = "STUB: not implemented"; return nil }

func getTagBytesWithoutSignature(tag *object.Tag) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
