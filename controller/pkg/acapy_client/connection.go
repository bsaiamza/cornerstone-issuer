package acapy

import (
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// CreateInvitation creates an invitation.
func (c *Client) CreateInvitation(request models.CreateInvitationRequest, params *models.CreateInvitationParams) (models.CreateInvitationResponse, error) {
	var invitation models.CreateInvitationResponse

	if params.Alias == "" {
		params.Alias = request.MyLabel
	}

	if !params.AutoAccept {
		params.AutoAccept = true
	}

	if !params.MultiUse {
		params.MultiUse = true
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":       params.Alias,
			"auto_accept": strconv.FormatBool(params.AutoAccept),
			"multi_use":   strconv.FormatBool(params.MultiUse),
			"public":      strconv.FormatBool(params.Public),
		}
	}

	err := c.post("/connections/create-invitation", queryParams, request, &invitation)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/create-invitation: %s", err.Error())
		return models.CreateInvitationResponse{}, err
	}
	return invitation, nil
}

// ReceiveInvitation receives an invitation.
func (c *Client) ReceiveInvitation(request models.Invitation, params *models.ReceiveInvitationParams) (models.ConnectionResponse, error) {
	var invitation models.ConnectionResponse

	if params.Alias == "" {
		params.Alias = request.Label
	}

	if !params.AutoAccept {
		params.AutoAccept = true
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":        params.Alias,
			"auto_accept":  strconv.FormatBool(params.AutoAccept),
			"mediation_id": params.MediationID,
		}
	}

	err := c.post("/connections/receive-invitation", queryParams, request, &invitation)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /connections/receive-invitation: %s", err.Error())
		return models.ConnectionResponse{}, err
	}
	return invitation, nil
}

// ListConnections returns all connections.
func (c *Client) ListConnections(params *models.ListConnectionsParams) (models.ListConnectionsResponse, error) {
	var connections models.ListConnectionsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":               params.Alias,
			"connection_protocol": params.ConnectionProtocol,
			"invitation_key":      params.InvitationKey,
			"my_did":              params.MyDID,
			"connection_state":    params.State,
			"their_did":           params.TheirDID,
			"their_public_did":    params.TheirPublicDID,
			"their_role":          params.TheirRole,
		}
	}

	err := c.get("/connections", queryParams, &connections)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections: ", err)
		return models.ListConnectionsResponse{}, err
	}
	return connections, nil
}
