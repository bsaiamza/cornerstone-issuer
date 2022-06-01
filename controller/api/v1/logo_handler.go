package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/server"
)

func getIamzaLogo(config *config.Config) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(getIamzaLogoHandler(config), mdw...)
}
func getIamzaLogoHandler(config *config.Config) http.HandlerFunc {
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

		log.Info.Print("Retrieving Iamza logo...")

		logo, err := ioutil.ReadFile(config.GetServerAddress() +"/static/media/"+ config.GetLogoName())
		if err != nil {
			log.Error.Printf("Failed to get logo: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get logo: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Println("Logo retrieved successfully!")

		w.Header().Set("Content-Type", "image/png")
		w.Write(logo)
	}
}
