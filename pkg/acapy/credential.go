package acapy

import (
	"cornerstone-issuer/pkg/log"
	"cornerstone-issuer/pkg/models"
)

func (c *Client) IssueCredential(request models.IssueCredentialRequest) (models.IssueCredentialResponse, error) {
	var credential models.IssueCredentialResponse

	arg := models.AcapyPostRequest{
		Endpoint: "/issue-credential/send",
		Body:     request,
		Response: &credential,
	}

	err := c.post(arg)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential/send: %s", err.Error())
		return models.IssueCredentialResponse{}, err
	}
	return credential, nil
}
