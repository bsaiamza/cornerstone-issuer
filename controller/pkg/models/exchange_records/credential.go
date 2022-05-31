package exchange_records

type ListCredentialExchangeRecordsParams struct {
	ConnectionID string `json:"connection_id"`
	Role         string `json:"role"`
	State        string `json:"state"`
	ThreadID     string `json:"thread_id"`
}

type ListCredentialExchangeRecordsResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	CredExRecord CredExRecord `json:"cred_ex_record"`
	Indy         Indy         `json:"indy"`
	LdProof      LdProof      `json:"ld_proof"`
}

type CredExRecord struct {
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

type ByFormat struct {
	CredIssue    Cred `json:"cred_issue"`
	CredOffer    Cred `json:"cred_offer"`
	CredProposal Cred `json:"cred_proposal"`
	CredRequest  Cred `json:"cred_request"`
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

type SAttach struct {
	ID          string `json:"@id"`
	ByteCount   int64  `json:"byte_count"`
	Data        Data   `json:"data"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	LastmodTime string `json:"lastmod_time"`
	MIMEType    string `json:"mime-type"`
}

type Data struct {
	Base64 string   `json:"base64"`
	JSON   string   `json:"json"`
	Jws    Jws      `json:"jws"`
	Links  []string `json:"links"`
	Sha256 string   `json:"sha256"`
}

type Jws struct {
	Header     Header `json:"header"`
	Protected  string `json:"protected"`
	Signature  string `json:"signature"`
	Signatures []Jws  `json:"signatures"`
}

type Header struct {
	Kid string `json:"kid"`
}

type Format struct {
	AttachID string `json:"attach_id"`
	Format   string `json:"format"`
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

type Attribute struct {
	MIMEType string `json:"mime-type"`
	Name     string `json:"name"`
	Value    string `json:"value"`
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

type Indy struct {
	CreatedAt           string `json:"created_at"`
	CredExID            string `json:"cred_ex_id"`
	CredExIndyID        string `json:"cred_ex_indy_id"`
	CredIDStored        string `json:"cred_id_stored"`
	CredRequestMetadata Cred   `json:"cred_request_metadata"`
	CredRevID           string `json:"cred_rev_id"`
	RevRegID            string `json:"rev_reg_id"`
	State               string `json:"state"`
	UpdatedAt           string `json:"updated_at"`
}

type LdProof struct {
	CreatedAt       string `json:"created_at"`
	CredExID        string `json:"cred_ex_id"`
	CredExLdProofID string `json:"cred_ex_ld_proof_id"`
	CredIDStored    string `json:"cred_id_stored"`
	State           string `json:"state"`
	UpdatedAt       string `json:"updated_at"`
}

type IssueCredentialRequest struct {
	Comment string `json:"comment,omitempty"`
}
