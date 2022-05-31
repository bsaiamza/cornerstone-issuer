package models

type CreateSchemaRequest struct {
	Attributes    []string `json:"attributes"`
	SchemaName    string   `json:"schema_name"`
	SchemaVersion string   `json:"schema_version"`
}

type CreateSchemaParams struct {
	ConnID                       string `json:"conn_id"`
	CreateTransactionForEndorser bool   `json:"create_transaction_for_endorser"`
}

type CreateSchemaResponse struct {
	SchemaID string `json:"schema_id"`
	Schema   Schema `json:"schema"`
}

type Schema struct {
	Ver       string   `json:"ver"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	AttrNames []string `json:"attrNames"`
	SeqNo     int64    `json:"seqNo"`
}

type LitSchemasParams struct {
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDID string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

type ListSchemasResponse struct {
	SchemaIDS []string `json:"schema_ids"`
}

type GetSchemaResponse struct {
	Schema Schema `json:"schema"`
}
