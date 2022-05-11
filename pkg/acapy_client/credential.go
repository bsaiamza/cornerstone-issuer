package acapy

import (
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"fmt"
)

// QueryCredentials returns all credentials.
func (c *Client) QueryCredentials(params *models.QueryCredentialsParams) (models.QueryCredentialsResponse, error) {
	var queryCredentialsResponse models.QueryCredentialsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"count": params.Count,
			"start": params.Start,
			"wql":   params.Wql,
		}
	}

	err := c.get("/credentials", queryParams, &queryCredentialsResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /credentials: ", err)
		return models.QueryCredentialsResponse{}, err
	}
	return queryCredentialsResponse, nil
}

// GetCredential returns a credential.
func (c *Client) GetCredential(credentialID string) (models.Credential, error) {
	var credential models.Credential

	err := c.get(fmt.Sprintf("/credential/%s", credentialID), nil, &credential)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /credential/{credential_id}: ", err)
		return models.Credential{}, err
	}
	return credential, nil
}

// ------------------ \\
// ------- V1 ------- \\
// ------------------ \\

// SendCredentialProposalV1 is used to send a proposal.
func (c *Client) SendCredentialProposalV1(credExID string, request models.SendProposalV1Request) (models.IssueCredentialResponse, error) {
	var credentialProposal models.IssueCredentialResponse

	err := c.post("/issue-credential/send-proposal", nil, request, &credentialProposal)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/send-proposal: ", err)
		return models.IssueCredentialResponse{}, err
	}
	return credentialProposal, nil
}

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

// SendCredentialRequestV1 is used to send a credential request.
func (c *Client) SendCredentialRequestV1(credExID string) (models.IssueCredentialResponse, error) {
	var credentialRequest models.IssueCredentialResponse

	err := c.post(fmt.Sprintf("/issue-credential/records/%s/send-request", credExID), nil, nil, &credentialRequest)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/send-request: ", err)
		return models.IssueCredentialResponse{}, err
	}
	return credentialRequest, nil
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

// StoreCredentialV1 is used to issue a credential.
func (c *Client) StoreCredentialV1(credExID string, request models.StoreCredentialRequest) (models.IssueCredentialResponse, error) {
	var store models.IssueCredentialResponse

	err := c.post(fmt.Sprintf("/issue-credential/records/%s/store", credExID), nil, request, &store)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/store: ", err)
		return models.IssueCredentialResponse{}, err
	}
	return store, nil
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

// SendCredentialProposalV2 is used to send a proposal.
func (c *Client) SendCredentialProposalV2(credExID string, request models.SendProposalV2Request) (models.IssueCredentialV2Response, error) {
	var credentialProposal models.IssueCredentialV2Response

	err := c.post("/issue-credential-2.0/send-proposal", nil, request, &credentialProposal)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential-2.0/send-proposal: ", err)
		return models.IssueCredentialV2Response{}, err
	}
	return credentialProposal, nil
}

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

// SendCredentialRequestV2 is used to send a credential request.
func (c *Client) SendCredentialRequestV2(credExID string, request models.SendRequestV2Request) (models.IssueCredentialV2Response, error) {
	var credentialRequest models.IssueCredentialV2Response

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/send-request", credExID), nil, request, &credentialRequest)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential-2.0/records/{cred_ex_id}/send-request: ", err)
		return models.IssueCredentialV2Response{}, err
	}
	return credentialRequest, nil
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

// StoreCredentialV2 is used to issue a credential.
func (c *Client) StoreCredentialV2(credExID string, request models.StoreCredentialRequest) (models.IssueCredentialV2Response, error) {
	var store models.IssueCredentialV2Response

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/store", credExID), nil, request, &store)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/store: ", err)
		return models.IssueCredentialV2Response{}, err
	}
	return store, nil
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
