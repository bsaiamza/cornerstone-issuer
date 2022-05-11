package models

// -------------------------- \\
// ----- REQUEST MODELS ----- \\
// -------------------------- \\

// ------------------------- \\
// ----- PARAMS MODELS ----- \\
// ------------------------- \\

// QueryRevocationRegestriesParams is the query revocation registries parameters.
type QueryRevocationRegestriesParams struct {
	CredDefID string `json:"cred_def_id"`
	State     string `json:"state"`
}

// ---------------------------- \\
// ----- RESPONSE MODELS  ----- \\
// ---------------------------- \\

// QueryRevocationRegestriesResponse is the query revocation registries response body object.
type QueryRevocationRegestriesResponse struct {
	RevocationRegestryIDs []string `json:"rev_reg_ids"`
}

// GetRevocationRegestryResponse is the get credential definition response body object.
type GetRevocationRegestryResponse struct {
	Result ResultClass `json:"result"`
}

// ------------------ \\
// ----- MODELS ----- \\
// ------------------ \\

type AccumKey struct {
	Z string `json:"z"`
}

type PublicKeys struct {
	AccumKey AccumKey `json:"accumKey"`
}

type ResultClass struct {
	CreatedAt      string        `json:"created_at"`
	CredDefID      string        `json:"cred_def_id"`
	ErrorMsg       string        `json:"error_msg"`
	IssuerDid      string        `json:"issuer_did"`
	MaxCredNum     int64         `json:"max_cred_num"`
	PendingPub     []string      `json:"pending_pub"`
	RecordID       string        `json:"record_id"`
	RevocDefType   string        `json:"revoc_def_type"`
	RevocRegDef    RevocRegDef   `json:"revoc_reg_def"`
	RevocRegEntry  RevocRegEntry `json:"revoc_reg_entry"`
	RevocRegID     string        `json:"revoc_reg_id"`
	State          string        `json:"state"`
	Tag            string        `json:"tag"`
	TailsHash      string        `json:"tails_hash"`
	TailsLocalPath string        `json:"tails_local_path"`
	TailsPublicURI string        `json:"tails_public_uri"`
	UpdatedAt      string        `json:"updated_at"`
}

type RevocRegDef struct {
	CredDefID    string           `json:"credDefId"`
	ID           string           `json:"id"`
	RevocDefType string           `json:"revocDefType"`
	Tag          string           `json:"tag"`
	Value        RevocRegDefValue `json:"value"`
	Ver          string           `json:"ver"`
}

type RevocRegDefValue struct {
	IssuanceType  string     `json:"issuanceType"`
	MaxCredNum    int64      `json:"maxCredNum"`
	PublicKeys    PublicKeys `json:"publicKeys"`
	TailsHash     string     `json:"tailsHash"`
	TailsLocation string     `json:"tailsLocation"`
}

type RevocRegEntry struct {
	Value RevocRegEntryValue `json:"value"`
	Ver   string             `json:"ver"`
}

type RevocRegEntryValue struct {
	Accum     string  `json:"accum"`
	PrevAccum string  `json:"prevAccum"`
	Revoked   []int64 `json:"revoked"`
}
