package models

// ----------------------------- \\
// ----- REQUEST MODELS V1 ----- \\
// ----------------------------- \\

// SendProposalV1Request is the send proposal request body object.
type SendProposalV1Request struct {
	AutoRemove         bool               `json:"auto_remove"`
	Comment            string             `json:"comment"`
	ConnectionID       string             `json:"connection_id"`
	CredDefID          string             `json:"cred_def_id"`
	CredentialProposal CredentialProposal `json:"credential_proposal"`
	IssuerDid          string             `json:"issuer_did"`
	SchemaID           string             `json:"schema_id"`
	SchemaIssuerDid    string             `json:"schema_issuer_did"`
	SchemaName         string             `json:"schema_name"`
	SchemaVersion      string             `json:"schema_version"`
	Trace              bool               `json:"trace"`
}

// SendOfferV1Request is the send offer request body object.
type SendOfferV1Request struct {
	CounterProposal CounterProposal `json:"counter_proposal"`
}

// IssueCredentialRequest is the issue credential request body object.
// Used for v1 and v2.
type IssueCredentialRequest struct {
	Comment string `json:"comment,omitempty"`
}

// StoreCredentialRequest is the store credential request body object.
// Used for v1 and v2.
type StoreCredentialRequest struct {
	CredentialID string `json:"credential_id"`
}

// ----------------------------- \\
// ----- REQUEST MODELS V2 ----- \\
// ----------------------------- \\

// SendProposalV2Request is the send proposal request body object.
type SendProposalV2Request struct {
	AutoRemove        bool              `json:"auto_remove"`
	Comment           string            `json:"comment"`
	ConnectionID      string            `json:"connection_id"`
	CredentialPreview CredentialPreview `json:"credential_preview"`
	Filter            Filter            `json:"filter"`
	Trace             bool              `json:"trace"`
}

// SendOfferV2Request is the send offer request body object.
type SendOfferV2Request struct {
	CounterPreview CounterPreview `json:"counter_preview"`
	Filter         Filter         `json:"filter"`
}

// SendRequestV2Request is the send request request body object.
type SendRequestV2Request struct {
	HolderDid string `json:"holder_did"`
}

// ---------------------------- \\
// ----- PARAMS MODELS V1 ----- \\
// ---------------------------- \\

// QueryExchangeRecordsParams is the query credential proposals (exchange records) parameters.
// Used for v1 and v2.
type QueryExchangeRecordsParams struct {
	ConnectionID string `json:"connection_id"`
	Role         string `json:"role"`
	State        string `json:"state"`
	ThreadID     string `json:"thread_id"`
}

// QueryCredentialsParams is the query credentials parameters.
type QueryCredentialsParams struct {
	Count string `json:"count"`
	Start string `json:"start"`
	Wql   string `json:"wql"`
}

// ------------------------------- \\
// ----- RESPONSE MODELS V1  ----- \\
// ------------------------------- \\

// IssueCredentialResponse is the issue credential process response body object.
// It's used as the send proposal response body object.
// It's used as the send offer response body object.
// It's used as the send request response body object.
// It's used as the issue credential response body object.
// It's used as the store credential response body object.
// It's used as the get credential response body object.
type IssueCredentialResponse struct {
	AutoIssue                 bool                      `json:"auto_issue"`
	AutoOffer                 bool                      `json:"auto_offer"`
	AutoRemove                bool                      `json:"auto_remove"`
	ConnectionID              string                    `json:"connection_id"`
	CreatedAt                 string                    `json:"created_at"`
	Credential                Credential                `json:"credential"`
	CredentialDefinitionID    string                    `json:"credential_definition_id"`
	CredentialExchangeID      string                    `json:"credential_exchange_id"`
	CredentialID              string                    `json:"credential_id"`
	CredentialOffer           CredentialOffer           `json:"credential_offer"`
	CredentialOfferDict       CredentialOfferDict       `json:"credential_offer_dict"`
	CredentialProposalDict    CredentialProposalDict    `json:"credential_proposal_dict"`
	CredentialRequest         CredentialRequest         `json:"credential_request"`
	CredentialRequestMetadata CredentialRequestMetadata `json:"credential_request_metadata"`
	ErrorMsg                  string                    `json:"error_msg"`
	Initiator                 string                    `json:"initiator"`
	ParentThreadID            string                    `json:"parent_thread_id"`
	RawCredential             RawCredential             `json:"raw_credential"`
	RevocRegID                string                    `json:"revoc_reg_id"`
	RevocationID              string                    `json:"revocation_id"`
	Role                      string                    `json:"role"`
	SchemaID                  string                    `json:"schema_id"`
	State                     string                    `json:"state"`
	ThreadID                  string                    `json:"thread_id"`
	Trace                     bool                      `json:"trace"`
	UpdatedAt                 string                    `json:"updated_at"`
}

// QueryExchangeRecordsV1Response is the query exchange records response body object.
type QueryExchangeRecordsV1Response struct {
	Results []IssueCredentialResponse `json:"results"`
}

// QueryCredentialsV1Response is the query credentials response body object.
type QueryCredentialsResponse struct {
	Results []Credential `json:"results"`
}

// ------------------------------- \\
// ----- RESPONSE MODELS V2  ----- \\
// ------------------------------- \\

// IssueCredentialV2Response is the issue credential process response body object.
// It's used as the send proposal response body object.
// It's used as the send offer response body object.
// It's used as the send request response body object.
// It's used as the issue credential response body object.
// It's used as the store credential response body object.
// It's used as the get credential response body object.
type IssueCredentialV2Response struct {
	AutoIssue      bool         `json:"auto_issue"`
	AutoOffer      bool         `json:"auto_offer"`
	AutoRemove     bool         `json:"auto_remove"`
	ByFormat       ByFormat     `json:"by_format"`
	ConnectionID   string       `json:"connection_id"`
	CreatedAt      string       `json:"created_at"`
	CredExID       string       `json:"cred_ex_id"`
	CredIssue      CredIssue    `json:"cred_issue"`
	CredOffer      CredOffer    `json:"cred_offer"`
	CredPreview    CredPreview  `json:"cred_preview"`
	CredProposal   CredProposal `json:"cred_proposal"`
	CredRequest    CredRequest  `json:"cred_request"`
	ErrorMsg       string       `json:"error_msg"`
	Initiator      string       `json:"initiator"`
	ParentThreadID string       `json:"parent_thread_id"`
	Role           string       `json:"role"`
	State          string       `json:"state"`
	ThreadID       string       `json:"thread_id"`
	Trace          bool         `json:"trace"`
	UpdatedAt      string       `json:"updated_at"`
}

// QueryExchangeRecordsV2Response is the query exchange records response body object.
type QueryExchangeRecordsV2Response struct {
	Results []IssueCredentialV2Response `json:"results"`
}

// ------------------ \\
// ----- MODELS ----- \\
// ------------------ \\

type AdditionalProp struct {
	Encoded string `json:"encoded"`
	Raw     string `json:"raw"`
}

type Attrs struct {
}

type Attribute struct {
	MIMEType string `json:"mime-type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type ByFormat struct {
	CredIssue    Cred `json:"cred_issue"`
	CredOffer    Cred `json:"cred_offer"`
	CredProposal Cred `json:"cred_proposal"`
	CredRequest  Cred `json:"cred_request"`
}

type CounterPreview struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CounterProposal struct {
	ID                 string             `json:"@id"`
	Comment            string             `json:"comment"`
	CredDefID          string             `json:"cred_def_id"`
	CredentialProposal CredentialProposal `json:"credential_proposal"`
	IssuerDid          string             `json:"issuer_did"`
	SchemaID           string             `json:"schema_id"`
	SchemaIssuerDid    string             `json:"schema_issuer_did"`
	SchemaName         string             `json:"schema_name"`
	SchemaVersion      string             `json:"schema_version"`
}

type Cred struct {
}

type CredIssue struct {
	ID                string    `json:"@id"`
	Type              string    `json:"@type"`
	Comment           string    `json:"comment"`
	CredentialsAttach []SAttach `json:"credentials~attach"`
	Formats           []Format  `json:"formats"`
	ReplacementID     string    `json:"replacement_id"`
}

type CredOffer struct {
	ID                string      `json:"@id"`
	Type              string      `json:"@type"`
	Comment           string      `json:"comment"`
	CredentialPreview CredPreview `json:"credential_preview"`
	Formats           []Format    `json:"formats"`
	OffersAttach      []SAttach   `json:"offers~attach"`
	ReplacementID     string      `json:"replacement_id"`
}

type CredPreview struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CredProposal struct {
	ID                string      `json:"@id"`
	Type              string      `json:"@type"`
	Comment           string      `json:"comment"`
	CredentialPreview CredPreview `json:"credential_preview"`
	FiltersAttach     []SAttach   `json:"filters~attach"`
	Formats           []Format    `json:"formats"`
}

type CredRequest struct {
	ID             string    `json:"@id"`
	Type           string    `json:"@type"`
	Comment        string    `json:"comment"`
	Formats        []Format  `json:"formats"`
	RequestsAttach []SAttach `json:"requests~attach"`
}

// It's used as the get credential response body object.
type Credential struct {
	Attrs     Attrs  `json:"attrs"`
	CredDefID string `json:"cred_def_id"`
	CredRevID string `json:"cred_rev_id"`
	Referent  string `json:"referent"`
	RevRegID  string `json:"rev_reg_id"`
	SchemaID  string `json:"schema_id"`
}

type CredentialV2 struct {
	Context           []string          `json:"@context"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Description       string            `json:"description"`
	Identifier        string            `json:"identifier"`
	IssuanceDate      string            `json:"issuanceDate"`
	Issuer            string            `json:"issuer"`
	Name              string            `json:"name"`
	Type              []string          `json:"type"`
}

type CredentialOffer struct {
	CredDefID           string              `json:"cred_def_id"`
	KeyCorrectnessProof KeyCorrectnessProof `json:"key_correctness_proof"`
	Nonce               string              `json:"nonce"`
	SchemaID            string              `json:"schema_id"`
}

type CredentialOfferDict struct {
	ID                string         `json:"@id"`
	Type              string         `json:"@type"`
	Comment           string         `json:"comment"`
	CredentialPreview CredentialPR   `json:"credential_preview"`
	OffersAttach      []OffersAttach `json:"offers~attach"`
}

type CredentialPR struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CredentialPreview struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CredentialProposal struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type CredentialProposalDict struct {
	ID                 string       `json:"@id"`
	Type               string       `json:"@type"`
	Comment            string       `json:"comment"`
	CredDefID          string       `json:"cred_def_id"`
	CredentialProposal CredentialPR `json:"credential_proposal"`
	IssuerDid          string       `json:"issuer_did"`
	SchemaID           string       `json:"schema_id"`
	SchemaIssuerDid    string       `json:"schema_issuer_did"`
	SchemaName         string       `json:"schema_name"`
	SchemaVersion      string       `json:"schema_version"`
}

type CredentialRequest struct {
	BlindedMS                 CredentialRequestMetadata `json:"blinded_ms"`
	BlindedMSCorrectnessProof CredentialRequestMetadata `json:"blinded_ms_correctness_proof"`
	CredDefID                 string                    `json:"cred_def_id"`
	Nonce                     string                    `json:"nonce"`
	ProverDid                 string                    `json:"prover_did"`
}

type CredentialRequestMetadata struct {
}

type CredentialSubject struct {
}

type Filter struct {
	Indy    Indy    `json:"indy"`
	LdProof LdProof `json:"ld_proof"`
}

type Format struct {
	AttachID string `json:"attach_id"`
	Format   string `json:"format"`
}

type Indy struct {
	CredDefID       string `json:"cred_def_id"`
	IssuerDid       string `json:"issuer_did"`
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDid string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

type KeyCorrectnessProof struct {
	C     string     `json:"c"`
	XrCap [][]string `json:"xr_cap"`
	XzCap string     `json:"xz_cap"`
}

type LdProof struct {
	Credential CredentialV2 `json:"credential"`
	Options    Options      `json:"options"`
}

type OffersAttach struct {
	ID          string `json:"@id"`
	ByteCount   int64  `json:"byte_count"`
	Data        Data   `json:"data"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	LastmodTime string `json:"lastmod_time"`
	MIMEType    string `json:"mime-type"`
}

type Options struct {
	ProofType string `json:"proofType"`
}

type RawCredential struct {
	CredDefID                 string                    `json:"cred_def_id"`
	RevReg                    CredentialRequestMetadata `json:"rev_reg"`
	RevRegID                  string                    `json:"rev_reg_id"`
	SchemaID                  string                    `json:"schema_id"`
	Signature                 CredentialRequestMetadata `json:"signature"`
	SignatureCorrectnessProof CredentialRequestMetadata `json:"signature_correctness_proof"`
	Values                    Values                    `json:"values"`
	Witness                   CredentialRequestMetadata `json:"witness"`
}

type SAttach struct {
	ID          string `json:"@id"`
	ByteCount   int64  `json:"byte_count"`
	Data        Data   `json:"data"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	LastmodTime string `json:"lastmod_time"`
	MIMEType    string `json:"mime-type"`
}

type Values struct {
}
