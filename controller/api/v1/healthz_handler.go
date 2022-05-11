package api

import (
	"encoding/json"
	"net/http"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/server"
)

func healthz(config *config.Config, c *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(healthzHandler(config, c), mdw...)
}
func healthzHandler(config *config.Config, c *acapy.Client) http.HandlerFunc {
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

		alive, err := c.IsAlive()
		if err != nil {
			log.Error.Printf("Failed to check cornerstone issuer agent health: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to check cornerstone issuer agent health!",
				"error":   err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		if alive {
			log.Info.Print("Cornerstone issuer agent is healthy!")
		}
		log.Info.Print("Cornerstone issuer server is healthy!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(alive)
	}
}
