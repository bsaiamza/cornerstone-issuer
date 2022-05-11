package api

import (
	"encoding/json"
	"net/http"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/server"
)

func queryProposalsV2(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(queryProposalV2Handler(config, c), mdw...)
}
func queryProposalV2Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Querying proposals...")

		connectionID := r.URL.Query().Get("connection_id")
		role := r.URL.Query().Get("role")
		state := r.URL.Query().Get("state")
		threadID := r.URL.Query().Get("thread_id")

		queryParams := models.QueryExchangeRecordsParams{
			ConnectionID: connectionID,
			Role:         role,
			State:        state,
			ThreadID:     threadID,
		}

		proposals, err := c.QueryCredentialProposalsV2(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to query proposals: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to query proposals: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Proposals queried!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(proposals)
	}
}

func sendOfferV2(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(sendOfferV2Handler(config, c), mdw...)
}
func sendOfferV2Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Sending offer...")

		credExID := r.URL.Query().Get("cred_ex_id")

		var sendOfferRequest models.SendOfferV2Request

		err := json.NewDecoder(r.Body).Decode(&sendOfferRequest)
		if err != nil {
			log.Error.Printf("Failed to decode offer: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode offer: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		offer, err := c.SendCredentialOfferV2(credExID, sendOfferRequest)
		if err != nil {
			log.Error.Printf("Failed to send offer: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to send offer: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Offer sent!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(offer)
	}
}

func issueCredentialV2(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(issueCredentialV2Handler(config, c), mdw...)
}
func issueCredentialV2Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Issuing credential...")

		credExID := r.URL.Query().Get("cred_ex_id")

		var issueRequest models.IssueCredentialRequest

		err := json.NewDecoder(r.Body).Decode(&issueRequest)
		if err != nil {
			log.Error.Printf("Failed to decode issue request: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode issue request: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		credential, err := c.IssueCredentialV2(credExID, issueRequest)
		if err != nil {
			log.Error.Printf("Failed to issue credential: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to issue credential: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential issued!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(credential)

	}
}
