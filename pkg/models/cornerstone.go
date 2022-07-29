package models

type CornerstoneCredentialRequest struct {
	IDNumber       string `json:"id_number"`
	FirstNames      string `json:"first_names"`
	Surname        string `json:"surname"`
	Gender         string `json:"gender"`
	DOB            string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
	Email          string `json:"email"`
}

