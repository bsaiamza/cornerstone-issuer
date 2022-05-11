package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// QueryConnections returns all connections.
func (c *Client) QueryConnections(params *models.QueryConnectionsParams) ([]models.Connection, error) {
	var queryConnectionsResponse models.QueryConnectionsResponse

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

	err := c.get("/connections", queryParams, &queryConnectionsResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections: ", err)
		return nil, err
	}
	return queryConnectionsResponse.Results, nil
}

// GetConnection returns a single connection.
func (c *Client) GetConnection(connID string) (models.Connection, error) {
	var connection models.Connection

	err := c.get(fmt.Sprintf("/connections/%s", connID), nil, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/{conn_id}: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// RemoveConnection removes a connection.
func (c *Client) RemoveConnection(connID string) error {
	return c.delete(fmt.Sprintf("/connections/%s", connID))
}

// Thread ...
type Thread struct {
	ThreadID string `json:"thread_id"`
}

// SendPing is a Trust Ping.
func (c *Client) SendPing(connID string) (Thread, error) {
	ping := struct {
		Comment string `json:"comment"`
	}{
		Comment: "ping",
	}
	var thread Thread
	err := c.post(fmt.Sprintf("/connections/%s/send-ping", connID), nil, ping, &thread)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/{conn_id}/send-ping: ", err)
		return Thread{}, err
	}
	return thread, nil
}

// SendBasicMessage sends a basic message.
func (c *Client) SendBasicMessage(connectionID string, message string) error {
	type BasicMessage struct {
		Content string `json:"content"`
	}
	var basicMessage = BasicMessage{
		Content: message,
	}

	return c.post(fmt.Sprintf("/connections/%s/send-message", connectionID), nil, basicMessage, nil)
}

// ------------------ \\
// ------- V1 ------- \\
// ------------------ \\

// CreateInvitationV1 is used to create an invitation to connect.
func (c *Client) CreateInvitationV1(request models.CreateInvitationV1Request, params *models.CreateInvitationV1Params) (models.CreateInvitationV1Response, error) {
	var createInvitationResponse models.CreateInvitationV1Response

	if params.Alias == "" {
		params.Alias = request.MyLabel
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

	err := c.post("/connections/create-invitation", queryParams, request, &createInvitationResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/create-invitation: ", err)
		return models.CreateInvitationV1Response{}, err
	}
	return createInvitationResponse, nil
}

// ReceiveInvitationV1 is used to receive an invitation to connect.
func (c *Client) ReceiveInvitationV1(request models.InvitationV1, params *models.ReceiveInvitationV1Params) (models.Connection, error) {
	var connection models.Connection

	if params.Alias == "" {
		params.Alias = request.Label
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":        params.Alias,
			"auto_accept":  strconv.FormatBool(params.AutoAccept),
			"mediation_id": params.MediationID,
		}
	}

	err := c.post("/connections/receive-invitation", queryParams, request, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/receive-invitation: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// AcceptInvitationV1 is used to accept an invitation to connect.
func (c *Client) AcceptInvitationV1(connID string, params *models.AcceptInvitationV1Params) (models.Connection, error) {
	var connection models.Connection

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"mediation_id": params.MediationID,
			"my_endpoint":  params.MyEndpoint,
			"my_label":     params.MyLabel,
		}
	}

	err := c.post(fmt.Sprintf("/connections/%s/accept-invitation", connID), queryParams, nil, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/{conn_id}/accept-invitation: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// AcceptRequestV1 is used to accept a request to connect.
func (c *Client) AcceptRequestV1(connID string, params *models.AcceptRequestV1Params) (models.Connection, error) {
	var connection models.Connection

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"my_endpoint": params.MyEndpoint,
		}
	}

	err := c.post(fmt.Sprintf("/connections/%s/accept-request", connID), queryParams, nil, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /connections/{conn_id}/accept-request: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// ------------------ \\
// ------- V2 ------- \\
// ------------------ \\

// CreateInvitationV2 is used to create an invitation to connect.
func (c *Client) CreateInvitationV2(request models.CreateInvitationV2Request, params *models.CreateInvitationV2Params) (models.CreateInvitationV2Response, error) {
	var createInvitationResponse models.CreateInvitationV2Response

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"auto_accept": strconv.FormatBool(params.AutoAccept),
			"multi_use":   strconv.FormatBool(params.MultiUse),
		}
	}

	err := c.post("/out-of-band/create-invitation", queryParams, request, &createInvitationResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /out-of-band/create-invitation: ", err)
		return models.CreateInvitationV2Response{}, err
	}
	return createInvitationResponse, nil
}

// ReceiveInvitationV2 is used to receive an invitation to connect.
func (c *Client) ReceiveInvitationV2(request models.InvitationV2, params *models.ReceiveInvitationV2Params) (models.Connection, error) {
	var connection models.Connection

	if params.Alias == "" {
		params.Alias = request.Label
	}

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"alias":                   params.Alias,
			"auto_accept":             strconv.FormatBool(params.AutoAccept),
			"mediation_id":            params.MediationID,
			"use_existing_connection": strconv.FormatBool(params.UseExistingConnection),
		}
	}

	err := c.post("/out-of-band/receive-invitation", queryParams, request, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /out-of-band/receive-invitation: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// AcceptInvitationV2 is used to accept an invitation to connect.
func (c *Client) AcceptInvitationV2(connID string, params *models.AcceptInvitationV2Params) (models.Connection, error) {
	var connection models.Connection

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"my_endpoint": params.MyEndpoint,
			"my_label":    params.MyLabel,
		}
	}

	err := c.post(fmt.Sprintf("/didexchange/%s/accept-invitation", connID), queryParams, nil, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py didexchange/{conn_id}/accept-invitation: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}

// AcceptRequestV2 is used to accept a request to connect.
func (c *Client) AcceptRequestV2(connID string, params *models.AcceptRequestV2Params) (models.Connection, error) {
	var connection models.Connection

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"mediation_id": params.MediationID,
			"my_endpoint":  params.MyEndpoint,
		}
	}

	err := c.post(fmt.Sprintf("/didexchange/%s/accept-request", connID), queryParams, nil, &connection)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py didexchange/{conn_id}/accept-request: ", err)
		return models.Connection{}, err
	}
	return connection, nil
}
