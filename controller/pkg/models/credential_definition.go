package models

type CreateCredentialDefinitionRequest struct {
	RevocationRegistrySize int64  `json:"revocation_registry_size"`
	SchemaID               string `json:"schema_id"`
	SupportRevocation      bool   `json:"support_revocation"`
	Tag                    string `json:"tag"`
}

type CreateCredentialDefinitionParams struct {
	ConnID                       string `json:"conn_id"`
	CreateTransactionForEndorser bool   `json:"create_transaction_for_endorser"`
}

type CreateCredentialDefinitionResponse struct {
	CredentialDefinitionID string `json:"credential_definition_id"`
}

type ListCredentialDefinitionsParams struct {
	CredDefID       string `json:"cred_def_id"`
	IssuerDID       string `json:"issuer_did"`
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDID string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

type ListCredentialDefinitionsResponse struct {
	CredentialDefinitionIDS []string `json:"credential_definition_ids"`
}

// * Below structs are specific to Cornerstone_Credential Definition

type GetCredentialDefinitionResponse struct {
	CredentialDefinition CredentialDefinition `json:"credential_definition"`
}

type CredentialDefinition struct {
	Ver      string `json:"ver"`
	ID       string `json:"id"`
	SchemaID string `json:"schemaId"`
	Type     string `json:"type"`
	Tag      string `json:"tag"`
	Value    Value  `json:"value"`
}

type Value struct {
	Primary Primary `json:"primary"`
}

type Primary struct {
	N     string `json:"n"`
	S     string `json:"s"`
	R     R      `json:"r"`
	Rctxt string `json:"rctxt"`
	Z     string `json:"z"`
}

type R struct {
	Forenames      string `json:"forenames"`
	Dateofbirth    string `json:"dateofbirth"`
	Idnumber       string `json:"idnumber"`
	Countryofbirth string `json:"countryofbirth"`
	Gender         string `json:"gender"`
	MasterSecret   string `json:"master_secret"`
	Surname        string `json:"surname"`
}
