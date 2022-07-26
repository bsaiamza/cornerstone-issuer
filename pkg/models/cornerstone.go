package models

type CornerstoneCredentialRequest struct {
	IDNumber       string `json:"id_number"`
	Surname        string `json:"surname"`
	Forenames      string `json:"forenames"`
	Gender         string `json:"gender"`
	DOB            string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
	Email          string `json:"email"`
}
