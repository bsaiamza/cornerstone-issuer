package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"cornerstone-issuer/pkg/acapy"
	"cornerstone-issuer/pkg/config"
	"cornerstone-issuer/pkg/log"
	"cornerstone-issuer/pkg/models"
	"cornerstone-issuer/pkg/server"
	"cornerstone-issuer/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
)

func getCredential(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(getCredentialHandler(config, acapy, cache), mdw...)
}

func getCredentialHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating credential request...")

		// Step 1: Retrieve user information
		var userInfo models.CornerstoneCredentialRequest
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		if err != nil {
			log.Error.Printf("Failed to decode credential data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to decode credential data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Validate ID number
		log.Info.Println("Validating ID number...")
		validID := userInfo.IdentityNumber
		validGender := userInfo.Gender
		validCOB := userInfo.CountryOfBirth
		userID, err := utils.IDValidator(validID, validGender, validCOB)
		if err != nil {
			log.Error.Printf("ID validation failed: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "ID validation failed: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		log.Info.Println("ID Validation passed")

		// Step 3: Call DHA API and compare user information
		dhaSwitch := config.GetDHAAPISwitch()

		if dhaSwitch == "1" {
			log.Info.Printf("Calling DHA API to get user data with the following ID number as a parameter: %s", userID)

			resp, err := http.Get(config.GetDHAAPI() + userID)
			if err != nil {
				log.Error.Printf("Failed on DHA API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed on DHA API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error.Printf("Failed to read DHA API response body: %s", err)
			}

			fmt.Println("\n")
			fmt.Println("DHA API response: ", string(body))

			var dhaData models.DHAResponse
			// err = xml.NewDecoder(resp.Body).Decode(&dhaData)
			err = xml.Unmarshal(body, &dhaData)
			if err != nil {
				log.Error.Printf("Failed to decode DHA API response: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "ID number not found!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// Check death status
			if dhaData.Person.DeathStatus == "DEAD" {
				log.Error.Println("Failed: user is deceased!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: user is deceased!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// dha
			dhaID := dhaData.Person.IDNumber
			dhaNames := strings.ToLower(dhaData.Person.Names)
			dhaSurname := strings.ToLower(dhaData.Person.Surname)
			dhaGender := strings.ToLower(dhaData.Person.Gender)
			// dhaNationality := strings.ToLower(dhaData.Root.Person.Nationality)
			dhaCOB := dhaData.Person.BirthPlace
			if dhaCOB == "SOUTH AFRICA" {
				dhaCOB = "RSA"
			}

			// user
			userNames := strings.ToLower(userInfo.Names)
			userSurname := strings.ToLower(userInfo.Surname)
			userGender := strings.ToLower(userInfo.Gender)
			// userNationality := strings.ToLower(userInfo.Nationality)
			userCOB := userInfo.CountryOfBirth
			if userInfo.CountryOfBirth == "South Africa" {
				userCOB = "RSA"
			}

			// compare
			fmt.Println("\n")
			fmt.Println("DHA Data vs User Data")
			fmt.Println("IDs: " + dhaID + " - " + userID)
			fmt.Println("Names: " + dhaNames + " - " + userNames)
			fmt.Println("Surname: " + dhaSurname + " - " + userSurname)
			fmt.Println("Gender: " + dhaGender + " - " + userGender)
			// fmt.Println("Nationality: " + dhaNationality + " - " + userNationality)
			fmt.Println("Country: " + dhaCOB + " - " + userCOB)
			fmt.Println("\n")

			if !(dhaID == userID && dhaNames == userNames && dhaSurname == userSurname &&
				dhaGender == userGender /*&& dhaNationality == userNationality*/ && dhaCOB == userCOB) {
				log.Error.Println("Failed: user data does not match DHA data!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: data does not match DHA data!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}
		} else {
			log.Info.Printf("Calling DHA API Simulator to get user data with the following ID number as a parameter: %s", userID)

			client := &http.Client{}
			req, err := http.NewRequest(http.MethodGet, config.GetDHASimulatorAPI()+userID, nil)
			if err != nil {
				log.Error.Printf("Failed to create request for DHA API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to create request for DHA API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Error.Printf("Failed on DHA simulator API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed on DHA simulator API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			defer resp.Body.Close()

			// err checking because theres no error struct
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error.Printf("Failed to read DHA simulator API response body: %s", err)
			}

			fmt.Println("\n")
			fmt.Println("DHA API Simulator response: ", string(body))

			if resp.StatusCode >= 400 {
				log.Error.Println("Failed to find user in DHA simulator records!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to find user in DHA simulator records!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			var dhaSimulatorData models.DHASimulatorResponse
			err = json.Unmarshal(body, &dhaSimulatorData)
			if err != nil {
				log.Error.Printf("Failed to decode dha simulator data: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to decode dha simulator data: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// dha
			dhaID := dhaSimulatorData.IDNumber
			dhaNames := strings.ToLower(dhaSimulatorData.Names)
			dhaSurname := strings.ToLower(dhaSimulatorData.Surname)
			dhaGender := strings.ToLower(dhaSimulatorData.Sex)
			dhaDOB := dhaSimulatorData.DateOfBirth
			dhaCOB := dhaSimulatorData.CountryOfBirth
			dhaNationality := dhaSimulatorData.Nationality

			// user
			userNames := strings.ToLower(userInfo.Names)
			userSurname := strings.ToLower(userInfo.Surname)
			userGender := strings.ToLower(string(userInfo.Gender[0]))
			userDOB := string(userInfo.DateOfBirth[0:4]) + "/" + string(userInfo.DateOfBirth[5:7]) + "/" + string(userInfo.DateOfBirth[8:10])
			userCOB := userInfo.CountryOfBirth
			if userCOB == "South Africa" {
				userCOB = "RSA"
			}
			userNationality := userInfo.Nationality

			fmt.Println("\n")
			fmt.Println("DHA Data vs User Data")
			fmt.Println("IDs: " + dhaID + " - " + userID)
			fmt.Println("Names: " + dhaNames + " - " + userNames)
			fmt.Println("Surname: " + dhaSurname + " - " + userSurname)
			fmt.Println("Gender: " + dhaGender + " - " + userGender)
			fmt.Println("DOB: " + dhaDOB + " - " + userDOB)
			fmt.Println("Country: " + dhaCOB + " - " + userCOB)
			fmt.Println("Nationality: " + dhaNationality + " - " + userNationality)
			fmt.Println("\n")

			if !(dhaID == userID && dhaNames == userNames && dhaSurname == userSurname &&
				dhaGender == userGender && dhaDOB == userDOB && dhaCOB == userCOB && dhaNationality == userNationality) {
				log.Error.Println("Failed: user data does not match DHA simulator data!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: data does not match DHA simulator data!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}
		}

		// Step 4: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := acapy.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 4: Cache user data for webhookEventsHandler
		err = cache.User(invitation.Invitation.RecipientKeys[0], userInfo)
		if err != nil {
			log.Error.Printf("Failed to cache user data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache user data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		w.WriteHeader(http.StatusOK)
		res := server.Response{
			"success":    true,
			"credential": invitation.InvitationURL,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func getCredentialByEmail(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(getCredentialByEmailHandler(config, acapy, cache), mdw...)
}

func getCredentialByEmailHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Creating credential request...")

		// Step 1: Retrieve user information
		var userInfo models.CornerstoneCredentialRequest
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		if err != nil {
			log.Error.Printf("Failed to decode credential data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to decode credential data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 2: Validate email address
		err = utils.ValidEmail(userInfo.Email)
		if err != nil {
			log.Error.Printf("Failed %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 3: Validate ID number
		log.Info.Println("Validating ID number...")
		validID := userInfo.IdentityNumber
		validGender := userInfo.Gender
		validCOB := userInfo.CountryOfBirth
		userID, err := utils.IDValidator(validID, validGender, validCOB)
		if err != nil {
			log.Error.Printf("ID validation failed: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "ID validation failed: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		log.Info.Println("ID Validation passed")

		// Step 4: Call DHA API and compare user information
		dhaSwitch := config.GetDHAAPISwitch()

		if dhaSwitch == "1" {
			log.Info.Printf("Calling DHA API to get user data with the following ID number as a parameter: %s", userID)

			resp, err := http.Get(config.GetDHAAPI() + userID)
			if err != nil {
				log.Error.Printf("Failed on DHA API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed on DHA API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error.Printf("Failed to read DHA API response body: %s", err)
			}

			fmt.Println("\n")
			fmt.Println("DHA API response: ", string(body))

			var dhaData models.DHAResponse
			// err = xml.NewDecoder(resp.Body).Decode(&dhaData)
			err = xml.Unmarshal(body, &dhaData)
			if err != nil {
				log.Error.Printf("Failed to decode DHA API response: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "ID number not found!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// Check death status
			if dhaData.Person.DeathStatus == "DEAD" {
				log.Error.Println("Failed: user is deceased!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: user is deceased!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// dha
			dhaID := dhaData.Person.IDNumber
			dhaNames := strings.ToLower(dhaData.Person.Names)
			dhaSurname := strings.ToLower(dhaData.Person.Surname)
			dhaGender := strings.ToLower(dhaData.Person.Gender)
			// dhaNationality := strings.ToLower(dhaData.Root.Person.Nationality)
			dhaCOB := dhaData.Person.BirthPlace
			if dhaCOB == "SOUTH AFRICA" {
				dhaCOB = "RSA"
			}

			// user
			userNames := strings.ToLower(userInfo.Names)
			userSurname := strings.ToLower(userInfo.Surname)
			userGender := strings.ToLower(userInfo.Gender)
			// userNationality := strings.ToLower(userInfo.Nationality)
			userCOB := userInfo.CountryOfBirth
			if userInfo.CountryOfBirth == "South Africa" {
				userCOB = "RSA"
			}

			// compare
			fmt.Println("\n")
			fmt.Println("DHA Data vs User Data")
			fmt.Println("IDs: " + dhaID + " - " + userID)
			fmt.Println("Names: " + dhaNames + " - " + userNames)
			fmt.Println("Surname: " + dhaSurname + " - " + userSurname)
			fmt.Println("Gender: " + dhaGender + " - " + userGender)
			// fmt.Println("Nationality: " + dhaNationality + " - " + userNationality)
			fmt.Println("Country: " + dhaCOB + " - " + userCOB)
			fmt.Println("\n")

			if !(dhaID == userID && dhaNames == userNames && dhaSurname == userSurname &&
				dhaGender == userGender /*&& dhaNationality == userNationality*/ && dhaCOB == userCOB) {
				log.Error.Println("Failed: user data does not match DHA data!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: data does not match DHA data!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}
		} else {
			log.Info.Printf("Calling DHA API Simulator to get user data with the following ID number as a parameter: %s", userID)

			client := &http.Client{}
			req, err := http.NewRequest(http.MethodGet, config.GetDHASimulatorAPI()+userID, nil)
			if err != nil {
				log.Error.Printf("Failed to create request for DHA API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to create request for DHA API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				log.Error.Printf("Failed on DHA simulator API call: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed on DHA simulator API call: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			defer resp.Body.Close()

			// err checking because theres no error struct
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error.Printf("Failed to read DHA simulator API response body: %s", err)
			}

			fmt.Println("\n")
			fmt.Println("DHA API Simulator response: ", string(body))

			if resp.StatusCode >= 400 {
				log.Error.Println("Failed to find user in DHA simulator records!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to find user in DHA simulator records!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			var dhaSimulatorData models.DHASimulatorResponse
			err = json.Unmarshal(body, &dhaSimulatorData)
			if err != nil {
				log.Error.Printf("Failed to decode dha simulator data: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed to decode dha simulator data: " + err.Error(),
				}
				json.NewEncoder(w).Encode(res)
				return
			}

			// dha
			dhaID := dhaSimulatorData.IDNumber
			dhaNames := strings.ToLower(dhaSimulatorData.Names)
			dhaSurname := strings.ToLower(dhaSimulatorData.Surname)
			dhaGender := strings.ToLower(dhaSimulatorData.Sex)
			dhaDOB := dhaSimulatorData.DateOfBirth
			dhaCOB := dhaSimulatorData.CountryOfBirth
			dhaNationality := dhaSimulatorData.Nationality

			// user
			userNames := strings.ToLower(userInfo.Names)
			userSurname := strings.ToLower(userInfo.Surname)
			userGender := strings.ToLower(string(userInfo.Gender[0]))
			userDOB := string(userInfo.DateOfBirth[0:4]) + "/" + string(userInfo.DateOfBirth[5:7]) + "/" + string(userInfo.DateOfBirth[8:10])
			userCOB := userInfo.CountryOfBirth
			if userInfo.CountryOfBirth == "South Africa" {
				userCOB = "RSA"
			}
			userNationality := userInfo.Nationality

			fmt.Println("\n")
			fmt.Println("DHA Data vs User Data")
			fmt.Println("IDs: " + dhaID + " - " + userID)
			fmt.Println("Names: " + dhaNames + " - " + userNames)
			fmt.Println("Surname: " + dhaSurname + " - " + userSurname)
			fmt.Println("Gender: " + dhaGender + " - " + userGender)
			fmt.Println("DOB: " + dhaDOB + " - " + userDOB)
			fmt.Println("Country: " + dhaCOB + " - " + userCOB)
			fmt.Println("Nationality: " + dhaNationality + " - " + userNationality)
			fmt.Println("\n")

			if !(dhaID == userID && dhaNames == userNames && dhaSurname == userSurname &&
				dhaGender == userGender && dhaDOB == userDOB && dhaCOB == userCOB && dhaNationality == userNationality) {
				log.Error.Println("Failed: user data does not match DHA simulator data!")
				w.WriteHeader(http.StatusInternalServerError)
				res := server.Response{
					"success": false,
					"msg":     "Failed: data does not match DHA simulator data!",
				}
				json.NewEncoder(w).Encode(res)
				return
			}
		}

		// Step 5: Create Invitation
		invitationRequest := models.CreateInvitationRequest{}

		invitation, err := acapy.CreateInvitation(invitationRequest)
		if err != nil {
			log.Error.Printf("Failed to create invitation: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create invitation: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 6: Cache user data for webhookEventsHandler
		err = cache.User(invitation.Invitation.RecipientKeys[0], userInfo)
		if err != nil {
			log.Error.Printf("Failed to cache user data: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to cache user data: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 7: Generate a qr code for email
		qrCodePng, err := qrcode.Encode(invitation.InvitationURL, qrcode.Medium, 256)
		if err != nil {
			log.Warning.Print("Failed to create QR code: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to create QR code: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 8: Send email
		prefix := string(userInfo.IdentityNumber[6])
		if prefix >= "5" {
			prefix = "Mr "
		} else {
			prefix = "Ms/Mrs "
		}

		err = utils.SendCredentialByEmail(prefix+userInfo.Surname, userInfo.Email, invitation.Invitation.RecipientKeys[0], qrCodePng, config)
		if err != nil {
			log.Warning.Print("Failed to send credential by email: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to send credential by email",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Step 9: Remove qr from os once email is sent
		err = os.Remove("./" + invitation.Invitation.RecipientKeys[0] + ".png")
		if err != nil {
			log.Warning.Print("Failed to remove QR code: ", err)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func webhookEvents(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(webhookEventsHandler(config, acapy, cache), mdw...)
}

func webhookEventsHandler(config *config.Config, acapy *acapy.Client, cache *utils.BigCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		topic := mux.Vars(r)["topic"]

		switch topic {
		case "connections":
			var request models.Connection
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				log.Error.Printf("Failed to decode request body: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "response" {
				pingRequest := models.PingConnectionRequest{
					Comment: "Ping",
				}

				_, err := acapy.PingConnection(request.ConnectionID, pingRequest)
				if err != nil {
					log.Error.Printf("Failed to ping holder: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			if request.State == "active" {
				userInfo, err := cache.ReadUser(request.InvitationKey)
				if err != nil {
					log.Error.Printf("Failed to read cached user data: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				dobJSON, _ := time.Parse("2006-01-02", userInfo.DateOfBirth)
				dob := dobJSON.Format("20060102")
				idPhoto := utils.ImageBase64()
				cd := time.Now().Format("20060102")

				credentialRequest := models.IssueCredentialRequest{
					AutoRemove:      false,
					ConnectionID:    request.ConnectionID,
					Comment:         "Cornerstone Credential",
					CredDefID:       config.GetCredDefID(),
					IssuerDid:       config.GetPublicDID(),
					SchemaID:        config.GetSchemaID(),
					SchemaIssuerDid: config.GetPublicDID(),
					SchemaName:      config.GetSchemaName(),
					SchemaVersion:   config.GetSchemaVersion(),
					Trace:           false,
					CredentialProposal: models.CredentialProposal{
						Type: "issue-credential/1.0/credential-preview",
						Attributes: []models.Attribute{
							{
								Name:  "identity_number",
								Value: userInfo.IdentityNumber,
							},
							{
								Name:  "names",
								Value: userInfo.Names,
							},
							{
								Name:  "surname",
								Value: userInfo.Surname,
							},
							{
								Name:  "gender",
								Value: userInfo.Gender,
							},
							{
								Name:  "date_of_birth",
								Value: dob,
							},
							{
								Name:  "country_of_birth",
								Value: userInfo.CountryOfBirth,
							},
							{
								Name:  "nationality",
								Value: userInfo.Nationality,
							},
							{
								Name:  "citizen_status",
								Value: userInfo.CitizenStatus,
							},
							{
								Name:  "identity_photo",
								Value: idPhoto,
							},
							{
								Name:  "cred_date",
								Value: cd,
							},
						},
					},
				}

				_, err = acapy.IssueCredential(credentialRequest)
				if err != nil {
					log.Error.Printf("Failed to send credential offer: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// For email credential notification
				err = cache.User(request.ConnectionID, userInfo)
				if err != nil {
					log.Error.Printf("Failed to cache user data: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteUser(request.InvitationKey)

				log.Info.Println("Credential offer sent")
			}

		case "issue_credential":
			var request models.IssueCredentialWebhookResponse
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				log.Error.Printf("Failed to decode request body: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			userInfo, err := cache.ReadUser(request.ConnectionID)
			if err != nil {
				log.Error.Printf("Failed to read cached user data: %s", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if request.State == "credential_issued" && userInfo.Email != "" {
				log.Info.Println("Sending credential issued notification...")

				prefix := string(userInfo.IdentityNumber[6])
				if prefix >= "5" {
					prefix = "Mr "
				} else {
					prefix = "Ms/Mrs "
				}

				err = utils.SendNotificationEmail(prefix+userInfo.Surname, userInfo.Email, config)
				if err != nil {
					log.Error.Printf("Failed to send credential notification email: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				cache.DeleteUser(request.ConnectionID)

				log.Info.Println("Notified user successfully about issued credential!")
			}

			if request.State == "credential_issued" {
				txnCounterSwitch := config.GetTxnCounterSwitch()

				if txnCounterSwitch == "1" {
					log.Info.Println("Calling Transaction Counter")

					txnID := utils.RandomTxnID(12)

					did := "BER7WwiAMK9igkiRjPYpEp"
					domain := "IAMZA Cornerstone Issuer"
					notes := map[string]interface{}{
						"txnid": txnID,
					}

					token, err := utils.CreateToken(jwt.SigningMethodRS256, did, domain, notes, config.GetTxnCounterPK())
					if err != nil {
						log.Error.Printf("Failed to create jwt: %s", err)
						return
					}

					url := config.GetTxnCounterAPI() + token

					req, err := http.NewRequest("POST", url, nil)
					if err != nil {
						log.Error.Printf("Failed to create new Transaction Counter API request: %s", err)
						return
					}

					req.Header.Add("Content-Type", "application/json")

					n, err := strconv.Atoi(config.GetTxnCounterLoopN())
					if err != nil {
						log.Error.Printf("Failed to parse loop n: %s", err)
						return
					}

					if config.GetTxnCounterLoopSwitch() == "1" {
						log.Info.Printf("Looping Txn Counter %d times", n)

						for i := 0; i <= n; i++ {
							res, err := http.DefaultClient.Do(req)
							if err != nil {
								log.Error.Printf("Failed on Transaction Counter API call: %s", err)
								return
							}

							defer res.Body.Close()
							body, _ := ioutil.ReadAll(res.Body)

							fmt.Println("\n")
							fmt.Println("Transaction Counter API response: ", string(body))
						}
					} else {
						res, err := http.DefaultClient.Do(req)
						if err != nil {
							log.Error.Printf("Failed on Transaction Counter API call: %s", err)
							return
						}

						defer res.Body.Close()
						body, _ := ioutil.ReadAll(res.Body)

						fmt.Println("\n")
						fmt.Println("Transaction Counter API response: ", string(body))
					}

				}
			}

		case "present_proof":
		case "basicmessages":
		case "revocation_registry":
		case "problem_report":
		case "issuer_cred_rev":

		default:
			log.Warning.Printf("Unexpected topic: %s", topic)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
