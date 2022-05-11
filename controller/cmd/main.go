package main

import (
	"cornerstone_issuer/api/v1"
	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/server"
)

func main() {
	config := config.GetConfig()

	host := config.GetServerHost()
	port := config.GetServerPort()
	serverAddress := host + ":" + port

	client := acapy.NewClient(config.GetAcapyURL())

	srv := server.NewServer().
		WithAddr(serverAddress).
		WithRouter(api.NewRouter(config, client)).
		WithErrLogger(log.ServerError)

	go func() {
		log.ServerInfo.Print("-------------------------------------------------")
		log.ServerInfo.Print("|		Cornerstone Issuer		|")
		log.ServerInfo.Print("-------------------------------------------------")
		log.ServerInfo.Print("		**ENV VARS**")
		log.ServerInfo.Print("	ACAPY_URL: ", config.GetAcapyURL())
		log.ServerInfo.Print("	CLIENT_URL: ", config.GetClientURL())
		log.ServerInfo.Print("	SERVER_HOST: ", config.GetServerHost(),)
		log.ServerInfo.Print("	SERVER_PORT: ", config.GetServerPort(),)
		log.ServerInfo.Print("-------------------------------------------------")
		log.ServerInfo.Print("")
		log.ServerInfo.Printf("Server started on: %s", serverAddress)
		if err := srv.Start(); err != nil {
			log.ServerError.Fatal(err)
		}
	}()

	server.GracefulExit(func() {
		if err := srv.Stop(); err != nil {
			log.ServerError.Print("Failed to shutdown server gracefully:", err)
		}
	})
}
