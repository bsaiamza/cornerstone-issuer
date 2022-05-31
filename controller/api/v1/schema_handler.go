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

func createSchema(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(createSchemaHandler(config, acapyClient), mdw...)
}
func createSchemaHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		err := json.NewDecoder(r.Body).Decode(&createSchemaRequest)
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

		schema, err := acapyClient.CreateSchema(createSchemaRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to create schema: %s", err.Error())
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
		json.NewEncoder(w).Encode(schema)
	}
}

func listSchemas(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(listSchemasHandler(config, acapyClient), mdw...)
}
func listSchemasHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Listing schemas...")

		schemaID := r.URL.Query().Get("schema_id")
		schemaIssuerDID := r.URL.Query().Get("schema_issuer_did")
		schemaName := r.URL.Query().Get("schema_name")
		schemaVersion := r.URL.Query().Get("schema_version")

		queryParams := models.LitSchemasParams{
			SchemaID:        schemaID,
			SchemaIssuerDID: schemaIssuerDID,
			SchemaName:      schemaName,
			SchemaVersion:   schemaVersion,
		}

		schemas, err := acapyClient.ListSchemas(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to list schemas: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to list schemas: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Schemas listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(schemas)
	}
}

func getSchema(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getSchemaHandler(config, acapyClient), mdw...)
}
func getSchemaHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
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

		log.Info.Println("Retrieving schema...")

		schemaID := r.URL.Query().Get("schema_id")

		schema, err := acapyClient.GetSchema(schemaID)
		if err != nil {
			log.Error.Printf("Failed to get schema: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get schema: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Println("Schema retrieved successfully!")

		json.NewEncoder(w).Encode(schema)
	}
}

// func getCornerstoneSchema(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
// 	mdw := []server.Middleware{
// 		server.NewLogRequest,
// 	}

// 	return server.ChainMiddleware(getCornerstoneSchemaHandler(config, acapyClient), mdw...)
// }
// func getCornerstoneSchemaHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		header := w.Header()
// 		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
// 		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
// 		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		if r.Method != http.MethodGet {
// 			log.Warning.Print("Incorrect request method!")
// 			w.WriteHeader(http.StatusMethodNotAllowed)
// 			res := server.Res{
// 				"success": false,
// 				"msg":     "Warning: Incorrect request method!",
// 			}
// 			json.NewEncoder(w).Encode(res)
// 			return
// 		}

// 		defer r.Body.Close()

// 		log.Info.Println("Retrieving schema...")

// 		schema, err := acapyClient.GetSchema(config.GetSchemaID())
// 		if err != nil {
// 			log.Error.Printf("Failed to get schema: %s", err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			res := server.Res{
// 				"success": false,
// 				"msg":     "Failed to get schema: " + err.Error(),
// 			}
// 			json.NewEncoder(w).Encode(res)
// 			return
// 		}

// 		log.Info.Println("Schema retrieved successfully!")

// 		json.NewEncoder(w).Encode(schema)
// 	}
// }
