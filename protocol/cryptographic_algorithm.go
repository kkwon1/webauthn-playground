package protocol

// COSEAlgorithmIdentifier From §5.10.5. A number identifying a cryptographic algorithm. The algorithm
// identifiers SHOULD be values registered in the IANA COSE Algorithms registry
// [https://www.w3.org/TR/webauthn/#biblio-iana-cose-algs-reg], for instance, -7 for "ES256"
//  and -257 for "RS256".
type COSEAlgorithmIdentifier int

// This enumeration defines the valid credential types.
// It is an extension point; values can be added to it in the future, as
// more credential types are defined. The values of this enumeration are used
// for versioning the Authentication Assertion and attestation structures according
// to the type of the authenticator.
// See §5.10.3. Credential Descriptor https://www.w3.org/TR/webauthn/#credentialType
type CredentialType string

// CredentialParameter is the credential type and algorithm
// that the relying party wants the authenticator to create
type CredentialParameter struct {
	Type      CredentialType   		    `json:"type"`
	Algorithm COSEAlgorithmIdentifier `json:"alg"`
}

const (
	// AlgES256 ECDSA with SHA-256
	AlgES256 COSEAlgorithmIdentifier = -7
	// AlgES384 ECDSA with SHA-384
	AlgES384 COSEAlgorithmIdentifier = -35
	// AlgES512 ECDSA with SHA-512
	AlgES512 COSEAlgorithmIdentifier = -36
	// AlgRS1 RSASSA-PKCS1-v1_5 with SHA-1
	AlgRS1 COSEAlgorithmIdentifier = -65535
	// AlgRS256 RSASSA-PKCS1-v1_5 with SHA-256
	AlgRS256 COSEAlgorithmIdentifier = -257
	// AlgRS384 RSASSA-PKCS1-v1_5 with SHA-384
	AlgRS384 COSEAlgorithmIdentifier = -258
	// AlgRS512 RSASSA-PKCS1-v1_5 with SHA-512
	AlgRS512 COSEAlgorithmIdentifier = -259
	// AlgPS256 RSASSA-PSS with SHA-256
	AlgPS256 COSEAlgorithmIdentifier = -37
	// AlgPS384 RSASSA-PSS with SHA-384
	AlgPS384 COSEAlgorithmIdentifier = -38
	// AlgPS512 RSASSA-PSS with SHA-512
	AlgPS512 COSEAlgorithmIdentifier = -39
	// AlgEdDSA EdDSA
	AlgEdDSA COSEAlgorithmIdentifier = -8
)

const (
	// PublicKeyCredentialType - Currently one credential type is defined, namely "public-key".
	PublicKeyCredentialType CredentialType = "public-key"
)