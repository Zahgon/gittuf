// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"errors"
)

type ObjectType uint

const (
	BlobObjectType ObjectType = iota + 1
	TreeObjectType
	CommitObjectType
	TagObjectType
)

var ErrInvalidObjectType = errors.New("unknown Git object type")

// HasObject returns true if an object with the specified Git ID exists in the
// repository.
func (r *Repository) HasObject(objectID Hash) bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetObjectType(objectID Hash) (ObjectType, error) {
	_ = "STUB: not implemented"
	return *new(ObjectType), nil
}

// GetObjectSize returns the size of the object with the specified Git ID.
func (r *Repository) GetObjectSize(objectID Hash) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
