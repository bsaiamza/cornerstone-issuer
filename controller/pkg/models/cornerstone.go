package models

type CredentialRequestData struct {
	Email          string `json:"email"`
	IDNumber       string `json:"id_number"`
	Surname        string `json:"surname"`
	Forenames      string `json:"forenames"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

type CornerstoneCredentialData struct {
	IDNumber       string `json:"id_number"`
	Surname        string `json:"surname"`
	Forenames      string `json:"forenames"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
}

type DhaData struct {
	IDNumber                   string `json:"Identity_Number"`
	Forenames                  string `json:"Names"`
	Surname                    string `json:"Surname"`
	Gender                     string `json:"Sex"`
	IssueDate                  string `json:"Issue_Date"`
	DateOfBirth                string `json:"Date_of_Birth"`
	BiometricsPhoto            string `json:"Biometrics-photo"`
	BiometricsFingerprint      string `json:"Biometrics-fingerprint"`
	BiometricsFingerprintMatch int    `json:"Biometrics-fingerprint_match"`
	Nationality                string `json:"Nationality"`
	CountryOfBirth             string `json:"Country_of_Birth"`
}

type DhaAPIError []DhaAPIErrorElement

type DhaAPIErrorElement struct {
	OriginatingDate string `json:"originatingDate"`
	ResponseCode    int64  `json:"responseCode"`
	ResponseDesc    string `json:"responseDesc"`
	ZaID            string `json:"ZA ID"`
	EndToEndID      string `json:"endToEndId"`
}
