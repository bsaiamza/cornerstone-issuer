package models

// ----------------------------- \\
// ----- REQUEST MODELS V1 ----- \\
// ----------------------------- \\

// CreateInvitationV1Request is the create invitation request body object.
type CreateInvitationV1Request struct {
	MediationID     string   `json:"mediation_id,omitempty"`
	Metadata        Metadata `json:"metadata,omitempty"`
	MyLabel         string   `json:"my_label,omitempty"`
	RecipientKeys   []string `json:"recipient_keys,omitempty"`
	RoutingKeys     []string `json:"routing_keys,omitempty"`
	ServiceEndpoint string   `json:"service_endpoint,omitempty"`
}

// SendBasicMessageRequest is the send basic message request body object.
type SendBasicMessageRequest struct {
	Content string `json:"content,omitempty"`
}

// TrustPingRequest is the trust ping request body object.
type TrustPingRequest struct {
	Comment string `json:"comment,omitempty"`
}

// ----------------------------- \\
// ----- REQUEST MODELS V2 ----- \\
// ----------------------------- \\

// CreateInvitationV2Request is the create invitation request body object.
type CreateInvitationV2Request struct {
	Alias              string        `json:"alias"`
	Attachments        []Attachments `json:"attachments,omitempty"`
	HandshakeProtocols []string      `json:"handshake_protocols"`
	MediationID        string        `json:"mediation_id,omitempty"`
	Metadata           Metadata      `json:"metadata,omitempty"`
	MyLabel            string        `json:"my_label"`
	UsePublicDID       bool          `json:"use_public_did"`
}

// ---------------------------- \\
// ----- PARAMS MODELS V1 ----- \\
// ---------------------------- \\

// CreateInvitationV1Params is the create invitation parameters.
type CreateInvitationV1Params struct {
	Alias      string `json:"alias"`
	AutoAccept bool   `json:"auto_accept"`
	MultiUse   bool   `json:"multi_use"`
	Public     bool   `json:"public"`
}

// ReceiveInvitationV1Params is the receive invitation parameters.
type ReceiveInvitationV1Params struct {
	Alias       string `json:"alias"`
	AutoAccept  bool   `json:"auto_accept"`
	MediationID string `json:"mediation_id"`
}

// AcceptInvitationV1Params is the accept invitation parameters.
type AcceptInvitationV1Params struct {
	MediationID string `json:"mediation_id"`
	MyEndpoint  string `json:"my_endpoint"`
	MyLabel     string `json:"my_label"`
}

// AcceptRequestV1Params is the accept request parameters.
type AcceptRequestV1Params struct {
	MyEndpoint string `json:"my_endpoint"`
}

// QueryConnectionsParams is the query connections parameters.
type QueryConnectionsParams struct {
	Alias              string `json:"alias"`
	ConnectionProtocol string `json:"connection_protocol"`
	InvitationKey      string `json:"invitation_key"`
	MyDID              string `json:"my_did"`
	State              string `json:"state"`
	TheirDID           string `json:"their_did"`
	TheirPublicDID     string `json:"their_public_did"`
	TheirRole          string `json:"their_role"`
}

// ---------------------------- \\
// ----- PARAMS MODELS V2 ----- \\
// ---------------------------- \\

// CreateInvitationV2Params is the create invitation parameters.
type CreateInvitationV2Params struct {
	AutoAccept bool `json:"auto_accept"`
	MultiUse   bool `json:"multi_use"`
}

// ReceiveInvitationV2Params is the receive invitation parameters.
type ReceiveInvitationV2Params struct {
	Alias                 string `json:"alias"`
	AutoAccept            bool   `json:"auto_accept"`
	MediationID           string `json:"mediation_id"`
	UseExistingConnection bool   `json:"use_existing_connection"`
}

// AcceptInvitationV2Params is the accept invitation parameters.
type AcceptInvitationV2Params struct {
	MyEndpoint string `json:"my_endpoint"`
	MyLabel    string `json:"my_label"`
}

// AcceptRequestV2Params is the accept request parameters.
type AcceptRequestV2Params struct {
	MediationID string `json:"mediation_id"`
	MyEndpoint  string `json:"my_endpoint"`
}

// ------------------------------ \\
// ----- RESPONSE MODELS V1 ----- \\
// ------------------------------ \\

// CreateInvitationV1Response is the create invitation response body object.
type CreateInvitationV1Response struct {
	ConnectionID  string       `json:"connection_id"`
	Invitation    InvitationV1 `json:"invitation"`
	InvitationURL string       `json:"invitation_url"`
}

// QueryConnectionsResponse is the create invitation response body object.
type QueryConnectionsResponse struct {
	Results []Connection `json:"results"`
}

// TrustPingResponse is the trust ping response body object.
type TrustPingResponse struct {
	ThreadID string `json:"thread_id"`
}

// ------------------------------ \\
// ----- RESPONSE MODELS V2 ----- \\
// ------------------------------ \\

// CreateInvitationV2Response is the create invitation response body object.
type CreateInvitationV2Response struct {
	CreatedAt           string       `json:"created_at"`
	InvitationMessageID string       `json:"invi_msg_id"`
	Invitation          InvitationV2 `json:"invitation"`
	InvitationID        string       `json:"invitation_id"`
	InvitationURL       string       `json:"invitation_url"`
	State               string       `json:"state"`
	Trace               bool         `json:"trace"`
	UpdatedAt           string       `json:"updated_at"`
}

// ------------------ \\
// ----- MODELS ----- \\
// ------------------ \\

// Attachments is a struct used to hold attachments data.
type Attachments struct {
	ID   string `json:"id"`   // either CredentialExchangeID or PresentationExchangeID
	Type string `json:"type"` // either credential-offer or present-proof
}

// Connection is a struct used to hold connection data.
// It's used as the receive invitation v1 response body object.
// It's used as the accept invitation v1 response body object.
// It's used as the accept request v1 response body object.
// It's used as the receive invitation v2 response body object.
// It's used as the accept invitation v2 response body object.
// It's used as the accept request v2 response body object.
type Connection struct {
	Accept              string `json:"accept"`
	Alias               string `json:"alias"`
	ConnectionID        string `json:"connection_id"`
	ConnectionProtocol  string `json:"connection_protocol"`
	CreatedAt           string `json:"created_at"`
	ErrorMsg            string `json:"error_msg"`
	InboundConnectionID string `json:"inbound_connection_id"`
	InvitationKey       string `json:"invitation_key"`
	InvitationMode      string `json:"invitation_mode"`
	InvitationMessageID string `json:"invitation_msg_id"`
	MyDID               string `json:"my_did"`
	RequestID           string `json:"request_id"`
	RFC23State          string `json:"rfc23_state"`
	RoutingState        string `json:"routing_state"`
	State               string `json:"state"`
	TheirDID            string `json:"their_did"`
	TheirLabel          string `json:"their_label"`
	TheirPublicDID      string `json:"their_public_did"`
	TheirRole           string `json:"their_role"`
	UpdatedAt           string `json:"updated_at"`
}

// Data is a struct part of RequestsAttach.
type Data struct {
	Base64 string   `json:"base64"`
	JSON   JSON     `json:"json"`
	Jws    Jws      `json:"jws"`
	Links  []string `json:"links"`
	Sha256 string   `json:"sha256"`
}

// Header is a struct part of RequestsAttach.
type Header struct {
	Kid string `json:"kid"`
}

// Invitation is a struct used to hold invitation data.
// It's used as receive invitation request body object.
type InvitationV1 struct {
	ID              string   `json:"@id,omitempty"`
	Type            string   `json:"@type,omitempty"`
	DID             string   `json:"did,omitempty"`
	ImageURL        string   `json:"imageUrl,omitempty"`
	Label           string   `json:"label,omitempty"`
	RecipientKeys   []string `json:"recipientKeys,omitempty"`
	RoutingKeys     []string `json:"routingKeys,omitempty"`
	ServiceEndpoint string   `json:"serviceEndpoint,omitempty"`
}

// InvitationV2 is a struct used to hold invitation data.
// It's used as receive invitation v2 request body object.
type InvitationV2 struct {
	ID                 string           `json:"@id"`
	Type               string           `json:"@type"`               // did:sov:BzCbsNYhMrjHiqZDTUASHg;spec/out-of-band/1.0/invitation
	HandshakeProtocols []string         `json:"handshake_protocols"` // did:sov:BzCbsNYhMrjHiqZDTUASHg;spec/didexchange/v1.0
	Label              string           `json:"label"`
	RequestsAttach     []RequestsAttach `json:"requests~attach"`
	Services           []Services       `json:"services"`
}

// JSON is a struct part of RequestsAttach.
type JSON struct {
	Sample string `json:"sample"`
}

// Jws is a struct part of RequestsAttach.
type Jws struct {
	Header     Header       `json:"header"`
	Protected  string       `json:"protected"`
	Signature  string       `json:"signature"`
	Signatures []Signatures `json:"signatures"`
}

// Metadata is a struct used to hold metadata data.
type Metadata struct{}

// RequestsAttach is a struct used to hold requests attach data.
type RequestsAttach struct {
	ID          string `json:"@id"`
	ByteCount   int64  `json:"byte_count"`
	Data        Data   `json:"data"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	LastmodTime string `json:"lastmod_time"`
	MimeType    string `json:"mime-type"`
}

// Signatures is a struct part of RequestsAttach.
type Signatures struct {
	Header    Header `json:"header"`
	Protected string `json:"protected"`
	Signature string `json:"signature"`
}

// Services is a struct used to hold services data
type Services struct {
	DID             string   `json:"did"`
	ID              string   `json:"id"`
	RecipientKeys   []string `json:"recipientKeys"`
	RoutingKeys     []string `json:"routingKeys"`
	ServiceEndpoint string   `json:"serviceEndpoint"`
	Type            string   `json:"type"` // did-communication
}
