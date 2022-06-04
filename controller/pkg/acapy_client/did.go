package acapy

import (
	"cornerstone_issuer/pkg/log"
	"cornerstone_issuer/pkg/models"
)

// GetDID returns a public DID.
// * actually a list function but the assumption is only one DID is registered and returned.
func (c *Client) GetDID() (models.GetPublicDidResponse, error) {
	var did models.GetPublicDidResponse

	queryParams := map[string]string{
		"did":      "",
		"key_type": "",
		"method":   "",
		"posture":  "",
		"verkey":   "",
	}

	err := c.get("/wallet/did/public", queryParams, &did)
	if err != nil {
		log.Error.Printf("Failed on ACA-py /wallet/did/public: %s", err.Error())
		return models.GetPublicDidResponse{}, err
	}
	return did, nil
}
