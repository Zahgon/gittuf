// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package ssh

import (
	"context"
	"crypto"
	"testing"

	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
	"golang.org/x/crypto/ssh"
)

const (
	SigNamespace = "git"
	KeyType      = "ssh"
)

// Verifier is a dsse.Verifier implementation for SSH keys.
type Verifier struct {
	keyID  string
	sshKey ssh.PublicKey
}

// Verify implements the dsse.Verifier.Verify interface for SSH keys.
func (v *Verifier) Verify(_ context.Context, data []byte, sig []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ssh-keygen uses sha512 to sign with **any*** key

// KeyID implements the dsse.Verifier.KeyID interface for SSH keys.
// FIXME: consider removing error in interface; a dsse.Verifier needs a keyid
func (v *Verifier) KeyID() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// Public implements the dsse.Verifier.Public interface for SSH keys.
		// FIXME: consider removing in interface, "Verify()" is all that's needed
		nil
}

func (v *Verifier) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}

func (v *Verifier) MetadataKey() *signerverifier.SSLibKey { _ = "STUB: not implemented"; return nil }

// Signer is a dsse.Signer implementation for SSH keys.
type Signer struct {
	Path string
	*Verifier
}

// Sign implements the dsse.Signer.Sign interface for SSH keys.
// It signs using "s.Path" to a public or private, encrypted or plaintext, rsa,
// ecdsa or ed25519 key file in a format supported by "ssh-keygen". This aligns
// with the git "user.signingKey" option.
// https://git-scm.com/docs/git-config#Documentation/git-config.txt-usersigningKey
func (s *Signer) Sign(_ context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

// NewKeyFromFile imports an ssh SSlibKey from the passed path.
// The path can point to a public or private, encrypted or plaintext, rsa,
// ecdsa or ed25519 key file in a format supported by "ssh-keygen". This aligns
// with the git "user.signingKey" option.
// https://git-scm.com/docs/git-config#Documentation/git-config.txt-usersigningKey
func NewKeyFromFile(path string) (*signerverifier.SSLibKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewKeyFromBytes returns an ssh SSLibKey from the passed bytes. It's meant to
// be used for tests as that's when we directly deal with key bytes.
func NewKeyFromBytes(t *testing.T, keyB []byte) *signerverifier.SSLibKey {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// NewVerifierFromKey creates a new Verifier from SSlibKey of type ssh.
func NewVerifierFromKey(key *signerverifier.SSLibKey) (*Verifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSignerFromFile creates an SSH signer from the passed path.
func NewSignerFromFile(path string) (*Signer, error) { _ = "STUB: not implemented"; return nil, nil }

// parseSSH2Body parses a base64-encoded SSH2 wire format key.
func parseSSH2Body(body string) (ssh.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(ssh.PublicKey), nil
}

// parseSSH2Key parses a SSH2 public key as defined in RFC4716 (section 3.)
// NOTE:
// - only supports "\n" as line termination character
// - does not validate line length, or header tag or value format
// - discards headers
func parseSSH2Key(data string) (ssh.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(ssh.PublicKey), nil
}

// Normalize and trim newlines

// Strip begin and end markers

// Strip headers

// Skip i==1, first line can not be a continued line

// Parse key material

func newSSHKey(key ssh.PublicKey, keyID string) *signerverifier.SSLibKey {
	_ = "STUB: not implemented"
	return nil
}
