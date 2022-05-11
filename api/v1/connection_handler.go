package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/server"
	// "github.com/skip2/go-qrcode"
)

func createInvitationV1(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(createInvitationV1Handler(config, c), mdw...)
}
func createInvitationV1Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Creating invitation...")

		alias := r.URL.Query().Get("alias")
		autoAccept, _ := strconv.ParseBool(r.URL.Query().Get("auto_accept"))
		multiuse, _ := strconv.ParseBool(r.URL.Query().Get("multi_use"))
		public, _ := strconv.ParseBool(r.URL.Query().Get("public"))

		queryParams := models.CreateInvitationV1Params{
			Alias:      alias,
			AutoAccept: autoAccept,
			MultiUse:   multiuse,
			Public:     public,
		}

		var createInvitationRequest models.CreateInvitationV1Request
		err := json.NewDecoder(r.Body).Decode(&createInvitationRequest)
		switch {
		case err == io.EOF:
			log.Warning.Print("Empty body!")
		case err != nil:
			log.Warning.Print("Failed to parse body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to parse body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		invitation, err := c.CreateInvitationV1(createInvitationRequest, &queryParams)
		if err != nil {
			log.Warning.Print("Failed to create invitation: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		// if err != nil {
		// 	log.Warning.Print("Failed to create QR code: ", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	res := server.Res{
		// 		"success": false,
		// 		"msg":     "Failed to create QR code: " + err.Error(),
		// 	}
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }

		log.Info.Print("Invitation created!")

		// w.Write(qrCodePng)
		// w.Header().Set("Content-Type", "image/png")

		res := server.Res{
			"success":    true,
			"msg":        "Invitation created!",
			"invitation": invitation,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func receiveInvitationV1(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(receiveInvitationV1Handler(config, c), mdw...)
}
func receiveInvitationV1Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Receiving invitation...")

		alias := r.URL.Query().Get("alias")
		autoAccept, _ := strconv.ParseBool(r.URL.Query().Get("auto_accept"))
		mediationID := r.URL.Query().Get("mediation_id")

		queryParams := models.ReceiveInvitationV1Params{
			Alias:       alias,
			AutoAccept:  autoAccept,
			MediationID: mediationID,
		}

		var receiveInvitationRequest models.InvitationV1
		if err := json.NewDecoder(r.Body).Decode(&receiveInvitationRequest); err != nil {
			log.Error.Printf("Failed to decode request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		invitation, err := c.ReceiveInvitationV1(receiveInvitationRequest, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to receive invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to receive invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Invitation received!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(invitation)
	}
}

func acceptInvitationV1(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(acceptInvitationV1Handler(config, c), mdw...)
}
func acceptInvitationV1Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Accepting invitation...")

		connID := r.URL.Query().Get("conn_id")
		mediationID := r.URL.Query().Get("mediation_id")
		myEndpoint := r.URL.Query().Get("my_endpoint")
		myLabel := r.URL.Query().Get("my_label")

		queryParams := models.AcceptInvitationV1Params{
			MediationID: mediationID,
			MyEndpoint:  myEndpoint,
			MyLabel:     myLabel,
		}

		invitation, err := c.AcceptInvitationV1(connID, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to accept invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to accept invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Invitation accepted!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(invitation)
	}
}

func acceptRequestV1(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(acceptRequestV1Handler(config, c), mdw...)
}
func acceptRequestV1Handler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Accepting request...")

		connID := r.URL.Query().Get("conn_id")
		myEndpoint := r.URL.Query().Get("my_endpoint")

		queryParams := models.AcceptRequestV1Params{
			MyEndpoint: myEndpoint,
		}

		request, err := c.AcceptRequestV1(connID, &queryParams)
		if err != nil {
			log.Error.Printf("Failed to accept request: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to accept request: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Request accepted!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(request)

		log.Info.Print("Sending a trust ping...")

		_, err = c.SendPing(connID)
		if err != nil {
			log.Error.Printf("Failed to send ping: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to send ping: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
}

func queryConnections(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(queryConnectionsHandler(config, c), mdw...)
}
func queryConnectionsHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Querying connections...")

		alias := r.URL.Query().Get("alias")
		connectionProtocol := r.URL.Query().Get("connection_protocol")
		invitationKey := r.URL.Query().Get("invitation_key")
		myDID := r.URL.Query().Get("my_did")
		state := r.URL.Query().Get("state")
		theirDID := r.URL.Query().Get("their_did")
		theirPublicDID := r.URL.Query().Get("their_public_did")
		theirRole := r.URL.Query().Get("their_role")

		queryParams := models.QueryConnectionsParams{
			Alias:              alias,
			ConnectionProtocol: connectionProtocol,
			InvitationKey:      invitationKey,
			MyDID:              myDID,
			State:              state,
			TheirDID:           theirDID,
			TheirPublicDID:     theirPublicDID,
			TheirRole:          theirRole,
		}

		connections, err := c.QueryConnections(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to query connections: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to query connections: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Connections queried!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(connections)
	}
}

func getConnection(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getConnectionHandler(config, c), mdw...)
}
func getConnectionHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Getting connection...")

		connID := r.URL.Query().Get("conn_id")

		connection, err := c.GetConnection(connID)
		if err != nil {
			log.Error.Printf("Failed to get connection: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get connection: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Connection gotten!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(connection)
	}
}

func removeConnection(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(removeConnectionHandler(config, c), mdw...)
}
func removeConnectionHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "DELETE, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodDelete {
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

		log.Info.Print("Removing connection...")

		connID := r.URL.Query().Get("conn_id")

		err := c.RemoveConnection(connID)
		if err != nil {
			log.Error.Printf("Failed to remove connection: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to remove connection: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Connection removed!")

		w.WriteHeader(http.StatusOK)
	}
}

func sendMessage(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(sendMessageHandler(config, c), mdw...)
}
func sendMessageHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		log.Info.Print("Sending message...")

		connID := r.URL.Query().Get("conn_id")
		var request = struct {
			Content string `json:"content"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Error.Printf("Failed to decode request: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		err = c.SendBasicMessage(connID, request.Content)
		if err != nil {
			log.Error.Printf("Failed to send message: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to send message: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Message sent!")

		w.WriteHeader(http.StatusOK)
	}
}
