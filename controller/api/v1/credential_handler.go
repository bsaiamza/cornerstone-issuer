package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	acapy "cornerstone_issuer/pkg/acapy_client"
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
	"cornerstone_issuer/pkg/models/exchange_records"
	"cornerstone_issuer/pkg/models/offer"
	"cornerstone_issuer/pkg/server"
)

func listCredentialRequests(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(listCredentialRequestsHandler(config, acapyClient), mdw...)
}
func listCredentialRequestsHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "GET, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Println("Listing credential requests...")

		connectionID := r.URL.Query().Get("connection_id")
		role := r.URL.Query().Get("role")
		state := r.URL.Query().Get("state")
		threadID := r.URL.Query().Get("thread_id")

		queryParams := exchange_records.ListCredentialExchangeRecordsParams{
			ConnectionID: connectionID,
			Role:         role,
			State:        state,
			ThreadID:     threadID,
		}

		exchangeRecords, err := acapyClient.ListCredentialExchangeRecords(&queryParams)
		if err != nil {
			log.Error.Printf("Failed to list credential requests: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to list credential requests: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential requests listed successfully!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(exchangeRecords.Results)
	}
}

func credentialOffer(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(credentialOfferHandler(config, acapyClient), mdw...)
}
func credentialOfferHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		credExID := r.URL.Query().Get("cred_ex_id")

		var requestBody offer.CredentialOfferBodyRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			log.Error.Printf("Failed to decode request body: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode request body: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		// log.Info.Print(credExID)
		// log.Info.Print(requestBody.IDNum)

		// TODO: validate ID number
		// Get user from home affairs api using ID number?
		dataFile, err := os.Open("users.json")
		if err != nil {
			log.Error.Printf("Failed to get DHA data: %s", err.Error())
			// w.WriteHeader(http.StatusInternalServerError)
			// res := server.Res{
			// 	"success": false,
			// 	"msg":     "Failed to get DHA data: " + err.Error(),
			// }
			// json.NewEncoder(w).Encode(res)
			// return
		}

		defer dataFile.Close()

		data, _ := ioutil.ReadAll(dataFile)

		var users models.DhaData

		json.Unmarshal(data, &users)

		var dhaUser models.User

		for _, user := range users.Users {
			if user.LprNumber == requestBody.IDNum {
				log.Info.Println("DHA says user exists!")

				dhaUser = models.User{
					GivenName:              user.GivenName,
					FamilyName:             user.FamilyName,
					Gender:                 user.Gender,
					LprNumber:              user.LprNumber,
					LprCategory:            user.LprCategory,
					ResidentSince:          user.ResidentSince,
					CommuterClassification: user.CommuterClassification,
					BirthDate:              user.BirthDate,
					BirthCountry:           user.BirthCountry,
				}
			} else {
				log.Warning.Println("DHA says user does not exists!")
				// w.WriteHeader(http.StatusInternalServerError)
				// res := server.Res{
				// 	"success": false,
				// 	"msg":     "DHA says user does not exists!",
				// }
				// json.NewEncoder(w).Encode(res)
				// return
			}
		}

		schema, err := acapyClient.GetSchema(requestBody.SchemaID)
		if err != nil {
			log.Error.Printf("Failed to get schema: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get schema: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		did, err := acapyClient.GetDID()
		if err != nil {
			log.Error.Printf("Failed to get DID: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to get DID: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		var request = offer.CredentialOfferRequest{
			AutoRemove: true,
			CounterPreview: offer.CounterPreview{
				Type: "issue-credential/2.0/credential-preview",
				Attributes: []offer.Attribute{
					{
						Name:  "id",
						Value: "did:sov:TwvLGLxQPgdmXuJXpq33mh",
					},
					{
						Name:  "givenName",
						Value: dhaUser.GivenName,
					},
					{
						Name:  "familyName",
						Value: dhaUser.FamilyName,
					},
					{
						Name:  "gender",
						Value: dhaUser.Gender,
					},
					{
						Name:  "lprNumber",
						Value: dhaUser.LprNumber,
					},
					{
						Name:  "lprCategory",
						Value: dhaUser.LprCategory,
					},
					{
						Name:  "residentSince",
						Value: dhaUser.ResidentSince,
					},
					{
						Name:  "commuterClassification",
						Value: dhaUser.CommuterClassification,
					},
					{
						Name:  "birthDate",
						Value: dhaUser.BirthDate,
					},
					{
						Name:  "birthCountry",
						Value: dhaUser.BirthCountry,
					},
				},
			},
			Filter: offer.Filter{
				Indy: offer.Indy{
					CredDefID:       requestBody.CredDefID,
					IssuerDid:       did.Result.Did,
					SchemaIssuerDid: did.Result.Did,
					SchemaID:        schema.Schema.ID,
					SchemaName:      schema.Schema.Name,
					SchemaVersion:   schema.Schema.Ver,
				},
				LdProof: offer.LdProof{
					Credential: offer.Credential{
						Context:      []string{"https://www.w3.org/2018/credentials/v1", "https://w3id.org/citizenship/v1"},
						ID:           "https://issuer.oidp.uscis.gov/credentials/83627465",
						Type:         []string{"VerifiableCredential", "PermanentResidentCard"},
						Issuer:       "did:sov:" + did.Result.Did,
						Identifier:   "83627465",
						Name:         "Identity Document",
						Description:  "Government issued Smart ID card.",
						IssuanceDate: time.Now().Format(time.RFC3339),
						CredentialSubject: offer.CredentialSubject{
							ID:                     "did:sov:TwvLGLxQPgdmXuJXpq33mh",
							Type:                   []string{"PermanentResident", "Person"},
							GivenName:              dhaUser.GivenName,
							FamilyName:             dhaUser.FamilyName,
							Gender:                 dhaUser.Gender,
							Image:                  "",
							ResidentSince:          dhaUser.ResidentSince,
							LprCategory:            dhaUser.LprCategory,
							LprNumber:              dhaUser.LprNumber,
							CommuterClassification: dhaUser.CommuterClassification,
							BirthCountry:           dhaUser.BirthCountry,
							BirthDate:              dhaUser.BirthDate,
						},
						Proof: offer.Proof{
							Type:               "Ed25519Signature2018",
							Created:            time.Now().Format(time.RFC3339),
							Jws:                "eyJhbGciOiJFZERTQSIsI...wRG2fNmAx60Vi4Ag",
							ProofPurpose:       "assertionMethod",
							VerificationMethod: "did:sov:TwvLGLxQPgdmXuJXpq33mh#key1",
						},
					},
					Options: offer.Options{
						ProofType: "Ed25519Signature2018",
					},
				},
			},
			Trace: false,
		}

		offer, err := acapyClient.CornerstoneCredentialOffer(credExID, request, config)
		if err != nil {
			log.Error.Printf("Failed to send credential offer: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to send credential offer: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential offer sent!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(offer)
	}
}

func issueCredential(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	mdw := []server.Middleware{
		server.NewLogRequest,
	}

	return server.ChainMiddleware(issueCredentialHandler(config, acapyClient), mdw...)
}
func issueCredentialHandler(config *config.Config, acapyClient *acapy.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", config.GetClientURL())
		header.Add("Access-Control-Allow-Methods", "POST, OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			log.Warning.Print("Incorrect request method!")
			w.WriteHeader(http.StatusMethodNotAllowed)
			res := server.Res{
				"success": false,
				"msg":     "Warning: Incorrect request method!",
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		defer r.Body.Close()

		log.Info.Print("Issuing credential...")

		credExID := r.URL.Query().Get("cred_ex_id")

		var issueRequest exchange_records.IssueCredentialRequest

		err := json.NewDecoder(r.Body).Decode(&issueRequest)
		if err != nil {
			log.Error.Printf("Failed to decode issue request: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			res := server.Res{
				"success": false,
				"msg":     "Failed to decode issue request: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		credential, err := acapyClient.IssueCredential(credExID, issueRequest)
		if err != nil {
			log.Error.Printf("Failed to issue credential: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			res := server.Res{
				"success": false,
				"msg":     "Failed to issue credential: " + err.Error(),
			}
			json.NewEncoder(w).Encode(res)
			return
		}

		log.Info.Print("Credential issued!")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(credential)
	}
}
