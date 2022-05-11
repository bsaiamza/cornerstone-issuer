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

func createCredentialDefinition(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(createCredentialDefinitionHandler(config, c), mdw...)
}
func createCredentialDefinitionHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Creating credential definition...")

		connID := r.URL.Query().Get("conn_id")
		createTransactionForEndorser, _ := strconv.ParseBool(r.URL.Query().Get("create_transaction_for_endorser"))

		queryParams := models.CreateCredentialDefinitionParams{
			ConnID:                       connID,
			CreateTransactionForEndorser: createTransactionForEndorser,
		}

		var createCredentialDefinitionRequest models.CreateCredentialDefinitionRequest

		if err := json.NewDecoder(r.Body).Decode(&createCredentialDefinitionRequest); err != nil {
			log.Error.Printf("Failed to decode request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		credentialDefinition, err := c.CreateCredentialDefinition(createCredentialDefinitionRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to create credential definition: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to create credential definition: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential definition created!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(credentialDefinition.Sent.CredentialDefinitionID)
	}
}

func queryCredentialDefinitions(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(queryCredentialDefinitionsHandler(config, c), mdw...)
}
func queryCredentialDefinitionsHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Querying credential definitions...")

		credDefID := r.URL.Query().Get("cred_def_id")
		issuerDID := r.URL.Query().Get("issuer_did")
		schemaID := r.URL.Query().Get("schema_id")
		schemaIssuerDID := r.URL.Query().Get("schema_issuer_did")
		schemaName := r.URL.Query().Get("schema_name")
		schemaVersion := r.URL.Query().Get("schema_version")

		queryParams := models.QueryCredentialDefinitionsParams{
			CredDefID:       credDefID,
			IssuerDID:       issuerDID,
			SchemaID:        schemaID,
			SchemaIssuerDID: schemaIssuerDID,
			SchemaName:      schemaName,
			SchemaVersion:   schemaVersion,
		}

		credentialDefinitions, err := c.QueryCredentialDefinitions(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to query credential definitions: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to query credential definitions: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		var credentialDefinitionIDs []models.CredentialDefinition
		for _, credentialDefinition := range credentialDefinitions.CredentialDefinitionIDs {
			ids := models.CredentialDefinition{
				ID: credentialDefinition,
			}

			credentialDefinitionIDs = append(credentialDefinitionIDs, ids)
		}

		log.Info.Print("Credential definitions queried!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(credentialDefinitionIDs)
	}
}

func getCredentialDefinition(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getCredentialDefinitionHandler(config, c), mdw...)
}
func getCredentialDefinitionHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Getting credential definition...")

		credDefID := r.URL.Query().Get("cred_def_id")

		credentialDefinition, err := c.GetCredentialDefinition(credDefID)
		if err != nil {
			log.Error.Printf("Failed to get credential definition: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get credential definition: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential definition gotten!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(credentialDefinition.CredentialDefinition)
	}
}
