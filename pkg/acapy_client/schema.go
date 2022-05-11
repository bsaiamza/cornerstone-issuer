package acapy

import (
	"fmt"
	"strconv"

	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// CreateSchema is used to create a schema.
func (c *Client) CreateSchema(request models.CreateSchemaRequest, params *models.CreateSchemaParams) (models.CreateSchemaResponse, error) {
	var createSchemaResponse models.CreateSchemaResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"conn_id":                         params.ConnID,
			"create_transaction_for_endorser": strconv.FormatBool(params.CreateTransactionForEndorser),
		}
	}

	err := c.post("/schemas", queryParams, request, &createSchemaResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /schemas: ", err)
		return models.CreateSchemaResponse{}, err
	}
	return createSchemaResponse, nil
}

// QuerySchemas returns all schemas.
func (c *Client) QuerySchemas(params *models.QuerySchemasParams) (models.QuerySchemasResponse, error) {
	var querySchemasResponse models.QuerySchemasResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"schema_id":         params.SchemaID,
			"schema_issuer_did": params.SchemaIssuerDID,
			"schema_name":       params.SchemaName,
			"schema_version":    params.SchemaVersion,
		}
	}

	err := c.get("/schemas/created", queryParams, &querySchemasResponse)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /schemas/created: ", err)
		return models.QuerySchemasResponse{}, err
	}
	return querySchemasResponse, nil
}

// GetSchema returns a schema.
func (c *Client) GetSchema(schemaID string) (models.GetSchemaResponse, error) {
	var schema models.GetSchemaResponse

	err := c.get(fmt.Sprintf("/schemas/%s", schemaID), nil, &schema)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /schemas/{schema_id}: ", err)
		return models.GetSchemaResponse{}, err
	}
	return schema, nil
}
