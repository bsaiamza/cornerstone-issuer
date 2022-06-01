package offer

type CredentialOfferBodyRequest struct {
	ConnectionID string `json:"connection_id"`
	IDNum     string `json:"id_no,omitempty"`
	GivenName              string `json:"givenName"`
	FamilyName             string `json:"familyName"`
	Gender                 string `json:"gender"`
	LprNumber              string `json:"lprNumber"`
	LprCategory            string `json:"lprCategory"`
	ResidentSince          string `json:"residentSince"`
	CommuterClassification string `json:"commuterClassification"`
	BirthDate              string `json:"birthDate"`
	BirthCountry           string `json:"birthCountry"`
	SchemaID  string `json:"schema_id,omitempty"`
	CredDefID string `json:"cred_def_id,omitempty"`
}

type CredentialOfferRequest struct {
	AutoIssue     bool           `json:"auto_issue"`
	AutoRemove     bool           `json:"auto_remove"`
	Comment     string           `json:"comment"`
	ConnectionID     string           `json:"connection_id"`
	CredentialPreview CredentialPreview `json:"credential_preview"`
	Filter         Filter         `json:"filter"`
	Trace          bool           `json:"trace"`
}

type CredentialPreview struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Filter struct {
	Indy    Indy    `json:"indy"`
	LdProof LdProof `json:"ld_proof"`
}

type Indy struct {
	CredDefID       string `json:"cred_def_id"`
	IssuerDid       string `json:"issuer_did"`
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDid string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

type LdProof struct {
	Credential Credential `json:"credential"`
	Options    Options    `json:"options"`
}

type Credential struct {
	Context      []string `json:"@context"`
	ID           string   `json:"id"`
	Type         []string `json:"type"`
	Issuer       string   `json:"issuer"`
	Identifier   string   `json:"identifier"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	IssuanceDate string   `json:"issuanceDate"`
	// ExpirationDate    string            `json:"expirationDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof             Proof             `json:"proof"`
}

type CredentialSubject struct {
	ID                     string   `json:"id"`
	Type                   []string `json:"type"`
	GivenName              string   `json:"givenName"`
	FamilyName             string   `json:"familyName"`
	Gender                 string   `json:"gender"`
	Image                  string   `json:"image"`
	ResidentSince          string   `json:"residentSince"`
	LprCategory            string   `json:"lprCategory"`
	LprNumber              string   `json:"lprNumber"`
	CommuterClassification string   `json:"commuterClassification"`
	BirthCountry           string   `json:"birthCountry"`
	BirthDate              string   `json:"birthDate"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	Jws                string `json:"jws"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
}

type Options struct {
	ProofType string `json:"proofType"`
}
