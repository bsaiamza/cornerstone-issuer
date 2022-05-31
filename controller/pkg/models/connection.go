package models

type CreateInvitationRequest struct {
	MediationID     string   `json:"mediation_id,omitempty"`
	Metadata        Metadata `json:"metadata,omitempty"`
	MyLabel         string   `json:"my_label,omitempty"`
	RecipientKeys   []string `json:"recipient_keys,omitempty"`
	RoutingKeys     []string `json:"routing_keys,omitempty"`
	ServiceEndpoint string   `json:"service_endpoint,omitempty"`
}

type Metadata struct {
}

type CreateInvitationParams struct {
	Alias      string `json:"alias"`
	AutoAccept bool   `json:"auto_accept"`
	MultiUse   bool   `json:"multi_use"`
	Public     bool   `json:"public"`
}

type CreateInvitationResponse struct {
	ConnectionID  string     `json:"connection_id"`
	Invitation    Invitation `json:"invitation"`
	InvitationURL string     `json:"invitation_url"`
}

type Invitation struct {
	ID              string   `json:"@id,omitempty"`
	Did             string   `json:"did,omitempty"`
	ImageURL        string   `json:"imageUrl,omitempty"`
	Label           string   `json:"label,omitempty"`
	RecipientKeys   []string `json:"recipientKeys,omitempty"`
	RoutingKeys     []string `json:"routingKeys,omitempty"`
	ServiceEndpoint string   `json:"serviceEndpoint,omitempty"`
}

type ReceiveInvitationParams struct {
	Alias       string `json:"alias"`
	AutoAccept  bool   `json:"auto_accept"`
	MediationID string `json:"mediation_id"`
}

type ConnectionResponse struct {
	Accept             string `json:"accept"`
	Alias              string `json:"alias"`
	ConnectionID       string `json:"connection_id"`
	ConnectionProtocol string `json:"connection_protocol"`
	CreatedAt          string `json:"created_at"`
	ErrorMsg           string `json:"error_msg"`
	InboundConnectID   string `json:"inbound_connect_id"`
	InvitationKey      string `json:"invitation_key"`
	InvitationMode     string `json:"invitation_mode"`
	InvitationMsgID    string `json:"invitation_msg_id"`
	MyDID              string `json:"my_did"`
	RequestID          string `json:"request_id"`
	Rfc23State         string `json:"rfc23_state"`
	RoutingState       string `json:"routing_state"`
	State              string `json:"state"`
	TheirDID           string `json:"their_did"`
	TheirLabel         string `json:"their_label"`
	TheirPublicDID     string `json:"their_public_did"`
	TheirRole          string `json:"their_role"`
	UpdatedAt          string `json:"updated_at"`
}

type ListConnectionsParams struct {
	Alias              string `json:"alias"`
	ConnectionProtocol string `json:"connection_protocol"`
	InvitationKey      string `json:"invitation_key"`
	MyDID              string `json:"my_did"`
	State              string `json:"state"`
	TheirDID           string `json:"their_did"`
	TheirPublicDID     string `json:"their_public_did"`
	TheirRole          string `json:"their_role"`
}

type ListConnectionsResponse struct {
	Results []ConnectionResponse `json:"results"`
}
