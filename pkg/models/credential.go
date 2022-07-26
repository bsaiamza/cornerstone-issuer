package models

type ListCredentialRecordsResponse struct {
	Results []Records `json:"results"`
}

type Records struct {
	Credential             *Credential            `json:"credential,omitempty"`
	CredentialOffer        CredentialOffer        `json:"credential_offer"`
	SchemaID               string                 `json:"schema_id"`
	Role                   string                 `json:"role"`
	UpdatedAt              string                 `json:"updated_at"`
	CredentialProposalDict CredentialProposalDict `json:"credential_proposal_dict"`
	ConnectionID           string                 `json:"connection_id"`
	CredentialDefinitionID string                 `json:"credential_definition_id"`
	Trace                  bool                   `json:"trace"`
	Initiator              string                 `json:"initiator"`
	AutoIssue              bool                   `json:"auto_issue"`
	AutoOffer              bool                   `json:"auto_offer"`
	CredentialExchangeID   string                 `json:"credential_exchange_id"`
	AutoRemove             bool                   `json:"auto_remove"`
	ThreadID               string                 `json:"thread_id"`
	State                  *string                `json:"state,omitempty"`
	CreatedAt              string                 `json:"created_at"`
	CredentialOfferDict    CredentialOfferDict    `json:"credential_offer_dict"`
	CredentialRequest      *CredentialRequest     `json:"credential_request,omitempty"`
	ErrorMsg               *string                `json:"error_msg,omitempty"`
}

type Credential struct {
	SchemaID  string `json:"schema_id"`
	CredDefID string `json:"cred_def_id"`
}

type CredentialRequest struct {
	ProverDid                 string                    `json:"prover_did"`
	CredDefID                 string                    `json:"cred_def_id"`
	BlindedMS                 BlindedMS                 `json:"blinded_ms"`
	BlindedMSCorrectnessProof BlindedMSCorrectnessProof `json:"blinded_ms_correctness_proof"`
	Nonce                     string                    `json:"nonce"`
}

type BlindedMS struct {
	U                   string      `json:"u"`
	Ur                  interface{} `json:"ur"`
	HiddenAttributes    []string    `json:"hidden_attributes"`
	CommittedAttributes Thread      `json:"committed_attributes"`
}

type Thread struct {
}

type BlindedMSCorrectnessProof struct {
	C        string `json:"c"`
	VDashCap string `json:"v_dash_cap"`
	MCaps    MCaps  `json:"m_caps"`
	RCaps    Thread `json:"r_caps"`
}

type MCaps struct {
	MasterSecret string `json:"master_secret"`
}

type CredentialOfferDict struct {
	Type              string         `json:"@type"`
	ID                string         `json:"@id"`
	Thread            Thread         `json:"~thread"`
	CredentialPreview CredentialPR   `json:"credential_preview"`
	OffersAttach      []OffersAttach `json:"offers~attach"`
	Comment           string         `json:"comment"`
}

type CredentialPR struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	MIMEType string `json:"mime-type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type OffersAttach struct {
	ID       string `json:"@id"`
	MIMEType string `json:"mime-type"`
	Data     Data   `json:"data"`
}

type Data struct {
	Base64 string `json:"base64"`
}

type CredentialOffer struct {
	SchemaID            string              `json:"schema_id"`
	CredDefID           string              `json:"cred_def_id"`
	Nonce               string              `json:"nonce"`
	KeyCorrectnessProof KeyCorrectnessProof `json:"key_correctness_proof"`
}

type KeyCorrectnessProof struct {
	C     string     `json:"c"`
	XzCap string     `json:"xz_cap"`
	XrCap [][]string `json:"xr_cap"`
}

type CredentialProposalDict struct {
	Type               string       `json:"@type"`
	ID                 string       `json:"@id"`
	CredentialProposal CredentialPR `json:"credential_proposal"`
	Comment            string       `json:"comment"`
	SchemaID           *string      `json:"schema_id,omitempty"`
	CredDefID          *string      `json:"cred_def_id,omitempty"`
	SchemaIssuerDid    *string      `json:"schema_issuer_did,omitempty"`
	IssuerDid          *string      `json:"issuer_did,omitempty"`
	SchemaName         *string      `json:"schema_name,omitempty"`
	SchemaVersion      *string      `json:"schema_version,omitempty"`
}

type IssueCredentialRequest struct {
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

type CredentialProposal struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type IssueCredentialResponse struct {
	CreatedAt              string                 `json:"created_at"`
	CredentialOfferDict    CredentialOfferDict    `json:"credential_offer_dict"`
	ThreadID               string                 `json:"thread_id"`
	CredentialProposalDict CredentialProposalDict `json:"credential_proposal_dict"`
	Initiator              string                 `json:"initiator"`
	ConnectionID           string                 `json:"connection_id"`
	CredentialDefinitionID string                 `json:"credential_definition_id"`
	CredentialOffer        CredentialOffer        `json:"credential_offer"`
	CredentialExchangeID   string                 `json:"credential_exchange_id"`
	AutoRemove             bool                   `json:"auto_remove"`
	AutoIssue              bool                   `json:"auto_issue"`
	Trace                  bool                   `json:"trace"`
	AutoOffer              bool                   `json:"auto_offer"`
	State                  string                 `json:"state"`
	UpdatedAt              string                 `json:"updated_at"`
	SchemaID               string                 `json:"schema_id"`
	Role                   string                 `json:"role"`
}

type CreateCredentialRequest struct {
	AutoRemove         bool               `json:"auto_remove"`
	Comment            string             `json:"comment"`
	CredDefID          string             `json:"cred_def_id"`
	CredentialProposal CredentialProposal `json:"credential_proposal"`
	IssuerDid          string             `json:"issuer_did"`
	SchemaID           string             `json:"schema_id"`
	SchemaIssuerDid    string             `json:"schema_issuer_did"`
	SchemaName         string             `json:"schema_name"`
	SchemaVersion      string             `json:"schema_version"`
	Trace              bool               `json:"trace"`
}

type IssueCredentialWebhookResponse struct {
	CredentialExchangeID   string                 `json:"credential_exchange_id"`
	ConnectionID           string                 `json:"connection_id"`
	ThreadID               string                 `json:"thread_id"`
	ParentThreadID         string                 `json:"parent_thread_id"`
	Initiator              string                 `json:"initiator"`
	State                  string                 `json:"state"`
	CredentialDefinitionID string                 `json:"credential_definition_id"`
	SchemaID               string                 `json:"schema_id"`
	CredentialProposalDict CredentialProposalDict `json:"credential_proposal_dict"`
	CredentialOffer        CredentialOffer        `json:"credential_offer"`
	CredentialRequest      CredentialRequest      `json:"credential_request"`
	CredentialRequestMeta  CredentialRequestMeta  `json:"credential_request_meta"`
	CredentialID           string                 `json:"credential_id"`
	RawCredential          string                 `json:"raw_credential"`
	Credential             Credential             `json:"credential"`
	AutoOffer              bool                   `json:"auto_offer"`
	AutoIssue              bool                   `json:"auto_issue"`
	ErrorMsg               string                 `json:"error_msg"`
}

type CredentialRequestMeta struct{}
