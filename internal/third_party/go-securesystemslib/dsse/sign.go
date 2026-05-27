/*
Package dsse implements the Dead Simple Signing Envelope (DSSE)
https://github.com/secure-systems-lab/dsse
*/
package dsse

import (
	"context"
	"errors"
)

// ErrNoSigners indicates that no signer was provided.
var ErrNoSigners = errors.New("no signers provided")

// EnvelopeSigner creates signed Envelopes.
type EnvelopeSigner struct {
	providers []Signer
}

/*
NewEnvelopeSigner creates an EnvelopeSigner that uses 1+ Signer algorithms to
sign the data.
*/
func NewEnvelopeSigner(p ...Signer) (*EnvelopeSigner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
NewMultiEnvelopeSigner creates an EnvelopeSigner that uses 1+ Signer
algorithms to sign the data. The threshold parameter is legacy and is ignored.

Deprecated: This function simply calls NewEnvelopeSigner, and that function should
be preferred.
*/
func NewMultiEnvelopeSigner(threshold int, p ...Signer) (*EnvelopeSigner, error) {
	_ = "STUB: not implemented"
	return nil, nil

	/*
	   SignPayload signs a payload and payload type according to DSSE.
	   Returned is an envelope as defined here:
	   https://github.com/secure-systems-lab/dsse/blob/master/envelope.md
	   One signature will be added for each Signer in the EnvelopeSigner.
	*/
}

func (es *EnvelopeSigner) SignPayload(ctx context.Context, payloadType string, body []byte) (*Envelope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
