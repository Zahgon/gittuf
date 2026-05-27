package dsse

import (
	"google.golang.org/protobuf/types/known/structpb"
)

/*
Envelope captures an envelope as described by the DSSE specification. See here:
https://github.com/secure-systems-lab/dsse/blob/master/envelope.md
*/
type Envelope struct {
	PayloadType string      `json:"payloadType"`
	Payload     string      `json:"payload"`
	Signatures  []Signature `json:"signatures"`
}

/*
DecodeB64Payload returns the serialized body, decoded from the envelope's
payload field. A flexible decoder is used, first trying standard base64, then
URL-encoded base64.
*/
func (e *Envelope) DecodeB64Payload() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		/*
		   Signature represents a generic in-toto signature that contains the identifier
		   of the key which was used to create the signature.
		   The used signature scheme has to be agreed upon by the signer and verifer
		   out of band.
		   The signature is a base64 encoding of the raw bytes from the signature
		   algorithm.
		*/nil
}

type Signature struct {
	KeyID     string     `json:"keyid"`
	Sig       string     `json:"sig"`
	Extension *Extension `json:"extension,omitempty"`
}

type Extension struct {
	Kind string           `json:"kind"`
	Ext  *structpb.Struct `json:"ext"`
}

/*
PAE implementes the DSSE Pre-Authentic Encoding
https://github.com/secure-systems-lab/dsse/blob/master/protocol.md#signature-definition
*/
func PAE(payloadType string, payload []byte) []byte { _ = "STUB: not implemented"; return nil }

/*
Both standard and url encoding are allowed:
https://github.com/secure-systems-lab/dsse/blob/master/envelope.md
*/
func b64Decode(s string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
