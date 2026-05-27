// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"context"

	"github.com/gittuf/gittuf/internal/common/set"
	sslibdsse "github.com/gittuf/gittuf/internal/third_party/go-securesystemslib/dsse"
	"github.com/gittuf/gittuf/internal/tuf"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

type SignatureVerifier struct {
	repository         *gitinterface.Repository
	name               string
	principals         []tuf.Principal
	threshold          int
	verifyExhaustively bool // verifyExhaustively checks all possible signatures and returns all matched principals, even if threshold is already met
}

func (v *SignatureVerifier) Name() string { _ = "STUB: not implemented"; return "" }

func (v *SignatureVerifier) Threshold() int { _ = "STUB: not implemented"; return 0 }

func (v *SignatureVerifier) TrustedPrincipalIDs() *set.Set[string] {
	_ = "STUB: not implemented"
	return nil
}

// Verify is used to check for a threshold of signatures using the verifier. The
// threshold of signatures may be met using a combination of at most one Git
// signature and signatures embedded in a DSSE envelope. Verify does not inspect
// the envelope's payload, but instead only verifies the signatures. The caller
// must ensure the validity of the envelope's contents.
func (v *SignatureVerifier) Verify(ctx context.Context, gitObjectID gitinterface.Hash, env *sslibdsse.Envelope) (*set.Set[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// usedPrincipalIDs is ultimately returned to track the set of principals
// who have been authenticated

// usedKeyIDs is tracked to ensure a key isn't duplicated between two
// principals, allowing two principals to meet a threshold using the same
// key

// gitObjectVerified is set to true if the gitObjectID's signature is
// verified

// First, verify the gitObject's signature if one is presented

// there are multiple keys we must try

// Signature verification succeeded

// No need to try the other keys for this principal, break

// TODO: this should be removed once we have unified signing
// methods across metadata and git signatures

// No need to try other principals, break

// If we don't have to verify exhaustively and threshold is 1 and the Git
// signature is verified, we can return

// Second, verify signatures on the envelope

// We have to verify the envelope independently for each principal
// trusted in the verifier as a principal may have multiple keys
// associated with them.

// Do not verify using this principal as they were verified for
// the Git signature

// this key has been encountered before, possibly because
// another Principal included this key

// We have the principal's verifiers: use that to verify the envelope

// TODO: remove this when we have signing method unification
// across git and dsse

// We set threshold to 1 as we only need one of the keys for this
// principal to be matched. If more than one key is matched and
// returned in acceptedKeys, we count this only once towards the
// principal and therefore the verifier's threshold. However, for
// safety, we count both keys. If two principals share keys, this
// can lead to a problem meeting thresholds. Arguably, they
// shouldn't be sharing keys, so this seems reasonable.

// Mark all accepted keys as used: this doesn't count towards
// the threshold directly, but if another principal has the same
// key, they may not be counted towards the threshold

// TODO: double check that this is okay!

// Return usedPrincipalIDs so the consumer can decide what to do with the
// principals that were used
