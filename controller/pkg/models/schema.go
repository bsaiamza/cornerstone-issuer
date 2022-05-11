package models

// -------------------------- \\
// ----- REQUEST MODELS ----- \\
// -------------------------- \\

// CreateSchemaRequest is the create schema request body object.
type CreateSchemaRequest struct {
	Attributes    []string `json:"attributes"`
	SchemaName    string   `json:"schema_name"`
	SchemaVersion string   `json:"schema_version"`
}

// ------------------------- \\
// ----- PARAMS MODELS ----- \\
// ------------------------- \\

// CreateSchemaParams is the create schema parameters.
type CreateSchemaParams struct {
	ConnID                       string `json:"conn_id"`
	CreateTransactionForEndorser bool   `json:"create_transaction_for_endorser"`
}

// QuerySchemasParams is the query schemas parameters.
type QuerySchemasParams struct {
	SchemaID        string `json:"schema_id"`
	SchemaIssuerDID string `json:"schema_issuer_did"`
	SchemaName      string `json:"schema_name"`
	SchemaVersion   string `json:"schema_version"`
}

// ---------------------------- \\
// ----- RESPONSE MODELS  ----- \\
// ---------------------------- \\

// CreateSchemaResponse is the create schema response body object.
type CreateSchemaResponse struct {
	Sent Sent `json:"sent"`
	Txn  Txn  `json:"txn"`
}

// QuerySchemasResponse is the query schemas response body object.
type QuerySchemasResponse struct {
	SchemaIDs []string `json:"schema_ids"`
}

// GetSchemaResponse is the get schema response body object.
type GetSchemaResponse struct {
	Schema Schema `json:"schema"`
}

// ------------------ \\
// ----- MODELS ----- \\
// ------------------ \\

// Formats is a struct used to hold formats data.
type Formats struct {
	AttachID string `json:"attach_id"`
	Format   string `json:"format"`
}

// MessagesAttach is a struct used to hold messages attach data.
type MessagesAttach struct {
	ID       string             `json:"@id"`
	Data     MessagesAttachData `json:"data"`
	MimeType string             `json:"mime"`
}

// MessagesAttachData is a struct used to hold messages attach data.
type MessagesAttachData struct {
	JSON string `json:"json"`
}

// MetaData is a struct used to hold metadata data.
type MetaData struct {
	Context     MetaDataContext     `json:"context"`
	PostProcess MetaDataPostProcess `json:"post_process"`
}

// MetaDataContext is a struct used to hold metadata context data.
type MetaDataContext struct {
}

// MetaDataPostProcess is a struct used to hold metadata post process data.
type MetaDataPostProcess struct {
	Topic string `json:"topic"`
	Other string `json:"other"`
}

// Schema is a struct used to hold schema data.
// It's used as the get schema response body object.
// It's used as the write record response body object.
type Schema struct {
	AttrNames []string `json:"attrNames"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	SeqNo     int64    `json:"seqNo"`
	Ver       string   `json:"ver"`
	Version   string   `json:"version"`
}

// Sent is a struct used to hold sent data.
type Sent struct {
	Schema   Schema `json:"schema"`
	SchemaID string `json:"schema_id"`
}

// SignatureRequest is a struct used to hold signature request data.
type SignatureRequest struct {
	AuthorGoalCode string `json:"author_goal_code"`
	Context        string `json:"context"`
	Method         string `json:"method"`
	SignatureType  string `json:"signature_type"`
	SignerGoalCode string `json:"signer_goal_code"`
}

// SignatureResponse is a struct used to hold signature response data.
type SignatureResponse struct {
	Context        string `json:"context"`
	MessageID      string `json:"message_id"`
	Method         string `json:"method"`
	SignerGoalCode string `json:"signer_goal_code"`
}

// Timing is a struct used to hold timing data.
type Timing struct {
	ExpiresTime string `json:"expires_time"`
}

// Txn is a struct used to hold txn data.
type Txn struct {
	Type              string              `json:"_type"`
	ConnectionID      string              `json:"connection_id"`
	CreatedAt         string              `json:"created_at"`
	EndorserWriteTxn  bool                `json:"endorser_write_txn"`
	Formats           []Formats           `json:"formats"`
	MessagesAttach    []MessagesAttach    `json:"messages_attach"`
	SignatureRequest  []SignatureRequest  `json:"signature_request"`
	SignatureResponse []SignatureResponse `json:"signature_response"`
	State             string              `json:"state"`
	ThreadID          string              `json:"thread_id"`
	Timing            Timing              `json:"timing"`
	Trace             bool                `json:"trace"`
	TransactionID     string              `json:"transaction_id"`
	UpdatedAt         string              `json:"updated_at"`
}
