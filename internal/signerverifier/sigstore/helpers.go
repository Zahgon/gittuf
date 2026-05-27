// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package sigstore

const fulcioConfigurationEndpoint = "/api/v2/configuration"

func parseTokenForIdentityAndIssuer(token, fulcioURL string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

//nolint:gosec

// Per the Fulcio spec, the subject domain is added after a '!'

type idToken struct {
	Issuer          string           `json:"iss"`
	Subject         string           `json:"sub"`
	Email           string           `json:"email"`
	EmailVerified   stringAsBool     `json:"email_verified"`
	FederatedClaims *federatedClaims `json:"federated_claims"`
}

type stringAsBool bool

func (sb *stringAsBool) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type federatedClaims struct {
	ConnectorID string `json:"connector_id"`
}

func issuerFromToken(tok *idToken) string { _ = "STUB: not implemented"; return "" }

func subjectFromToken(tok *idToken) string { _ = "STUB: not implemented"; return "" }
