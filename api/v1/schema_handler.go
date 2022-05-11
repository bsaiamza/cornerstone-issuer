package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/server"
)

func createSchema(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(createSchemaHandler(config, c), mdw...)
}
func createSchemaHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Print("Creating schema...")

		connID := r.URL.Query().Get("conn_id")
		createTransactionForEndorser, _ := strconv.ParseBool(r.URL.Query().Get("create_transaction_for_endorser"))

		queryParams := models.CreateSchemaParams{
			ConnID:                       connID,
			CreateTransactionForEndorser: createTransactionForEndorser,
		}

		var createSchemaRequest models.CreateSchemaRequest

		if err := json.NewDecoder(r.Body).Decode(&createSchemaRequest); err != nil {
			log.Error.Printf("Failed to decode request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		schema, err := c.CreateSchema(createSchemaRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to create schema: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to create schema: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Schema created!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(schema.Sent.Schema)
	}
}

func querySchemas(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(querySchemasHandler(config, c), mdw...)
}
func querySchemasHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Print("Querying schemas...")

		schemaID := r.URL.Query().Get("schema_id")
		schemaIssuerDID := r.URL.Query().Get("schema_issuer_did")
		schemaName := r.URL.Query().Get("schema_name")
		schemaVersion := r.URL.Query().Get("schema_version")

		queryParams := models.QuerySchemasParams{
			SchemaID:        schemaID,
			SchemaIssuerDID: schemaIssuerDID,
			SchemaName:      schemaName,
			SchemaVersion:   schemaVersion,
		}

		schemas, err := c.QuerySchemas(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to query schemas: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to query schemas: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		var schemaIDs []models.Schema
		for _, schema := range schemas.SchemaIDs {
			ids := models.Schema{
				ID: schema,
			}

			schemaIDs = append(schemaIDs, ids)
		}

		log.Info.Print("Schemas queried!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(schemaIDs)
	}
}

func getSchema(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getSchemaHandler(config, c), mdw...)
}
func getSchemaHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetAcapyURL())
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Print("Getting schema...")

		schemaID := r.URL.Query().Get("schema_id")

		schema, err := c.GetSchema(schemaID)
		if err != nil {
			log.Error.Printf("Failed to get schema: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get schema: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Schema gotten!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(schema.Schema)
	}
}
