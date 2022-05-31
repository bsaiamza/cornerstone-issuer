package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/cache"
	"cornerstone_issuer/pkg/config"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, acapyClient *acapy.Client, cache *cache.BigCache) *http.ServeMux {
	r := http.NewServeMux()

	apiBaseURL := config.GetAPIBaseURL()

	// health
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/health", health(config))
	// logo
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/logo", getIamzaLogo(config))
	// did
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/did", getDID(config, acapyClient))
	// schema
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schema/create", createSchema(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schema", getSchema(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schemas", listSchemas(config, acapyClient))
	// credential definition
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definition/create", createCredentialDefinition(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definition", getCredentialDefinition(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definitions", listCredentialDefinitions(config, acapyClient))
	// connection
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/connection/invitation", invitation(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/connections", listConnections(config, acapyClient))
	// credential
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/credential/requests", listCredentialRequests(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/credential/offer", credentialOffer(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/credential/issue", issueCredential(config, acapyClient))

	r.Handle("/", http.FileServer(getFileSystem()))

	return r
}

func getFileSystem() http.FileSystem {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		fmt.Println(err)
	}

	return http.FS(fsys)
}
