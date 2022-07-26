package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"cornerstone_issuer/pkg/client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/utils"

	"github.com/gorilla/mux"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, client *client.Client, cache *utils.BigCache) *mux.Router {
	r := mux.NewRouter()

	path := r.PathPrefix("/api/v1/cornerstone-issuer").Subrouter()
	path.HandleFunc("/health", health(config))
	path.HandleFunc("/connections", listConnections(config, client))
	path.HandleFunc("/credentials", listCredentials(config, client))
	path.HandleFunc("/credential", getCredential(config, client, cache))
	path.HandleFunc("/credential-email", getCredentialByEmail(config, client, cache))
	path.HandleFunc("/topic/{topic}/", webhookEvents(config, client, cache))

	r.PathPrefix("/").Handler(http.FileServer(getFileSystem()))

	return r
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		fmt.Println(err)
	}

	return http.FS(fsys)
}
