package api

import (
	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"net/http"
)

func NewRouter(config *config.Config, client *acapy.Client) *http.ServeMux {
	r := http.NewServeMux()

	// health check
	r.HandleFunc("/api/v1/cornerstone/issuer/healthz", healthz(config, client))
	// schema routes
	r.HandleFunc("/api/v1/cornerstone/issuer/schema/create", createSchema(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/schemas", querySchemas(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/schema", getSchema(config, client))
	// credential definitions routes
	r.HandleFunc("/api/v1/cornerstone/issuer/credential-definition/create", createCredentialDefinition(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/credential-definitions", queryCredentialDefinitions(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/credential-definition", getCredentialDefinition(config, client))
	// connection routes
	r.HandleFunc("/api/v1/cornerstone/issuer/connections", queryConnections(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/connection", getConnection(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/connection/remove", removeConnection(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/connection/send-message", sendMessage(config, client))
	// connection v1 routes
	r.HandleFunc("/api/v1/cornerstone/issuer/connection/v1/create-invitation", createInvitationV1(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/connection/v1/accept-request", acceptRequestV1(config, client))
	// connection v2 routes //TODO: v2 oob & didexchange connection
	// credential routes
	r.HandleFunc("/api/v1/cornerstone/issuer/credential/v2/proposals", queryProposalsV2(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/credential/v2/offer-proposal", sendOfferV2(config, client))
	r.HandleFunc("/api/v1/cornerstone/issuer/credential/v2/issue", issueCredentialV2(config, client))

	return r
}
