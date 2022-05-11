package models

// -------------------------- \\
// ----- REQUEST MODELS ----- \\
// -------------------------- \\

// CreateCredentialDefinitionRequest is the create credential definition request body object.
type CreateCredentialDefinitionRequest struct {
	RevocationRegistrySize int64  `json:"revocation_registry_size,omitempty"`
	SchemaID               string `json:"schema_id"`
	SupportRevocation      bool   `json:"support_revocation,omitempty"`
	Tag                    string `json:"tag"`
}

// ------------------------- \\
// ----- PARAMS MODELS ----- \\
// ------------------------- \\

// CreateCredentialDefinitionParams is the create credential definition parameters.
type CreateCredentialDefinitionParams struct {
	ConnID                       string `json:"conn_id"`
	CreateTransactionForEndorser bool   `json:"create_transaction_for_endorser"`
}

// QueryCredentialDefinitionsParams is the query credential definitions parameters.
type QueryCredentialDefinitionsParams struct {
	CredDefID       string `json:"cred_def_id"`
	IssuerDID       string `json:"issuer_did"`
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDID string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

// ---------------------------- \\
// ----- RESPONSE MODELS  ----- \\
// ---------------------------- \\

// CreateCredentialDefinitionResponse is the create credential definition response body object.
type CreateCredentialDefinitionResponse struct {
	Sent CredDefSent `json:"sent"`
	Txn  Txn         `json:"txn"`
}

// QueryCredentialDefinitionsResponse is the query credential definitions response body object.
type QueryCredentialDefinitionsResponse struct {
	CredentialDefinitionIDs []string `json:"credential_definition_ids"`
}

// GetCredentialDefinitionResponse is the get credential definition response body object.
type GetCredentialDefinitionResponse struct {
	CredentialDefinition CredentialDefinition `json:"credential_definition"`
}

// ------------------ \\
// ----- MODELS ----- \\
// ------------------ \\

// CredentialDefinition is a struct used to hold credential definition data.
// It's used as the get credential definition response body object.
// It's used as the write record response body object.
type CredentialDefinition struct {
	ID       string `json:"id"`
	SchemaID string `json:"schemaId"`
	Tag      string `json:"tag"`
	Type     string `json:"type"`
	Value    Value  `json:"value"`
	Ver      string `json:"ver"`
}

// CredDefSent is a struct used to hold sent data.
type CredDefSent struct {
	CredentialDefinitionID string `json:"credential_definition_id"`
}

// Primary is a struct used to hold primary data.
type Primary struct {
	N     string `json:"n"`
	R     R      `json:"r"`
	Rctxt string `json:"rctxt"`
	S     string `json:"s"`
	Z     string `json:"z"`
}

// R is a struct used to hold r data.
type R struct {
	MasterSecret string `json:"master_secret"`
	Number       string `json:"number"`
	Remainder    string `json:"remainder"`
}

// Revocation is a struct used to hold revocation data.
type Revocation struct {
	G      string `json:"g"`
	GDash  string `json:"g_dash"`
	H      string `json:"h"`
	H0     string `json:"h0"`
	H1     string `json:"h1"`
	H2     string `json:"h2"`
	HCap   string `json:"h_cap"`
	Htilde string `json:"htilde"`
	Pk     string `json:"pk"`
	U      string `json:"u"`
	Y      string `json:"y"`
}

// Value is a struct used to hold value data.
type Value struct {
	Primary    Primary `json:"primary"`
	Revocation Primary `json:"revocation"`
}
