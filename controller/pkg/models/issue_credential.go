package models

type IssueCornerstoneCredentialRequest struct {
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

type Attribute struct {
	MIMEType string `json:"mime-type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

type IssueCornerstoneCredentialResponse struct {
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

// type CredentialOfferDict struct {
// 	Type              string         `json:"@type"`
// 	ID                string         `json:"@id"`
// 	Thread            Thread         `json:"~thread"`
// 	CredentialPreview CredentialPR   `json:"credential_preview"`
// 	Comment           string         `json:"comment"`
// 	OffersAttach      []OffersAttach `json:"offers~attach"`
// }

type CredentialPR struct {
	Type       string      `json:"@type"`
	Attributes []Attribute `json:"attributes"`
}

type OffersAttach struct {
	ID       string `json:"@id"`
	MIMEType string `json:"mime-type"`
	Data     Data   `json:"data"`
}

// type Data struct {
// 	Base64 string `json:"base64"`
// }

type Thread struct {
}

// type CredentialProposalDict struct {
// 	Type               string       `json:"@type"`
// 	ID                 string       `json:"@id"`
// 	CredentialProposal CredentialPR `json:"credential_proposal"`
// 	IssuerDid          string       `json:"issuer_did"`
// 	CredDefID          string       `json:"cred_def_id"`
// 	SchemaIssuerDid    string       `json:"schema_issuer_did"`
// 	SchemaID           string       `json:"schema_id"`
// 	SchemaVersion      string       `json:"schema_version"`
// 	SchemaName         string       `json:"schema_name"`
// 	Comment            string       `json:"comment"`
// }

type ListCredentialRecordsParams struct {
	ConnectionID string `json:"connection_id"`
	Role         string `json:"role"`
	State        string `json:"state"`
	ThreadID     string `json:"thread_id"`
}

type ListCredentialRecordsResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
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

type CredentialOfferDict struct {
	Type              string         `json:"@type"`
	ID                string         `json:"@id"`
	Thread            Thread         `json:"~thread"`
	CredentialPreview CredentialPR   `json:"credential_preview"`
	OffersAttach      []OffersAttach `json:"offers~attach"`
	Comment           string         `json:"comment"`
}

type Data struct {
	Base64 string `json:"base64"`
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

type BlindedMSCorrectnessProof struct {
	C        string `json:"c"`
	VDashCap string `json:"v_dash_cap"`
	MCaps    MCaps  `json:"m_caps"`
	RCaps    Thread `json:"r_caps"`
}

type MCaps struct {
	MasterSecret string `json:"master_secret"`
}
