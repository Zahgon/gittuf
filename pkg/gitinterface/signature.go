// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"context"
	"errors"

	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
)

const (
	rekorPublicGoodInstance           = "https://rekor.sigstore.dev"
	namespaceSSHSignature      string = "git"
	gpgPrivateKeyPEMHeader     string = "PGP PRIVATE KEY"
	opensshPrivateKeyPEMHeader string = "OPENSSH PRIVATE KEY"
	rsaPrivateKeyPEMHeader     string = "RSA PRIVATE KEY"
	genericPrivateKeyPEMHeader string = "PRIVATE KEY"
	signingFormatGPG           string = "gpg"
	signingFormatSSH           string = "ssh"
)

var (
	ErrNotCommitOrTag             = errors.New("invalid object type, expected commit or tag for signature verification")
	ErrSigningKeyNotSpecified     = errors.New("signing key not specified in git config")
	ErrUnknownSigningMethod       = errors.New("unknown signing method (not one of gpg, ssh, x509)")
	ErrIncorrectVerificationKey   = errors.New("incorrect key provided to verify signature")
	ErrVerifyingSigstoreSignature = errors.New("unable to verify Sigstore signature")
	ErrVerifyingSSHSignature      = errors.New("unable to verify SSH signature")
	ErrInvalidSignature           = errors.New("unable to parse signature / signature has unexpected header")
)

// CanSign inspects the Git configuration to determine if commit / tag signing
// is possible.
func (r *Repository) CanSign() error { _ = "STUB: not implemented"; return nil }

// Format is one of GPG, SSH, X509

// If format is GPG or X509, the signing key parameter is optional
// However, for SSH, the signing key must be set

// VerifySignature verifies the cryptographic signature associated with the
// specified object. The `objectID` must point to a Git commit or tag object.
func (r *Repository) VerifySignature(ctx context.Context, objectID Hash, key *signerverifier.SSLibKey) error {
	_ = "STUB: not implemented"
	return nil
}

func signGitObjectUsingKey(contents, pemKeyBytes []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// openpgp implements its own armor-decode method, pem.Decode considers
// the input invalid. We haven't tested if this is universal, so in case
// pem.Decode does succeed on a GPG key, we catch it below.

func signGitObjectUsingGPGKey(contents, pemKeyBytes []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func signGitObjectUsingSSHKey(contents, pemKeyBytes []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// verifyGitsignSignature handles the Sigstore-specific workflow involved in
// verifying commit or tag signatures issued by gitsign.
func verifyGitsignSignature(ctx context.Context, repo *Repository, key *signerverifier.SSLibKey, data, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Check git config to see if rekor server must be overridden

// gitsignRekor.NewWithOptions invokes cosign.GetRekorPubs which looks at
// the env var, so we don't have to do anything here

// cosign.GetCTLogPubs already looks at the env var, so we don't have to do
// anything here

// verifySSHKeySignature verifies Git signatures issued by SSH keys.
func verifySSHKeySignature(ctx context.Context, key *signerverifier.SSLibKey, data, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func getSigningMethod(gitConfig map[string]string) string { _ = "STUB: not implemented"; return "" }

// default to gpg

func getSigningKeyInfo(gitConfig map[string]string) string { _ = "STUB: not implemented"; return "" }
