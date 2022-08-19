package main

import (
	"cornerstone_issuer/api"
	"cornerstone_issuer/pkg/client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/server"
	"cornerstone_issuer/pkg/utils"
	"crypto/tls"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	runServer()
}

func runServer() {
	config := config.LoadConfig()
	acapyClient := client.NewClient(config.GetAcapyURL())
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
