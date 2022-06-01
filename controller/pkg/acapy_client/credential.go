package acapy

import (
	"cornerstone_issuer/pkg/config"
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models/exchange_records"
	"cornerstone_issuer/pkg/models/offer"
	"cornerstone_issuer/pkg/models/proposals"
	"fmt"
)

// ListCredentialExchangeRecords returns all credential exchange records.
func (c *Client) ListCredentialExchangeRecords(params *exchange_records.ListCredentialExchangeRecordsParams) (exchange_records.ListCredentialExchangeRecordsResponse, error) {
	var credentialExchangeRecords exchange_records.ListCredentialExchangeRecordsResponse

	var queryParams = map[string]string{}
	if params != nil {
		queryParams = map[string]string{
			"connection_id": params.ConnectionID,
			"role":          params.Role,
			"state":         params.State,
			"thread_id":     params.ThreadID,
		}
	}

	err := c.get("/issue-credential-2.0/records", queryParams, &credentialExchangeRecords)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential-2.0/records: %s", err.Error())
		return exchange_records.ListCredentialExchangeRecordsResponse{}, err
	}
	return credentialExchangeRecords, nil
}

// CornerstoneCredentialProposal sends a credential offer.
func (c *Client) CornerstoneCredentialOffer(credExID string, request offer.CredentialOfferRequest, config *config.Config) (proposals.CornerstoneCredentialProposalResponse, error) {
	var proposal proposals.CornerstoneCredentialProposalResponse

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/send-offer", credExID), nil, request, &proposal)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential-2.0/{cred_ex_id}/send-offer: %s", err.Error())
		return proposals.CornerstoneCredentialProposalResponse{}, err
	}
	return proposal, nil
}

// CornerstoneCredentialProposal sends a credential offer.
func (c *Client) CornerstoneCredentialOfferV2(request offer.CredentialOfferRequest) (proposals.CornerstoneCredentialProposalResponse, error) {
	var proposal proposals.CornerstoneCredentialProposalResponse

	err := c.post("/issue-credential-2.0/send-offer", nil, request, &proposal)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /issue-credential-2.0/send-offer: %s", err.Error())
		return proposals.CornerstoneCredentialProposalResponse{}, err
	}
	return proposal, nil
}

// IssueCredential issues a credential.
func (c *Client) IssueCredential(credExID string, request exchange_records.IssueCredentialRequest) (exchange_records.CredExRecord, error) {
	var credential exchange_records.CredExRecord

	err := c.post(fmt.Sprintf("/issue-credential-2.0/records/%s/issue", credExID), nil, request, &credential)
	if err != nil {
		log.ServerError.Print("Failed on ACA-py /issue-credential/records/{cred_ex_id}/issue: ", err)
		return exchange_records.CredExRecord{}, err
	}
	return credential, nil
}

// IssueCornerstoneCredential
// func (c *Client) IssueCornerstoneCredential(connectionID, did, idNumber, familyName, firstName, dob, gender, address, credentialSchemaID string) (models.IssueCornerstoneCredentialResponse, error) {
// 	var request = models.CornerstoneCredential{
// 		ConnectionID: connectionID,
// 		Comment:      "Your Cornerstone Credential.",
// 		Filter: models.Filter{
// 			LdProof: models.LdProof{
// 				Credential: models.Credential{
// 					Context:        []string{"https://www.w3.org/2018/credentials/v1", "https://json-schema.org/draft/2020-12/schema"},
// 					Type:           []string{"VerifiableCredential", "VerifiableAttestation", "VerifiableId"},
// 					Issuer:         did,
// 					IssuanceDate:   time.Now().String(),
// 					ExpirationDate: "", // time.Now().Add(time.Hour * 24 * 365).String(),
// 					CredentialSubject: models.CredentialSubject{
// 						IDNumber:   idNumber,
// 						FamilyName: familyName,
// 						FirstName:  firstName,
// 						DOB:        dob,
// 						Gender:     gender,
// 						Address:    address,
// 					},
// 					CredentialStatus: models.CredentialStatus{
// 						ID:   "Cornerstone_Credential",
// 						Type: "CornerstoneCredentialStatusList2022",
// 					},
// 					CredentialSchema: models.CredentialSchema{
// 						ID:   credentialSchemaID,
// 						Type: "CornerstoneCredentialSchema",
// 					},
// 					Evidence: models.Evidence{
// 						ID:               "http://example.com/evidence/government-issued-id",
// 						Type:             []string{"DocumentVerification"},
// 						Verifier:         did,
// 						EvidenceDocument: "SmartID",
// 						SubjectPresence:  "Physical",
// 						DocumentPresence: "Physical",
// 					},
// 					Proof: models.Proof{
// 						Type:               "Ed25519Signature2018",
// 						Created:            time.Now().String(),
// 						Jws:                "",
// 						ProofPurpose:       "assertionMethod",
// 						VerificationMethod: did,
// 					},
// 				},
// 			},
// 		},
// 	}

// 	var response models.IssueCornerstoneCredentialResponse

// 	err := c.post("/issue-credential-2.0/send", nil, request, &response)
// 	if err != nil {
// 		log.Error.Printf("Failed on ACA-py /issue-credential-2.0/send: %s", err)
// 		return models.IssueCornerstoneCredentialResponse{}, err
// 	}
// 	return response, nil
// }
