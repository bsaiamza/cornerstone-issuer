package acapy

import (
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"fmt"
)

// ------------------ \\
// ------- V1 ------- \\
// ------------------ \\

// SendCredentialOfferV1 is used to send a proposal offer.
func (c *Client) SendCredentialOfferV1(credExID string, request models.SendOfferV1Request) (models.IssueCredentialResponse, error) {
	var credentialOffer models.IssueCredentialResponse

	err := c.post(fmt.Sprintf("/issue-credential/records/%s/send-offer", credExID), nil, request, &credentialOffer)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/send-offer: ", err)
		return models.IssueCredentialResponse{}, err
	}
	return credentialOffer, nil
}

// IssueCredentialV1 is used to issue a credential.
func (c *Client) IssueCredentialV1(credExID string, request models.IssueCredentialRequest) (models.IssueCredentialResponse, error) {
	var credential models.IssueCredentialResponse

	err := c.post(fmt.Sprintf("/issue-credential/records/%s/issue", credExID), nil, request, &credential)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/issue: ", err)
		return models.IssueCredentialResponse{}, err
	}
	return credential, nil
}

// QueryCredentialProposalsV1 returns all credential proposals.
func (c *Client) QueryCredentialProposalsV1(params *models.QueryExchangeRecordsParams) (models.QueryExchangeRecordsV1Response, error) {
	var queryCredentialProposalsResponse models.QueryExchangeRecordsV1Response

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"connection_id": params.ConnectionID,
			"role":          params.Role,
			"state":         params.State,
			"thread_id":     params.ThreadID,
		}
	}

	err := c.get("/issue-credential/records", queryParams, &queryCredentialProposalsResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records: ", err)
		return models.QueryExchangeRecordsV1Response{}, err
	}
	return queryCredentialProposalsResponse, nil
}

// ------------------ \\
// ------- V2 ------- \\
// ------------------ \\

// SendCredentialOfferV2 is used to send a proposal offer.
func (c *Client) SendCredentialOfferV2(credExID string, request models.SendOfferV2Request) (models.IssueCredentialV2Response, error) {
	var credentialOffer models.IssueCredentialV2Response

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/send-offer", credExID), nil, request, &credentialOffer)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential-2.0/records/{cred_ex_id}/send-offer: ", err)
		return models.IssueCredentialV2Response{}, err
	}
	return credentialOffer, nil
}

// IssueCredentialV2 is used to issue a credential.
func (c *Client) IssueCredentialV2(credExID string, request models.IssueCredentialRequest) (models.IssueCredentialV2Response, error) {
	var credential models.IssueCredentialV2Response

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/issue", credExID), nil, request, &credential)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/issue: ", err)
		return models.IssueCredentialV2Response{}, err
	}
	return credential, nil
}

// QueryCredentialProposalsV2 returns all credential proposals.
func (c *Client) QueryCredentialProposalsV2(params *models.QueryExchangeRecordsParams) (models.QueryExchangeRecordsV2Response, error) {
	var queryCredentialProposalsResponse models.QueryExchangeRecordsV2Response

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"connection_id": params.ConnectionID,
			"role":          params.Role,
			"state":         params.State,
			"thread_id":     params.ThreadID,
		}
	}

	err := c.get("/issue-credential-2.0/records", queryParams, &queryCredentialProposalsResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential-2.0/records: ", err)
		return models.QueryExchangeRecordsV2Response{}, err
	}
	return queryCredentialProposalsResponse, nil
}
