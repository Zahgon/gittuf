// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package sigstore

import (
	"context"
	"crypto"

	signeropts "github.com/gittuf/gittuf/internal/signerverifier/sigstore/options/signer"
	verifieropts "github.com/gittuf/gittuf/internal/signerverifier/sigstore/options/verifier"
	"github.com/secure-systems-lab/go-securesystemslib/signerverifier"
	"github.com/sigstore/sigstore-go/pkg/sign"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	KeyType   = "sigstore-oidc"
	KeyScheme = "fulcio"

	ExtensionMimeType = "application/vnd.dev.sigstore.verificationmaterial;version=0.3"

	GitConfigIssuer      = "gitsign.issuer"
	GitConfigClientID    = "gitsign.clientid"
	GitConfigFulcio      = "gitsign.fulcio"
	GitConfigRekor       = "gitsign.rekor"
	GitConfigRedirectURL = "gitsign.redirecturl"

	EnvSigstoreRootFile = "SIGSTORE_ROOT_FILE"

	sigstoreBundleMimeType = "application/vnd.dev.sigstore.bundle+json;version=0.3"
)

type Verifier struct {
	rekorURL string
	issuer   string
	identity string
	ext      *structpb.Struct
}

func NewVerifierFromIdentityAndIssuer(identity, issuer string, opts ...verifieropts.Option) *Verifier {
	_ = "STUB: not implemented"
	return nil
}

func (v *Verifier) Verify(_ context.Context, data, sig []byte) error {
	_ = "STUB: not implemented"
	// data is PAE(envelope)
	// sig is raw sigBytes
	// extension is set in the verifier
	return nil
}

// create protobuf bundle

func (v *Verifier) KeyID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (v *Verifier) Public() crypto.PublicKey {
	_ = "STUB: not implemented"
	// TODO
	return *new(crypto.PublicKey)
}

func (v *Verifier) SetExtension(ext *structpb.Struct) { _ = "STUB: not implemented"; return }

func (v *Verifier) ExpectedExtensionKind() string {
	_ = "STUB: not implemented"
	// TODO: versioning?
	return ""
}

type Signer struct {
	issuerURL   string
	clientID    string
	redirectURL string
	fulcioURL   string
	rekorURL    string
	token       string
	*Verifier
}

func NewSigner(opts ...signeropts.Option) *Signer { _ = "STUB: not implemented"; return nil }

func (s *Signer) Sign(_ context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: support private sigstore by reading config

// We reuse the token if it's already been fetched once for this signer
// object
// getIDToken also populates the Verifier's identity and issuer pieces

// TODO: TSA support?

func (s *Signer) KeyID() (string, error) {
	_ = "STUB: not implemented"
	// verifier can't return error
	return "", nil
}

//nolint:errcheck

// verifier.identity and verifier.issuer are empty resulting in this
// return value

// getIDToken will populate verifier

// MetadataKey returns the securesystemslib representation of the key, used for
// its representation in gittuf metadata.
func (s *Signer) MetadataKey() (*signerverifier.SSLibKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Signer) getIDToken() (string, error) {
	_ = "STUB: not implemented"

	// TODO: support client secret?
	return "", nil
}

// Set identity and issuer pieces

func (s *Signer) getFulcioInstance() *sign.Fulcio { _ = "STUB: not implemented"; return nil }

func (s *Signer) getRekorInstance() *sign.Rekor { _ = "STUB: not implemented"; return nil }
