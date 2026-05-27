// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gpg

import (
	"context"
	"crypto"
	"testing"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
)

const (
	KeyType = "gpg"

	defaultGPGProgram = "gpg"
)

// Verifier is a dsse.Verifier implementation for GPG keys.
type Verifier struct {
	metadataKey *signerverifier.SSLibKey
	keyID       string
	entity      *openpgp.Entity
}

// KeyID implements the dsse.Verifier.KeyID interface for GPG keys.
// FIXME: consider removing error in interface; a dsse.Verifier needs a keyid
func (v *Verifier) KeyID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Public implements the dsse.Verifier.Public interface for GPG keys.
		// FIXME: consider removing in interface, "Verify()" is all that's needed
		nil
}

func (v *Verifier) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (v *Verifier) MetadataKey() *signerverifier.SSLibKey { _ = "STUB: not implemented"; return nil }

// Verify implements the dsse.Verifier.Verify interface for GPG keys.
func (v *Verifier) Verify(_ context.Context, data []byte, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// NewVerifierFromKey creates a new verifier from SSLibKey of type GPG.
func NewVerifierFromKey(key *signerverifier.SSLibKey) (*Verifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Signer struct {
	*Verifier
	program string
}

func (s *Signer) KeyID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Sign implements the dsse.Signer.Sign interface for GPG keys.
		nil
}

func (s *Signer) Sign(_ context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

func NewSignerFromKeyID(keyID string, opts ...SignerOption) (*Signer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LoadGPGKeyFromBytes returns a signerverifier.SSLibKey for a GPG / PGP key passed in as
// armored bytes. The returned signerverifier.SSLibKey uses the primary key's fingerprint as the
// key ID.
func LoadGPGKeyFromBytes(contents []byte) (*signerverifier.SSLibKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check if this is correct for subkeys

// TODO: this should use the underlying key algorithm

func getPublicKeyForKeyID(keyID, program string) (*signerverifier.SSLibKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// SetupTestGPGHomeDir is a test helper used only to prepare a temporary GPG
// home dir with the specified keys added in.
func SetupTestGPGHomeDir(t *testing.T, privateKeyBytes ...[]byte) {
	_ = "STUB: not implemented"
	return
}

// We use os.MkdirTemp because t.TempDir can result in a path
// that's too long for socket files, used for gpg-agent.

//nolint:errcheck

type SignerOptions struct {
	program string
}

type SignerOption func(*SignerOptions)

func WithGPGProgram(program string) SignerOption {
	_ = "STUB: not implemented"
	return *new(SignerOption)
}
