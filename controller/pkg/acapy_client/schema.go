package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// CreateSchema creates a schema.
func (c *Client) CreateSchema(request models.CreateSchemaRequest, params *models.CreateSchemaParams) (models.CreateSchemaResponse, error) {
	var schema models.CreateSchemaResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"conn_id":                         params.ConnID,
			"create_transaction_for_endorser": strconv.FormatBool(params.CreateTransactionForEndorser),
		}
	}

	err := c.post("/schemas", queryParams, request, &schema)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /schemas: %s", err.Error())
		return models.CreateSchemaResponse{}, err
	}
	return schema, nil
}

// ListSchemas returns all schemas.
func (c *Client) ListSchemas(params *models.LitSchemasParams) (models.ListSchemasResponse, error) {
	var schemas models.ListSchemasResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"schema_id":         params.SchemaID,
			"schema_issuer_did": params.SchemaIssuerDID,
			"schema_name":       params.SchemaName,
			"schema_version":    params.SchemaVersion,
		}
	}

	err := c.get("/schemas/created", queryParams, &schemas)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /schemas/created: %s", err.Error())
		return models.ListSchemasResponse{}, err
	}
	return schemas, nil
}

// GetSchema returns a schema.
func (c *Client) GetSchema(schemaID string) (models.GetSchemaResponse, error) {
	var schema models.GetSchemaResponse

	err := c.get(fmt.Sprintf("/schemas/%s", schemaID), nil, &schema)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /schemas/{schema_id}: %s", err.Error())
		return models.GetSchemaResponse{}, err
	}
	return schema, nil
}
