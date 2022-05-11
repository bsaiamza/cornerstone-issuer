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

	role := config.GetRole()
	host := config.GetServerHost()
	port := config.GetServerPort()
	serverAddress := host + ":" + port

	startStr := "+ Cornerstone Issuer +"

	if role == "holder" {
		startStr = "+ Cornerstone Holder +"
	}

	client := acapy.NewClient(config.GetAcapyURL())

	srv := server.NewServer().
		WithAddr(serverAddress).
		WithRouter(api.NewRouter(config, client)).
		WithErrLogger(log.ServerError)

	go func() {
		log.ServerInfo.Print("----------------------")
		log.ServerInfo.Print(startStr)
		log.ServerInfo.Print("----------------------")
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
