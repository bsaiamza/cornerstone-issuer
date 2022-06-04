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

func createCredentialDefinition(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(createCredentialDefinitionHandler(config, acapyClient), mdw...)
}
func createCredentialDefinitionHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		err := json.NewDecoder(r.Body).Decode(&createCredentialDefinitionRequest)
		if err != nil {
			log.Error.Printf("Failed to decode request body: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		definition, err := acapyClient.CreateCredentialDefinition(createCredentialDefinitionRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to create credential definition: %s", err.Error())
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
		json.NewEncoder(w).Encode(definition)
	}
}

func listCredentialDefinitions(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(listCredentialDefinitionsHandler(config, acapyClient), mdw...)
}
func listCredentialDefinitionsHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Listing credential definitions...")

		credDefID := r.URL.Query().Get("cred_def_id")
		issuerDID := r.URL.Query().Get("issuer_did")
		schemaID := r.URL.Query().Get("schema_id")
		schemaIssuerDID := r.URL.Query().Get("schema_issuer_did")
		schemaName := r.URL.Query().Get("schema_name")
		schemaVersion := r.URL.Query().Get("schema_version")

		queryParams := models.ListCredentialDefinitionsParams{
			CredDefID:       credDefID,
			IssuerDID:       issuerDID,
			SchemaID:        schemaID,
			SchemaIssuerDID: schemaIssuerDID,
			SchemaName:      schemaName,
			SchemaVersion:   schemaVersion,
		}

		definitions, err := acapyClient.ListCredentialDefinitions(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to list credential definitions: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to list credential definitions: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential definitions listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(definitions)
	}
}

func getCredentialDefinition(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getCredentialDefinitionHandler(config, acapyClient), mdw...)
}
func getCredentialDefinitionHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		log.Info.Println("Retrieving credential definition...")

		credDefID := r.URL.Query().Get("cred_def_id")

		definition, err := acapyClient.GetCredentialDefinition(credDefID)
		if err != nil {
			log.Error.Printf("Failed to get credential definition: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get credential definition: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Println("Credential definition retrieved successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(definition.CredentialDefinition)
	}
}
