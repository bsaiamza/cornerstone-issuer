package main

import (
	"cornerstone_issuer/api/v1"
	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/server"
	"cornerstone_issuer/pkg/util"
)

func main() {
	config := config.GetConfig()

	serverAddress := config.GetServerAddress()

	acapyClient := acapy.NewClient(config.GetAcapyURL())

	cache := util.NewBigCache()

	// init connection for ui
	createInvitationRequest := models.CreateInvitationRequest{}
	queryParams := models.CreateInvitationParams{}

	invitation, err := acapyClient.CreateInvitation(createInvitationRequest, &queryParams)
	if err != nil {
		log.Error.Printf("Failed to generate multi invitation: %s", err.Error())
		return
	}

	err = cache.UpdateString("invitationURL", invitation.InvitationURL)
	if err != nil {
		log.Error.Printf("Failed to update cache: %s", err.Error())
		return
	}

	err = cache.UpdateStruct("invitation", invitation.Invitation)
	if err != nil {
		log.Error.Printf("Failed to update cache: %s", err.Error())
		return
	}

	srv := server.NewServer().
		WithAddress(serverAddress).
		WithRouter(api.NewRouter(config, acapyClient, cache)).
		WithErrLogger(log.ServerError)

	go func() {
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("|		Cornerstone Issuer		|")
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("		**ENV VARS**")
		log.ServerInfo.Println("	CLIENT_URL: ", config.GetClientURL())
		log.ServerInfo.Println("	SERVER_ADDRESS: ", config.GetServerAddress())
		log.ServerInfo.Println("	API_BASE_URL: ", config.GetAPIBaseURL())
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("")
		log.ServerInfo.Printf("Server started on: %s", serverAddress)
		if err := srv.Start(); err != nil {
			log.ServerError.Fatal(err)
		}
	}()

	server.GracefulExit(func() {
		if err := srv.Stop(); err != nil {
			log.ServerError.Printf("Failed to stop server: %s", err.Error())
		}
	})
}
