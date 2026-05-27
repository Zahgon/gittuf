// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

// ReadBlob returns the contents of the blob referenced by blobID.
func (r *Repository) ReadBlob(blobID Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteBlob creates a blob object with the specified contents and returns the
// ID of the resultant blob.
func (r *Repository) WriteBlob(contents []byte) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}
