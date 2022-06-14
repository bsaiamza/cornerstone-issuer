package api

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/util"
)

//go:embed build
var embeddedFiles embed.FS

func NewRouter(config *config.Config, acapyClient *acapy.Client, cache *util.BigCache) *http.ServeMux {
	r := http.NewServeMux()

	apiBaseURL := config.GetAPIBaseURL()

	// health
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/health", health(config))
	// did
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/did", getPublicDID(config, acapyClient))
	// schema
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schema/create", createSchema(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schema", getSchema(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/schemas", listSchemas(config, acapyClient))
	// credential definition
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definition/create", createCredentialDefinition(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definition", getCredentialDefinition(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/definitions", listCredentialDefinitions(config, acapyClient))
	// connection
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/connection/invitation", invitation(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/connection/multi-invitation", multiInvitation(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/connections", listConnections(config, acapyClient))
	// credential
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/topic/connections/", issueCredential(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/credential", displayCredentialRequest(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/email-credential", emailCredentialRequest(config, acapyClient, cache))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/credentials", listCredentials(config, acapyClient))
	r.HandleFunc(apiBaseURL+"/cornerstone/issuer/topic/issue_credential/", credentialPing(config, acapyClient, cache))

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
