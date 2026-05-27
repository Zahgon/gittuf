// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package common //nolint:revive

import (
	"crypto/x509"
	"errors"
)

var (
	ErrSignatureVerificationFailed = errors.New("failed to verify signature")
	ErrNotPrivateKey               = errors.New("loaded key is not a private key")
	ErrUnknownKeyType              = errors.New("unknown key type")
	ErrInvalidThreshold            = errors.New("threshold is either less than 1 or greater than number of provided public keys")
)

// LoadCertsFromPath opens the file at the specified path and parses the
// certificates present in PEM form. This is similar to a helper in
// https://github.com/sigstore/sigstore and is used in gittuf's sigstore signing
// and verification flows.
func LoadCertsFromPath(path string) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
