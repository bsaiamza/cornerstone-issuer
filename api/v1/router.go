package api

import (
	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"net/http"
)

func NewRouter(config *config.Config, client *acapy.Client) *http.ServeMux {
	r := http.NewServeMux()

	// health check
	r.HandleFunc("/api/v1/issuer/healthz", healthz(config, client))
	// schema routes
	r.HandleFunc("/api/v1/issuer/create-schema", createSchema(config, client))
	r.HandleFunc("/api/v1/issuer/schemas", querySchemas(config, client))
	r.HandleFunc("/api/v1/issuer/schema", getSchema(config, client))
	// credential definitions routes
	r.HandleFunc("/api/v1/issuer/create-credential-definition", createCredentialDefinition(config, client))
	r.HandleFunc("/api/v1/issuer/credential-definitions", queryCredentialDefinitions(config, client))
	r.HandleFunc("/api/v1/issuer/credential-definition", getCredentialDefinition(config, client))
	// connection routes
	r.HandleFunc("/api/v1/issuer/connections", queryConnections(config, client))
	r.HandleFunc("/api/v1/issuer/connection", getConnection(config, client))
	r.HandleFunc("/api/v1/issuer/remove-connection", removeConnection(config, client))
	r.HandleFunc("/api/v1/issuer/send-message", sendMessage(config, client))
	// connection v1 routes
	r.HandleFunc("/api/v1/issuer/create-v1", createInvitationV1(config, client))
	r.HandleFunc("/api/v1/issuer/receive-v1", receiveInvitationV1(config, client))
	r.HandleFunc("/api/v1/issuer/accept-invitation-v1", acceptInvitationV1(config, client))
	r.HandleFunc("/api/v1/issuer/accept-request-v1", acceptRequestV1(config, client))
	// connection v2 routes
	// r.HandleFunc("/api/v1/issuer/create-v2", createInvitationV2(config, client))
	// r.HandleFunc("/api/v1/issuer/receive-v2", receiveInvitationV2(config, client))
	// r.HandleFunc("/api/v1/issuer/accept-invitation-v2", acceptInvitationV2(config, client))
	// r.HandleFunc("/api/v1/issuer/accept-request-v2", acceptRequestV2(config, client))
	// credential routes
	r.HandleFunc("/api/v1/issuer/proposals", queryProposalsV2(config, client))
	r.HandleFunc("/api/v1/issuer/proposal", sendProposalV2(config, client))
	r.HandleFunc("/api/v1/issuer/offer", sendOfferV2(config, client))
	r.HandleFunc("/api/v1/issuer/request", sendRequestV2(config, client))
	r.HandleFunc("/api/v1/issuer/issue", issueV2(config, client))

	return r
}
