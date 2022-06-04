package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// CreateCredentialDefinition creates a credential definition.
func (c *Client) CreateCredentialDefinition(request models.CreateCredentialDefinitionRequest, params *models.CreateCredentialDefinitionParams) (models.CreateCredentialDefinitionResponse, error) {
	var credentialDefinition models.CreateCredentialDefinitionResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"conn_id":                         params.ConnID,
			"create_transaction_for_endorser": strconv.FormatBool(params.CreateTransactionForEndorser),
		}
	}

	err := c.post("/credential-definitions", queryParams, request, &credentialDefinition)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /credential-definitions: %s", err.Error())
		return models.CreateCredentialDefinitionResponse{}, err
	}
	return credentialDefinition, nil
}

// ListCredentialDefinitions returns all credential definitions.
func (c *Client) ListCredentialDefinitions(params *models.ListCredentialDefinitionsParams) (models.ListCredentialDefinitionsResponse, error) {
	var credentialDefinitions models.ListCredentialDefinitionsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"cred_def_id":       params.CredDefID,
			"issuer_did":        params.IssuerDID,
			"schema_id":         params.SchemaID,
			"schema_issuer_did": params.SchemaIssuerDID,
			"schema_name":       params.SchemaName,
			"schema_version":    params.SchemaVersion,
		}
	}

	err := c.get("/credential-definitions/created", queryParams, &credentialDefinitions)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /credential-definitions/created: %s", err.Error())
		return models.ListCredentialDefinitionsResponse{}, err
	}
	return credentialDefinitions, nil
}

// GetCredentialDefinition returns a credential definition.
// * this function will be specific to the cornerstone credential definition or alike for other definitions.
func (c *Client) GetCredentialDefinition(credDefID string) (models.GetCredentialDefinitionResponse, error) {
	var credentialDefinition models.GetCredentialDefinitionResponse

	err := c.get(fmt.Sprintf("/credential-definitions/%s", credDefID), nil, &credentialDefinition)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /credential-definitions/{cred_def_id}: %s", err.Error())
		return models.GetCredentialDefinitionResponse{}, err
	}
	return credentialDefinition, nil
}
