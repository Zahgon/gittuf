package dsse

import (
	"context"
	"crypto"
	"errors"
)

// ErrNoSignature indicates that an envelope did not contain any signatures.
var ErrNoSignature = errors.New("no signature found")

type EnvelopeVerifier struct {
	providers []Verifier
	threshold int
}

type AcceptedKey struct {
	Public crypto.PublicKey
	KeyID  string
	Sig    Signature
}

func (ev *EnvelopeVerifier) Verify(ctx context.Context, e *Envelope) ([]AcceptedKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decode payload (i.e serialized body)

// Generate PAE(payloadtype, serialized body)

// If *any* signature is found to be incorrect, it is skipped

// Loop over the providers.
// If provider and signature include key IDs but do not match skip.
// If a provider recognizes the key, we exit
// the loop and use the result.

// Verifiers that do not provide a keyid will be generated one using public.

// See https://github.com/in-toto/in-toto/pull/251

// Sanity if with some reflect magic this happens.

func NewEnvelopeVerifier(v ...Verifier) (*EnvelopeVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMultiEnvelopeVerifier(threshold int, p ...Verifier) (*EnvelopeVerifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SHA256KeyID(pub crypto.PublicKey) (string, error) {
	_ = "STUB: not implemented"
	// Generate public key fingerprint
	return "", nil
}

func removeIndex(v []Verifier, index int) []Verifier { _ = "STUB: not implemented"; return nil }
