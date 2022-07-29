package api

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strconv"

	"cornerstone_issuer/pkg/client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/server"
)

func GetDHAUser(config *config.Config, client *client.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(getDHAUserHandler(config, client), mdw...)
}
func getDHAUserHandler(config *config.Config, client *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.DHAWarning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		// Step 1: Retrieve ID query string 
		idNumber := r.URL.Query().Get("id_number")

		dhaData, err := client.DHARequest(config.GetDHAAPI() + idNumber)
		if err != nil {
			log.DHAError.Printf("Failed on DHA API call: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed on DHA API call: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		dob := ""
		dobYear, err := strconv.Atoi(dhaData.Root.Person.IDNumber[0:1])
		if err != nil {
			log.DHAError.Printf("Failed to convert DOB: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed to convert DOB: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		if dobYear > 2 {
			dob = "19" + dhaData.Root.Person.IDNumber[0:2] + "-" + dhaData.Root.Person.IDNumber[2:4] + "-" + dhaData.Root.Person.IDNumber[4:6]
		} else {
			dob = "20" + dhaData.Root.Person.IDNumber[0:2] + "-" + dhaData.Root.Person.IDNumber[2:4] + "-" + dhaData.Root.Person.IDNumber[4:6]
		}

		data := models.DHAUser{
			IDNumber: dhaData.Root.Person.IDNumber,
			FirstNames: dhaData.Root.Person.Names,
			Surname: dhaData.Root.Person.Surname,
			Gender: dhaData.Root.Person.Gender,
			DateOfBirth: dob,
			CountryOfBirth: dhaData.Root.Person.BirthPlace,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func GetDHASimulatorUser(config *config.Config, client *client.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.LogAPIRequest,
	}

	return server.ChainMiddleware(getDHASimulatorUserHandler(config, client), mdw...)
}
func getDHASimulatorUserHandler(config *config.Config, client *client.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.DHAWarning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Response{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		idNumber := r.URL.Query().Get("id_number")

		dhaData, err := client.DHASimulatorRequest(config.GetDHASimulatorAPI() + idNumber)
		if err != nil {
			log.DHAError.Printf("Failed on DHA simulator API call: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Response{
				"success": false,
				"msg":     "Failed on DHA simulator API call: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		data := models.DHAUser{
			IDNumber: dhaData.IDNumber,
			FirstNames: dhaData.Names,
			Surname: dhaData.Surname,
			Gender: dhaData.Sex,
			DateOfBirth: dhaData.DateOfBirth,
			CountryOfBirth: dhaData.CountryOfBirth,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

// TEST DHA

		// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		// resp, err := http.Get("https://npr-prod.dha.gov.za:8093/NATCGI/BVRSTEST/NWWAPS/SYSWEB/WEB046N?IDNO=9602295518082")
		// if err != nil {
		// 	log.ServerError.Println(err)
		// }

		// defer resp.Body.Close()

		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.ServerError.Fatal(err)
		// }

		// fmt.Println(string(body))

		// var userInfo models.RealDHASuccessResponse
		// err = xml.NewDecoder(resp.Body).Decode(&userInfo)
		// if err != nil {
		// 	log.Error.Printf("Failed to decode dha data: %s", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	res := server.Response{
		// 		"success": false,
		// 		"msg":     "Failed to decode dha data: " + err.Error(),
		// 	}
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }

		// fmt.Println(userInfo)

		// end

		// log.Info.Println("Listing connections...")

		// connections, err := client.ListConnections()
		// if err != nil {
		// 	log.Error.Printf("Failed to list connections: %s", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	res := server.Response{
		// 		"success": false,
		// 		"msg":     "Failed to list connections: " + err.Error(),
		// 	}
		// 	json.NewEncoder(w).Encode(res)
		// 	return
		// }

		// log.Info.Print("Connections listed successfully!")