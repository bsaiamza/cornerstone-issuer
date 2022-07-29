package api

import (
	"cornerstone_issuer/pkg/client"
	"cornerstone_issuer/pkg/config"

	"github.com/gorilla/mux"
)

func NewRouter(config *config.Config, client *client.Client) *mux.Router {
	r := mux.NewRouter()

	path := r.PathPrefix("/api/v1/dha").Subrouter()
	path.HandleFunc("/user", GetDHAUser(config, client))
	path.HandleFunc("/simulator/user", GetDHASimulatorUser(config, client))

	return r
}
