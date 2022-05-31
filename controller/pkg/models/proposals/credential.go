package proposals

type CornerstoneCredentialProposalRequest struct {
	AutoRemove     bool           `json:"auto_remove"`
	Comment        string         `json:"comment"`
	ConnectionID   string         `json:"connection_id"`
	CounterPreview CounterPreview `json:"counter_preview"`
	Filter         Filter         `json:"filter"`
	Trace          bool           `json:"trace"`
}

type CounterPreview struct {
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
	Context           []string          `json:"@context"`
	ID                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	Identifier        string            `json:"identifier"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	IssuanceDate      string            `json:"issuanceDate"`
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

type CornerstoneCredentialProposalResponse struct {
	AutoOffer    bool                                              `json:"auto_offer"`
	CreatedAt    string                                            `json:"created_at"`
	ByFormat     ByFormat                                          `json:"by_format"`
	Trace        bool                                              `json:"trace"`
	State        string                                            `json:"state"`
	ThreadID     string                                            `json:"thread_id"`
	CredPreview  CredPreview                                       `json:"cred_preview"`
	CredProposal CornerstoneCredentialProposalResponseCredProposal `json:"cred_proposal"`
	AutoRemove   bool                                              `json:"auto_remove"`
	UpdatedAt    string                                            `json:"updated_at"`
	Role         string                                            `json:"role"`
	AutoIssue    bool                                              `json:"auto_issue"`
	ConnectionID string                                            `json:"connection_id"`
	CredExID     string                                            `json:"cred_ex_id"`
	Initiator    string                                            `json:"initiator"`
}

type ByFormat struct {
	CredProposal ByFormatCredProposal `json:"cred_proposal"`
}

type ByFormatCredProposal struct {
	Indy    Indy    `json:"indy"`
	LdProof LdProof `json:"ld_proof"`
}

type CredPreview struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CornerstoneCredentialProposalResponseCredProposal struct {
	Type              string          `json:"@type"`
	ID                string          `json:"@id"`
	CredentialPreview CredPreview     `json:"credential_preview"`
	Comment           string          `json:"comment"`
	Formats           []Format        `json:"formats"`
	FiltersAttach     []FiltersAttach `json:"filters~attach"`
}

type FiltersAttach struct {
	ID       string `json:"@id"`
	MIMEType string `json:"mime-type"`
	Data     Data   `json:"data"`
}

type Data struct {
	Base64 string `json:"base64"`
}

type Format struct {
	AttachID string `json:"attach_id"`
	Format   string `json:"format"`
}

type ProposalRequest struct {
	ConnectionID           string `json:"connection_id,omitempty"`
	ID                     string `json:"id,omitempty"`
	GivenName              string `json:"given_name,omitempty"`
	FamilyName             string `json:"family_name,omitempty"`
	Gender                 string `json:"gender,omitempty"`
	LprNumber              string `json:"lprNumber,omitempty"`
	LprCategory            string `json:"lprCategory,omitempty"`
	ResidentSince          string `json:"residentSince,omitempty"`
	CommuterClassification string `json:"commuterClassification,omitempty"`
	BirthDate              string `json:"birthDate,omitempty"`
	BirthCountry           string `json:"birthCountry,omitempty"`
	CredDefID              string `json:"credDefID,omitempty"`
	IssuerDid              string `json:"issuerDid,omitempty"`
	SchemaID               string `json:"schemaID,omitempty"`
	SchemaIssuerDid        string `json:"schemaIssuerDid,omitempty"`
	SchemaName             string `json:"schemaName,omitempty"`
	SchemaVersion          string `json:"schemaVersion,omitempty"`
	Issuer                 string `json:"issuer,omitempty"`
}
