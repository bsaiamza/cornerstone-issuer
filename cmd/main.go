package main

import (
	"cornerstone-issuer/api"
	"cornerstone-issuer/pkg/acapy"
	"cornerstone-issuer/pkg/config"
	"cornerstone-issuer/pkg/log"
	"cornerstone-issuer/pkg/server"
	"cornerstone-issuer/pkg/utils"
)

func main() {
	runServer()
}

func runServer() {
	config := config.LoadConfig()
	acapyClient := acapy.NewClient(config.GetAcapyURL())
	cache := utils.NewBigCache()

	srv := server.NewServer().
		WithAddress(config.GetServerAddress()).
		WithRouter(api.NewRouter(config, acapyClient, cache)).
		WithErrorLogger(log.ServerError)

	go func() {
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("|		Cornerstone Issuer		|")
		log.ServerInfo.Println("-------------------------------------------------")
		log.ServerInfo.Println("")
		log.ServerInfo.Printf("Server started on: %s", config.GetServerAddress())
		if err := srv.Start(); err != nil {
			log.ServerError.Fatal(err)
		}
	}()

	utils.GracefulServerExit(func() {
		if err := srv.Stop(); err != nil {
			log.ServerError.Printf("Failed to stop server: %s", err.Error())
		}
	})
}
