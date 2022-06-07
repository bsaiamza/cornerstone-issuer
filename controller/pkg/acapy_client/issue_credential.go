package acapy

import (
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// IssueCornerstoneCredential issues a credential.
func (c *Client) IssueCornerstoneCredential(request models.IssueCornerstoneCredentialRequest) (models.IssueCornerstoneCredentialResponse, error) {
	var credential models.IssueCornerstoneCredentialResponse

	err := c.post("/issue-credential/send", nil, request, &credential)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential/send: %s", err.Error())
		return models.IssueCornerstoneCredentialResponse{}, err
	}
	return credential, nil
}

// ListCredentialRecords returns all credential exchange records.
func (c *Client) ListCredentialRecords(params *models.ListCredentialRecordsParams) (models.ListCredentialRecordsResponse, error) {
	var credentialRecords models.ListCredentialRecordsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"connection_id": params.ConnectionID,
			"role":          params.Role,
			"state":         params.State,
			"thread_id":     params.ThreadID,
		}
	}

	err := c.get("/issue-credential/records", queryParams, &credentialRecords)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential/records: %s", err.Error())
		return models.ListCredentialRecordsResponse{}, err
	}
	return credentialRecords, nil
}
