package models

type CornerstoneCredentialRequest struct {
	IdentityNumber string `json:"identity_number"`
	Names          string `json:"names"`
	Surname        string `json:"surname"`
	Gender         string `json:"gender"`
	DateOfBirth    string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
	Nationality    string `json:"nationality"`
	CitizenStatus  string `json:"citizen_status"`
	IdentityPhoto  string `json:"identity_photo"`
	CredDate       string `json:"cred_date"`
	Email          string `json:"email"`
}
