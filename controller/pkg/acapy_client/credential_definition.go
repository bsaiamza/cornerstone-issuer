package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// CreateCredentialDefinition is used to create a credential definition.
func (c *Client) CreateCredentialDefinition(request models.CreateCredentialDefinitionRequest, params *models.CreateCredentialDefinitionParams) (models.CreateCredentialDefinitionResponse, error) {
	var createCredentialDefinitionResponse models.CreateCredentialDefinitionResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"conn_id":                         params.ConnID,
			"create_transaction_for_endorser": strconv.FormatBool(params.CreateTransactionForEndorser),
		}
	}

	err := c.post("/credential-definitions", queryParams, request, &createCredentialDefinitionResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /credential-definitions: ", err)
		return models.CreateCredentialDefinitionResponse{}, err
	}
	return createCredentialDefinitionResponse, nil
}

// QueryCredentialDefinitions returns all credential definitions.
func (c *Client) QueryCredentialDefinitions(params *models.QueryCredentialDefinitionsParams) (models.QueryCredentialDefinitionsResponse, error) {
	var queryCredentialDefinitionsResponse models.QueryCredentialDefinitionsResponse

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

	err := c.get("/credential-definitions/created", queryParams, &queryCredentialDefinitionsResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /credential-definitions/created: ", err)
		return models.QueryCredentialDefinitionsResponse{}, err
	}
	return queryCredentialDefinitionsResponse, nil
}

// GetCredentialDefinition returns a credential definition.
func (c *Client) GetCredentialDefinition(credDefID string) (models.GetCredentialDefinitionResponse, error) {
	var credentialDefinition models.GetCredentialDefinitionResponse

	err := c.get(fmt.Sprintf("/credential-definitions/%s", credDefID), nil, &credentialDefinition)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /credential-definitions/{cred_def_id}: ", err)
		return models.GetCredentialDefinitionResponse{}, err
	}
	return credentialDefinition, nil
}
