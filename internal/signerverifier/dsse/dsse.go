// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package dsse

import (
	"context"

	"github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
)

const PayloadType = "application/vnd.gittuf+json"

// CreateEnvelope is an opinionated interface to create a DSSE envelope. It
// accepts instances of tuf.RootMetadata, tuf.TargetsMetadata, etc. and marshals
// the input prior to storing it as the envelope's payload.
func CreateEnvelope(v any) (*dsse.Envelope, error) { _ = "STUB: not implemented"; return nil, nil }

// SignEnvelope is an opinionated API to sign DSSE envelopes. It's opinionated
// because it assumes the payload is Base 64 encoded, which is the expectation
// for gittuf metadata. If one or more signatures from the provided signing key
// already exist, they are all removed in favor of the new signature from that
// key.
func SignEnvelope(ctx context.Context, envelope *dsse.Envelope, signer dsse.Signer) (*dsse.Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unpack the bundle to get the signature + verification material
// Set extension in the signature object

// Preserve signatures that aren't from signer

// Attach new signature from signer

// Replace existing list of signatures with new signatures in envelope

// VerifyEnvelope verifies a DSSE envelope against an expected threshold using
// a slice of verifiers passed into it. Threshold indicates the number of
// providers that must validate the envelope.
func VerifyEnvelope(ctx context.Context, envelope *dsse.Envelope, verifiers []dsse.Verifier, threshold int) ([]dsse.AcceptedKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We verify with threshold == 1 because we want control over the threshold
// checks: we get all the verified keys back
