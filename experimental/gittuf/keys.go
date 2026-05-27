// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"errors"

	"github.com/gittuf/gittuf/internal/signerverifier/gpg"
	sigstoresigneropts "github.com/gittuf/gittuf/internal/signerverifier/sigstore/options/signer"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
)

const (
	GPGKeyPrefix = "gpg:"
	FulcioPrefix = "fulcio:"
)

type signingMethod int

const (
	signingMethodGPG signingMethod = iota
	signingMethodSSH
	signingMethodX509
)

var (
	ErrUnsupportedSigningMethod = errors.New("unsupported signing method specified in Git configuration")
	ErrSigningKeyNotSpecified   = errors.New("signing key not specified in Git configuration")
	ErrUnsupportedX509Method    = errors.New("unsupported X509 certificate specified in Git configuration")
)

// LoadPublicKey returns a signerverifier.SSLibKey object for a PGP / Sigstore
// Fulcio / SSH (on-disk) key for use in gittuf metadata.
func LoadPublicKey(keyRef string) (tuf.Principal, error) {
	_ = "STUB: not implemented"
	return *new(tuf.Principal), nil
}

// LoadPublicKeyFromGitConfig loads a public key as with LoadPublicKey above,
// but from the key specified in the Git configuration of the target repository.
func LoadPublicKeyFromGitConfig(repo *Repository) (tuf.Principal, error) {
	_ = "STUB: not implemented"
	return *new(tuf.Principal), nil
}

// Attempt to determine what type of key is specified by the user's Git
// config

// GPG is assumed if "gpg" is specified, or if nothing is specified

// If some other format specified, return error

// Get the path to the signing key, required if using an SSH or GPG key

// GPG
// Load a GPG signer from the specified key

// SSH
// Load an SSH signer from the specified key

// X.509
// We only support sigstore X.509, so check that gitsign is specified

// gitsign

// LoadSigner loads a metadata signer for the specified key bytes. The signer
// must be for a GPG key (in which case the `key` is the GPG key ID), an SSH key
// (in which case the `key` is a path to the private key) or for signing with
// Sigstore (where `key` has a prefix `fulcio:`). If no key ID is specified,
// this function calls LoadSignerFromGitConfig.
func LoadSigner(repo *Repository, key string) (sslibdsse.SignerVerifier, error) {
	_ = "STUB: not implemented"
	return *new(sslibdsse.SignerVerifier), nil
}

// LoadSignerFromGitConfig loads a metadata signer for the signing key specified
// in the Git configuration of the target repository.
func LoadSignerFromGitConfig(repo *Repository) (sslibdsse.SignerVerifier, error) {
	_ = "STUB: not implemented"
	return *new(sslibdsse.SignerVerifier), nil
}

// Attempt to determine what type of key is specified by the user's Git
// config

// GPG is assumed if "gpg" is specified, or if nothing is specified

// If some other format specified, return error

// Get the path to the signing key, required if using an SSH or GPG key

// GPG
// Load a GPG signer from the specified key

// SSH
// Load an SSH signer from the specified key

// X.509
// We only support sigstore X.509, so check that gitsign is specified

// gitsign

func getGPGOptions(config map[string]string) []gpg.SignerOption {
	_ = "STUB: not implemented"
	return nil

	// Parse relevant gpg.<name> config values
}

func getSigstoreOptions(config map[string]string) []sigstoresigneropts.Option {
	_ = "STUB: not implemented"
	return nil
}

// Parse relevant gitsign.<> config values
